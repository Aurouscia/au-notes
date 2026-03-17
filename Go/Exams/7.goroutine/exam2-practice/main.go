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
	time.Sleep(500)
	output := fmt.Sprintf("TASK %d: %s processed", task.ID, task.Name)
	return Result{
		TaskID: task.ID,
		Output: output,
		Error:  nil,
	}
}

// WorkerPool 使用 n 个 worker 并发处理任务
// tasks: 待处理任务队列
// n: worker 数量
// 返回：所有任务的结果切片
func WorkerPool(tasks []Task, n int) []Result {
	c := make(chan Task, n)
	go func() {
		defer close(c)
		for _, t := range tasks {
			fmt.Println("task queued: ", t.Name)
			c <- t
		}
	}()
	results := []Result{}
	for t, ok := <-c; ok; {
		fmt.Println("task taken: ", t.Name)
		result := ProcessTask(t)
		results = append(results, result)
	}
	return results
}

// ProducerConsumer 实现生产者-消费者模式
// producerCount: 生产者数量
// consumerCount: 消费者数量
// taskCount: 每个生产者产生的任务数
// 返回：所有消费的结果
func ProducerConsumer(producerCount, consumerCount, taskCount int) []Result {
	c := make(chan Task, 100)
	for i := 0; i < producerCount; i++ {
		producerId := i // 这个得在 go 外面捕获，在里面捕获可能出问题（gopls警告）
		go func() {
			for j := 0; j < taskCount; j++ {
				taskId := producerId*10000 + j
				taskName := "task-" + fmt.Sprint(taskId)
				c <- Task{
					ID:   taskId,
					Name: taskName,
				}
			}
		}()
	}
	close(c)
	results := []Result{}
	wg := sync.WaitGroup{}
	for i := 0; i < consumerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t, ok := <-c; ok; {
				res := ProcessTask(t)
				results = append(results, res)
			}
		}()
	}
	wg.Wait()
	return results
}

func main() {
	fmt.Println("=== 并发任务处理器测试 ===")

	// 1. 测试 WorkerPool
	//    - 创建 10 个任务
	//    - 使用 3 个 worker 并发处理
	//    - 打印所有结果

	tasks := []Task{}
	for i := 0; i < 10; i++ {
		taskName := "task-" + fmt.Sprint(i)
		tasks = append(tasks, Task{
			ID:   i,
			Name: taskName,
		})
	}
	results := WorkerPool(tasks, 3)
	for _, r := range results {
		fmt.Printf("taskID: %d: %s\n", r.TaskID, r.Output)
	}

	// 2. 测试 ProducerConsumer
	//    - 2 个生产者，每个生产 5 个任务
	//    - 3 个消费者并发处理
	//    - 打印所有结果

	results = ProducerConsumer(2, 3, 5)
	for _, r := range results {
		fmt.Printf("taskID: %d: %s\n", r.TaskID, r.Output)
	}

	// 3. 测试带超时的任务处理
	//    - 使用 select + time.After 实现超时控制
	//    - 如果任务 1 秒内未完成，输出 "timeout"
	ch := make(chan Result, 5)
	go func() {
		ch <- ProcessTask(Task{Name: "some task 1", ID: 1})
		ch <- ProcessTask(Task{Name: "some task 2", ID: 2})
		ch <- ProcessTask(Task{Name: "some task 3", ID: 3})
	}()
	select {
	case result := <-ch:
		fmt.Println(result)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
