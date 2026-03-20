package main

import (
	"context"
	"fmt"
	"log"
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
		tasks:   make(chan Task, 5),
		results: make(chan TaskResult, 5),
		workers: workerCount,
		wg:      sync.WaitGroup{},
	}
}

// Submit 提交任务 // ⚠️ 不符合惯例：一般把 context 放第一个参数
func (q *TaskQueue) Submit(task Task, ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		// 如果ctx已取消：拒绝添加，返回错误
		return err
	}
	select {
	case q.tasks <- task:
	// case <-ctx.Done(): // 这样不好：当“tasks可接收”和“ctx.Done”同时就绪时，可能仍会加入
	// 	return ctx.Err()
	default:
		// 如果队列满，拒绝添加，返回错误
		return fmt.Errorf("queue full")
	}
	return nil
}

// worker 处理任务
func (q *TaskQueue) worker(ctx context.Context, id int) {
	// 1. 从 tasks channel 获取任务
	// 2. 使用 select 监听 ctx.Done() 实现取消
	// 3. 模拟处理（使用 time.After 或 sleep）
	// 4. 发送结果到 results channel
	// 5. 注意：当 ctx 被取消或 tasks 关闭时，优雅退出

	defer q.wg.Done() // 不管怎么退出，都把 waitGroup 减 1

	// 用 drain 函数处理“ctx已经取消，但channel仍有剩下的任务”这种情况的处理逻辑
	// 读取剩余的 task，全部转化为失败的 result（活要见人死要见尸）
	// ❌ 设计问题：最好是做成有超时的（卡住太久则放弃写入结果）避免外部原因阻塞
	drain := func(err error) {
		for t := range q.tasks {
			q.results <- TaskResult{
				TaskID: t.ID,
				Err:    err,
			}
		}
	}

	for {
		select {
		case t, ok := <-q.tasks:
			if !ok {
				return // channel 已关闭且没有任务了，直接返回
			}
			if err := ctx.Err(); err != nil {
				// ❌ 忘记处理当前已取出的任务 t，造成一个任务丢失！
				q.results <- TaskResult{
					TaskID: t.ID,
					Err:    err,
				}
				drain(err) // 已经cancel，将任务channel drain掉
				return
			}
			q.results <- processWithCtx(t, ctx, id)
		case <-ctx.Done():
			drain(ctx.Err())
			return
		}
	}
}

// 带有 context 的实际处理函数 // ⚠️ 不符合惯例：一般把 context 放第一个参数
func processWithCtx(task Task, ctx context.Context, id int) TaskResult {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*2001) // 每个任务限时 2 秒
	defer cancel()
	select {
	// 用 time.After 模拟实际任务完成
	case <-time.After(task.Duration):
		return TaskResult{
			TaskID: task.ID,
			Output: fmt.Sprintf("output from worker %d", id),
			Err:    nil,
		}
	// 如果在任务完成前，ctx 被取消：直接返回（time.After 会被 GC 收集的）
	case <-ctx.Done():
		return TaskResult{
			TaskID: task.ID,
			Err:    ctx.Err(),
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

// Stop 优雅关闭 // ⚠️ 会等待所有 workder 结束，更合适的命名是 Shutdown 或 Close
func (q *TaskQueue) Stop() {
	// 1. 关闭 tasks channel（通知 worker 不再接受新任务）
	// 2. 等待所有 worker 完成
	// 3. 关闭 results channel
	close(q.tasks)
	q.wg.Wait() // wg 返回 0，代表所有 worker 都退出了，不会有 results 的发送
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

	// 启动 worker（非阻塞性操作，无需 go）
	queue.Start(ctx)

	// 收集结果的 goroutine
	collectWg := sync.WaitGroup{}
	collectWg.Add(1)
	go func() {
		for result := range queue.Results() {
			if result.Err != nil {
				log.Printf("任务 %d 失败: %v", result.TaskID, result.Err)
			} else {
				log.Printf("任务 %d 成功: %s", result.TaskID, result.Output)
			}
		}
		collectWg.Done()
	}()

	// 提交任务
	fmt.Println("提交任务...")
	for i := 1; i <= 10; i++ {
		secs := i%4 + 1
		task := Task{
			ID:       i,
			Duration: time.Duration(secs) * time.Second,
		}
		if err := queue.Submit(task, ctx); err != nil {
			log.Printf("提交任务 %d 失败: %v", i, err)
		} else {
			fmt.Printf("任务 %d 已提交，预计耗时 %ds \n", i, secs)
		}
	}

	// 3 秒后取消所有任务
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("\n取消所有任务...")
	cancel()

	// 等待关闭
	queue.Stop()
	collectWg.Wait() // 等待结果全部读取
	fmt.Println("任务队列已关闭")
}
