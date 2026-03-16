package main

import (
	"fmt"
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
	// 请在此实现
	return Result{}
}

// WorkerPool 使用 n 个 worker 并发处理任务
// tasks: 待处理任务队列
// n: worker 数量
// 返回：所有任务的结果切片
func WorkerPool(tasks []Task, n int) []Result {
	// 请在此实现
	return nil
}

// ProducerConsumer 实现生产者-消费者模式
// producerCount: 生产者数量
// consumerCount: 消费者数量
// taskCount: 每个生产者产生的任务数
// 返回：所有消费的结果
func ProducerConsumer(producerCount, consumerCount, taskCount int) []Result {
	// 请在此实现
	return nil
}

func main() {
	fmt.Println("=== 并发任务处理器测试 ===")

	// 1. 测试 WorkerPool
	//    - 创建 10 个任务
	//    - 使用 3 个 worker 并发处理
	//    - 打印所有结果

	// 2. 测试 ProducerConsumer
	//    - 2 个生产者，每个生产 5 个任务
	//    - 3 个消费者并发处理
	//    - 打印所有结果

	// 3. 测试带超时的任务处理
	//    - 使用 select + time.After 实现超时控制
	//    - 如果任务 1 秒内未完成，输出 "timeout"
}
