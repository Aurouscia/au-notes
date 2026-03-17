package main

import (
	"fmt"
	"sync"
	"time"
)

// Task 表示一个待处理的任务
type Task struct {
	ID   int
	Name string
}

// Result 表示任务处理结果
type Result struct {
	TaskID int
	Output string
	Error  error
}

// ProcessTask 处理单个任务，返回结果
func ProcessTask(task Task) Result {
	time.Sleep(500 * time.Millisecond)
	output := fmt.Sprintf("Task %d: %s processed", task.ID, task.Name)
	return Result{
		TaskID: task.ID,
		Output: output,
		Error:  nil,
	}
}

// ============================================
// 最佳实践 1: WorkerPool - 真正的并发工作池
// ============================================

// WorkerPool 使用 n 个 worker 并发处理任务
func WorkerPool(tasks []Task, n int) []Result {
	taskCh := make(chan Task)
	resultCh := make(chan Result, len(tasks))

	// 启动 n 个 worker goroutine
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// ✅ 最佳实践：使用 for-range 遍历 channel
			// 当 channel 关闭时，range 会自动退出
			for task := range taskCh {
				fmt.Printf("[Worker %d] Processing: %s\n", workerID, task.Name)
				result := ProcessTask(task)
				resultCh <- result
			}
		}(i)
	}

	// 发送任务（在单独的 goroutine 中，避免阻塞）
	go func() {
		for _, task := range tasks {
			taskCh <- task
		}
		// ✅ 最佳实践：发送方负责关闭 channel
		close(taskCh)
	}()

	// 等待所有 worker 完成，然后关闭结果 channel
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 收集结果
	var results []Result
	// ✅ 最佳实践：使用 for-range 接收结果
	for result := range resultCh {
		results = append(results, result)
	}

	return results
}

// ============================================
// 最佳实践 2: ProducerConsumer - 多生产者多消费者
// ============================================

// ProducerConsumer 实现生产者-消费者模式
func ProducerConsumer(producerCount, consumerCount, taskCount int) []Result {
	taskCh := make(chan Task, 100)
	resultCh := make(chan Result, producerCount*taskCount)

	// ✅ 最佳实践：使用 WaitGroup 等待生产者完成
	var producerWg sync.WaitGroup
	for i := 0; i < producerCount; i++ {
		producerWg.Add(1)
		// ✅ 最佳实践：将循环变量作为参数传入，避免闭包陷阱
		go func(id int) {
			defer producerWg.Done()
			for j := 0; j < taskCount; j++ {
				taskID := id*10000 + j
				task := Task{
					ID:   taskID,
					Name: fmt.Sprintf("task-%d", taskID),
				}
				taskCh <- task
				fmt.Printf("[Producer %d] Created: %s\n", id, task.Name)
			}
		}(i)
	}

	// ✅ 最佳实践：在单独的 goroutine 中等待生产者完成并关闭 channel
	go func() {
		producerWg.Wait()
		close(taskCh)
		fmt.Println("[System] All producers done, taskCh closed")
	}()

	// 启动消费者
	var consumerWg sync.WaitGroup
	for i := 0; i < consumerCount; i++ {
		consumerWg.Add(1)
		go func(id int) {
			defer consumerWg.Done()
			// ✅ 最佳实践：使用 for-range 遍历 channel
			for task := range taskCh {
				fmt.Printf("[Consumer %d] Processing: %s\n", id, task.Name)
				result := ProcessTask(task)
				resultCh <- result
			}
			fmt.Printf("[Consumer %d] Exiting (taskCh closed)\n", id)
		}(i)
	}

	// 等待消费者完成并关闭结果 channel
	go func() {
		consumerWg.Wait()
		close(resultCh)
		fmt.Println("[System] All consumers done, resultCh closed")
	}()

	// 收集结果
	var results []Result
	for result := range resultCh {
		results = append(results, result)
	}

	return results
}

// ============================================
// 最佳实践 3: 带超时的任务处理
// ============================================

// ProcessWithTimeout 处理任务，带超时控制
func ProcessWithTimeout(task Task, timeout time.Duration) (Result, bool) {
	resultCh := make(chan Result, 1)

	go func() {
		resultCh <- ProcessTask(task)
	}()

	// ✅ 最佳实践：使用 select + time.After 实现超时
	select {
	case result := <-resultCh:
		return result, true // 成功
	case <-time.After(timeout):
		return Result{}, false // 超时
	}
}

// ============================================
// 主函数测试
// ============================================

func main() {
	fmt.Print("=== 并发任务处理器测试（最佳实践版本）===\n\n")

	// 1. 测试 WorkerPool
	fmt.Println("--- Test 1: WorkerPool ---")
	tasks := make([]Task, 10)
	for i := 0; i < 10; i++ {
		tasks[i] = Task{ID: i, Name: fmt.Sprintf("task-%d", i)}
	}
	results := WorkerPool(tasks, 3)
	fmt.Printf("WorkerPool completed: %d tasks processed\n\n", len(results))

	// 2. 测试 ProducerConsumer
	fmt.Println("--- Test 2: ProducerConsumer ---")
	results = ProducerConsumer(2, 3, 5)
	fmt.Printf("ProducerConsumer completed: %d tasks processed\n\n", len(results))

	// 3. 测试带超时的任务处理
	fmt.Println("--- Test 3: ProcessWithTimeout ---")

	// 3.1 正常情况（超时 1 秒，任务 500ms）
	result, ok := ProcessWithTimeout(Task{ID: 1, Name: "quick-task"}, time.Second)
	if ok {
		fmt.Printf("Task completed: %s\n", result.Output)
	} else {
		fmt.Println("Task timeout!")
	}

	// 3.2 超时情况（超时 200ms，任务 500ms）
	result, ok = ProcessWithTimeout(Task{ID: 2, Name: "slow-task"}, 200*time.Millisecond)
	if ok {
		fmt.Printf("Task completed: %s\n", result.Output)
	} else {
		fmt.Println("Task timeout!")
	}
}

// ============================================
// 关键知识点总结
// ============================================
//
// 1. Channel 遍历方式
//    ✅ for v := range ch     - 推荐，channel 关闭时自动退出
//    ✅ for { v, ok := <-ch; if !ok { break } } - 需要手动检查 ok
//    ❌ for v, ok := <-ch; ok; { } - 错误！ok 只在初始化时求值
//
// 2. Channel 关闭原则
//    - 只能由发送方关闭
//    - 关闭后不能再发送（panic）
//    - 关闭后可以继续接收已缓冲的数据
//    - range 会在 channel 关闭且无数据时自动退出
//
// 3. WaitGroup 使用模式
//    - Add() 要在启动 goroutine 之前调用
//    - Done() 在 goroutine 结束时调用（通常用 defer）
//    - Wait() 阻塞等待所有 goroutine 完成
//
// 4. 生产者-消费者关闭模式
//    - 启动生产者 goroutine，使用 WaitGroup 等待
//    - 在单独的 goroutine 中：Wait() -> close(channel)
//    - 消费者使用 for-range，channel 关闭后自动退出
//
// 5. 闭包陷阱
//    - 循环变量在 goroutine 中捕获要小心
//    - 推荐：将变量作为参数传入 go func(v Type) { ... }(v)
