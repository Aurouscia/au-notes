package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

// TaskScheduler 任务调度器
type TaskScheduler struct {
	db          *gorm.DB
	taskQueue   chan uint      // 任务 ID 队列
	quit        chan struct{}  // 退出信号
	workerCount int            // Worker 数量
}

// NewTaskScheduler 创建任务调度器
// TODO: 初始化 Channel 和 Worker
// 要求：
// 1. taskQueue 缓冲区大小设置为 100
// 2. 启动指定数量的 Worker Goroutine
// 3. 返回初始化好的 TaskScheduler
func NewTaskScheduler(db *gorm.DB, workerCount int) *TaskScheduler {
	// 请在此处实现
	return nil
}

// SubmitTask 提交任务到队列
// TODO: 将任务 ID 发送到 Channel
// 注意：使用 select 防止 Channel 满时阻塞，如果 Channel 已满返回错误
func (s *TaskScheduler) SubmitTask(taskID uint) error {
	// 请在此处实现
	return nil
}

// worker 工作协程
// TODO: 从 Channel 中获取任务 ID 并执行
// 要求：
// 1. 监听 taskQueue 和 quit 信号
// 2. 收到 quit 信号时优雅退出
// 3. 调用 processTask 处理任务
func (s *TaskScheduler) worker(id int) {
	// 请在此处实现
}

// processTask 处理单个任务
// TODO: 实现任务执行逻辑
// 要求：
// 1. 从数据库获取任务信息
// 2. 更新状态为 running
// 3. 根据任务类型执行不同逻辑（调用 executeTaskLogic）
// 4. 更新任务状态为 completed 或 failed，并记录结果
// 5. 使用事务确保数据一致性
func (s *TaskScheduler) processTask(taskID uint) {
	// 请在此处实现
}

// executeTaskLogic 执行任务逻辑
// TODO: 根据任务类型执行不同的模拟逻辑
// 要求：
// 1. email: 模拟耗时 1-3 秒，成功率 90%
// 2. report: 模拟耗时 3-5 秒，成功率 80%
// 3. cleanup: 模拟耗时 0.5-1 秒，成功率 95%
// 4. 使用随机数模拟成功/失败
// 5. 返回执行结果字符串或错误
func (s *TaskScheduler) executeTaskLogic(taskType string, params string) (string, error) {
	// 请在此处实现
	return "", nil
}

// Shutdown 优雅关闭调度器
// TODO: 发送退出信号并等待 Worker 完成
// 要求：
// 1. 关闭 quit Channel 通知所有 Worker 退出
// 2. 等待一小段时间让正在执行的任务完成
// 3. 关闭 taskQueue
func (s *TaskScheduler) Shutdown(ctx context.Context) error {
	// 请在此处实现
	return nil
}

// randomSleep 随机睡眠（辅助函数）
func randomSleep(min, max time.Duration) {
	time.Sleep(min + time.Duration(rand.Float64()*float64(max-min)))
}
