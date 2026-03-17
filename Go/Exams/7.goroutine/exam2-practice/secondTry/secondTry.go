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
// 模拟处理时间：500ms
// 输出格式："Task [ID]: [Name] processed"
func ProcessTask(task Task) Result {
	time.Sleep(500 * time.Millisecond)
	return Result{
		TaskID: task.ID,
		Output: fmt.Sprintf("Task %d: %s processed", task.ID, task.Name),
		Error:  nil,
	}
}

// WorkerPool 使用 n 个 worker 并发处理任务
// tasks: 待处理任务队列
// n: worker 数量
// 返回：所有任务的结果切片
func WorkerPool(tasks []Task, n int) []Result {
	taskCount := len(tasks)
	taskCh := make(chan Task, taskCount)
	resultCh := make(chan Result, taskCount) // 使用 channel 收集多个 routine 的结果（优于共享切片）

	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				// 无限循环，一直接任务，直到 taskCh 关闭（可以改为使用 for-range 更好看）
				task, ok := <-taskCh
				if !ok {
					break // 没有任务了，才跳出无限循环
				}
				fmt.Printf("worker %d started task: %s\n", i, task.Name)
				resultCh <- ProcessTask(task)
				fmt.Printf("worker %d finished task: %s\n", i, task.Name)
			}
		}(i) // routine 中依赖循环变量：应该使用 func 的参数传入。如果直接使用 i 会读取到错误的值（循环内仅捕获了 i 的引用）
	}

	// 在单独的 routine 中发送任务，避免阻塞
	go func() {
		for _, t := range tasks {
			taskCh <- t
		}
		close(taskCh) // 全部发送完成后，由发送方关闭
	}()

	go func() {
		wg.Wait() // 全部 worker 处理完成后，关闭 channel（适合在单独 routine 处理避免阻塞）
		close(resultCh)
		// 由于发送方分散在多个 routine 中，所以应该用 WaitGroup 等待全部结束后统一关闭）
	}()

	// wg.Wait() // 没有必要等待，可以由 range 来逐个等待
	results := make([]Result, 0, taskCount)
	for r := range resultCh {
		results = append(results, r)
	}
	return results
}

// ProducerConsumer 实现生产者-消费者模式
// producerCount: 生产者数量
// consumerCount: 消费者数量
// taskCount: 每个生产者产生的任务数
// 返回：所有消费的结果
func ProducerConsumer(producerCount, consumerCount, taskCount int) []Result {
	taskCh := make(chan Task, 5)
	resultCh := make(chan Result, 5) // 使用 channel 收集多个 routine 的结果（优于共享切片）

	produceWg := sync.WaitGroup{} // 使用 waitGroup 确保所有生产者都结束发送后，统一关闭 channel
	for i := 0; i < producerCount; i++ {
		produceWg.Add(1)
		go func(i int) {
			defer produceWg.Done()
			for j := 0; j < taskCount; j++ {
				taskID := (i+1)*10000 + j
				taskName := "task-" + fmt.Sprint(taskID)
				taskCh <- Task{ID: taskID, Name: taskName}
				fmt.Printf("producer %d created task: %s\n", i, taskName)
			}
		}(i)
	}

	go func() {
		produceWg.Wait()
		close(taskCh)
	}()

	resultWg := sync.WaitGroup{}
	for i := 0; i < consumerCount; i++ {
		resultWg.Add(1)
		go func(i int) {
			defer resultWg.Done()
			for {
				// 无限循环，一直读取 taskCh 直到其关闭（使用 for-range 更好）
				task, ok := <-taskCh
				if !ok {
					break
				}
				fmt.Printf("consumer %d started task: %s\n", i, task.Name)
				resultCh <- ProcessTask(task)
				fmt.Printf("consumer %d finished task: %s\n", i, task.Name)
			}
		}(i)
	}

	go func() {
		resultWg.Wait()
		close(resultCh)
	}()

	results := make([]Result, 0, producerCount*taskCount)
	for r := range resultCh {
		results = append(results, r)
	}
	return results
}

func ProcessWithTimeout(task Task, timeout time.Duration) (Result, bool) {
	// resultCh := make(chan Result)
	// select {
	// case resultCh <- ProcessTask(task):
	// 	return <-resultCh, true
	// case <-time.After(timeout):
	// 	return Result{}, false
	// }

	// resultCh := make(chan Result) // ❌ 无缓冲区：如果 timeout，resultCh<- 会永久卡住（goroutine 泄漏）
	// 正确写法：
	resultCh := make(chan Result, 1) // 即使 timeout，也会让 goroutine 最终完成（resultCh 会自动回收，无需手动 close）

	go func() {
		resultCh <- ProcessTask(task)
	}()

	select {
	case result := <-resultCh:
		return result, true
	case <-time.After(timeout):
		return Result{}, false
	}
}

func main() {
	fmt.Print("=== 并发任务处理器测试（第二次尝试）===\n\n")

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
