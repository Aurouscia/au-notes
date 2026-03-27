package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
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

	quit := false

	// 开一个 routine，监听 s.quit 的关闭信号
	go func() {
		for {
			_, ok := <-s.quit
			if !ok {
				quit = true
				break
			}
		}
	}()

Loop:
	for {
		taskId, ok := <-s.taskQueue
		if !ok {
			break Loop
		}
		err := s.processTask(taskId)
		if err != nil {
			fmt.Printf("worker %d encountered error when processing task %d\n", id, taskId)
		}
		if quit {
			break Loop
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
		if err := tx.Find(&task, taskID).Error; err != nil {
			return err
		}
		err := tx.Model(&task).Update("Status", "running").Error
		if err != nil {
			return err
		}
		result, err := s.executeTaskLogic(task.Type, task.Params)
		if err != nil {
			_ = tx.Model(&task).Update("Status", "failed")
			return err
		}
		// ❌ Update 会重置当前链式状态，多个 Update 不会合并为一个操作，要更新多个字段必须使用 Updates
		// tx.Model(&task).Update("Status", "completed").Update("Result", result)
		err = tx.Model(&task).Updates(Task{Status: "completed", Result: result}).Error
		return err
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
	close(s.quit)
	s.wg.Wait()
	close(s.taskQueue)
	return nil
}
