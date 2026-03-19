package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// Task 表示一个任务
type Task struct {
	ID       int
	Duration time.Duration // 模拟处理耗时
}

// TaskResult 表示任务处理结果
type TaskResult struct {
	TaskID int
	Output string
	Err    error
}

// TaskQueue 任务队列
type TaskQueue struct {
	tasks   chan Task
	results chan TaskResult
	workers int
	wg      sync.WaitGroup
}

// NewTaskQueue 创建任务队列
func NewTaskQueue(workerCount int) *TaskQueue {
	return &TaskQueue{
		tasks:   make(chan Task, 10),
		results: make(chan TaskResult, 10),
		workers: workerCount,
		wg:      sync.WaitGroup{},
	}
}

// Submit 提交任务
func (q *TaskQueue) Submit(task Task) error {
	// 如果队列满，返回错误
	if len(q.tasks) == cap(q.tasks) {
		return fmt.Errorf("queue full")
	}
	q.tasks <- task
	return nil
}

// worker 处理任务
func (q *TaskQueue) worker(ctx context.Context, id int) {
	// 1. 从 tasks channel 获取任务
	// 2. 使用 select 监听 ctx.Done() 实现取消
	// 3. 模拟处理（使用 time.After 或 sleep）
	// 4. 发送结果到 results channel
	// 5. 注意：当 ctx 被取消或 tasks 关闭时，优雅退出

	defer q.wg.Done()
	for {
		workingOn := Task{}
		select {
		case t, ok := <-q.tasks:
			if !ok {
				return
			}
			workingOn = t
		case <-ctx.Done():
			return
		}

		// 模拟处理任务，1-3秒随机延迟
		sec := rand.Int31n(3) + 1
		dur := time.Duration(sec) * time.Second
		time.Sleep(dur)
		q.results <- TaskResult{
			TaskID: workingOn.ID,
			Output: fmt.Sprintf("output from worker %d", id),
			Err:    nil,
		}
	}
}

// Start 启动 worker 池
func (q *TaskQueue) Start(ctx context.Context) {
	// 启动指定数量的 worker goroutine
	for i := 0; i < q.workers; i++ {
		q.wg.Add(1)
		go q.worker(ctx, i)
	}
}

// Stop 优雅关闭
func (q *TaskQueue) Stop() {
	// 1. 关闭 tasks channel（通知 worker 不再接受新任务）
	// 2. 等待所有 worker 完成
	// 3. 关闭 results channel
	close(q.tasks)
	q.wg.Wait()
	close(q.results)
}

// Results 获取结果 channel
func (q *TaskQueue) Results() <-chan TaskResult {
	return q.results
}

func main() {
	queue := NewTaskQueue(3) // 3 个 worker

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动 worker
	go queue.Start(ctx)

	// 收集结果的 goroutine
	go func() {
		for result := range queue.Results() {
			if result.Err != nil {
				log.Printf("任务 %d 失败: %v", result.TaskID, result.Err)
			} else {
				log.Printf("任务 %d 成功: %s", result.TaskID, result.Output)
			}
		}
	}()

	// 提交任务
	fmt.Println("提交任务...")
	for i := 1; i <= 10; i++ {
		task := Task{
			ID:       i,
			Duration: time.Duration(i%5+1) * time.Second,
		}
		if err := queue.Submit(task); err != nil {
			log.Printf("提交任务 %d 失败: %v", i, err)
		} else {
			fmt.Printf("任务 %d 已提交\n", i)
		}
	}

	// 3 秒后取消所有任务
	time.Sleep(3 * time.Second)
	fmt.Println("\n取消所有任务...")
	cancel()

	// 等待关闭
	queue.Stop()
	fmt.Println("任务队列已关闭")
}
