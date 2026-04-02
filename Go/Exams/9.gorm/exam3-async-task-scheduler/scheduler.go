package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"gorm.io/gorm"
)

// TaskScheduler 任务调度器
type TaskScheduler struct {
	db          *gorm.DB
	taskQueue   chan uint     // 任务 ID 队列
	quit        chan struct{} // 退出信号
	workerCount int           // Worker 数量
	wg          *sync.WaitGroup
	closed      int32 // 标记是否已关闭，使用 atomic 操作
}

// NewTaskScheduler 创建任务调度器
// 要求：
// 1. taskQueue 缓冲区大小设置为 100
// 2. 启动指定数量的 Worker Goroutine
// 3. 返回初始化好的 TaskScheduler
func NewTaskScheduler(db *gorm.DB, workerCount int) *TaskScheduler {
	ts := TaskScheduler{
		db:          db,
		taskQueue:   make(chan uint, 100),
		quit:        make(chan struct{}),
		workerCount: workerCount,
		wg:          &sync.WaitGroup{},
	}
	for i := range workerCount {
		ts.wg.Add(1)
		go ts.worker(i)
	}
	return &ts
}

// SubmitTask 提交任务到队列
// 注意：使用 select 防止 Channel 满时阻塞，如果 Channel 已满返回错误
func (s *TaskScheduler) SubmitTask(taskID uint) error {
	// 检查调度器是否已关闭，避免向已关闭的 channel 写入导致 panic
	if atomic.LoadInt32(&s.closed) == 1 {
		return fmt.Errorf("调度器已关闭，无法提交任务 taskID: %d", taskID)
	}
	select {
	case s.taskQueue <- taskID:
	default:
		return fmt.Errorf("任务队列已满，任务添加失败 taskID: %d", taskID)
	}
	return nil
}

// worker 工作协程
// 要求：
// 1. 监听 taskQueue 和 quit 信号
// 2. 收到 quit 信号时优雅退出
// 3. 调用 processTask 处理任务
func (s *TaskScheduler) worker(id int) {
	defer s.wg.Done()

	// ❌ 使用 select 同时监听 taskQueue 和 quit，避免使用额外的 goroutine 和共享变量
	// 这样可以避免竞态条件（quit 变量没有同步保护）和 goroutine 泄漏
	for {
		// 虽然这个 select 可能继续执行，可能直接退出，但问题不大（这并非竞态条件，不会出bug）
		// 可按需求将 <-s.quit 改为“drain掉所有剩余任务”
		select {
		case taskID, ok := <-s.taskQueue:
			if !ok {
				// Channel 已关闭，优雅退出
				return
			}
			if err := s.processTask(taskID); err != nil {
				// 最佳实践：使用结构化日志而不是 fmt.Printf
				fmt.Printf("worker %d: failed to process task %d: %v\n", id, taskID, err)
			}
		case <-s.quit:
			// 收到退出信号，处理完当前任务后退出
			// 最佳实践：退出前清空队列中的任务（可选，取决于业务需求）
			// 这里选择直接退出，让剩余任务在下次启动时处理或由其他机制处理
			return
		}
	}
}

// processTask 处理单个任务
// 要求：
// 1. 从数据库获取任务信息
// 2. 更新状态为 running
// 3. 根据任务类型执行不同逻辑（调用 executeTaskLogic）
// 4. 更新任务状态为 completed 或 failed，并记录结果
// 5. 使用事务确保数据一致性
func (s *TaskScheduler) processTask(taskID uint) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		var task Task
		// ⚠️ 使用 First 而不是 Find，First 在记录不存在时会返回 ErrRecordNotFound
		if err := tx.First(&task, taskID).Error; err != nil {
			return err
		}
		// ❌ 字段名应该使用小写的 "status" 而不是 "Status"（GORM 会自动处理，但保持一致性更好）
		// 同时更新 UpdatedAt 字段
		if err := tx.Model(&task).Update("status", "running").Error; err != nil {
			return err
		}
		result, err := s.executeTaskLogic(task.Type, task.Params)
		if err != nil {
			// ❌ 失败时应该记录错误信息到 Error 字段
			// 最佳实践：使用 Updates + 结构体 一次性更新多个字段，减少数据库操作
			return tx.Model(&task).Updates(Task{Status: "failed", Error: err.Error()}).Error
		}
		// ❌ 使用 Updates 一次性更新状态和结果
		return tx.Model(&task).Updates(Task{Status: "completed", Result: result}).Error
	})
}

// executeTaskLogic 执行任务逻辑
// 要求：
// 1. email: 模拟耗时 1-3 秒，成功率 90%
// 2. report: 模拟耗时 3-5 秒，成功率 80%
// 3. cleanup: 模拟耗时 0.5-1 秒，成功率 95%
// 4. 使用随机数模拟成功/失败
// 5. 返回执行结果字符串或错误
func (s *TaskScheduler) executeTaskLogic(taskType string, params string) (string, error) {
	_ = params
	secs := 0.0
	failRate := 0.0
	switch taskType {
	case "email":
		secs = rand.Float64()*2 + 1
		failRate = 0.9
	case "report":
		secs = rand.Float64()*2 + 3
		failRate = 0.8
	case "cleanup":
		secs = rand.Float64() / 2
		failRate = 0.95
	}
	dur := time.Duration(secs) * time.Second
	time.Sleep(dur)
	if rand.Float64() > failRate {
		return "", fmt.Errorf("task failed: %s", taskType)
	}
	return fmt.Sprintf("task done: %s", taskType), nil
}

// Shutdown 优雅关闭调度器
// 要求：
// 1. 关闭 quit Channel 通知所有 Worker 退出
// 2. 等待一小段时间让正在执行的任务完成
// 3. 关闭 taskQueue
func (s *TaskScheduler) Shutdown(ctx context.Context) error {
	// ⚠️ 最佳实践：

	// 第1步：标记调度器已关闭，阻止新任务提交
	// 这样可以确保 SubmitTask 在关闭过程中返回错误，而不是 panic
	atomic.StoreInt32(&s.closed, 1)

	// 第2步：关闭 taskQueue，阻止新任务进入队列
	// 必须先关闭队列，再通知 Worker 退出，否则等待期间仍可能有新任务提交
	close(s.taskQueue)

	// 第3步：关闭 quit channel，通知所有 Worker 退出
	close(s.quit)

	// 第4步：使用 context 控制等待时间，等待所有 Worker 完成
	// 开一个 goroutine 等待 waitGroup 归零，归零后设置 done
	done := make(chan struct{})
	go func() {
		s.wg.Wait()
		close(done)
	}()

	// 等待 done（所有 worker 退出）或超时
	select {
	case <-done:
		// 所有 Worker 已退出
		return nil
	case <-ctx.Done():
		// 超时，强制退出
		return ctx.Err()
	}
}
