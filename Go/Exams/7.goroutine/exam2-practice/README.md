# Goroutine 与 Channel - 实践题

## 需求

实现一个简单的**并发任务处理器**，使用 goroutine 和 channel 来并行处理多个任务。

### 1. 定义任务类型

```go
// Task 表示一个待处理的任务
type Task struct {
    ID   int
    Name string
    // 其他字段...
}

// Result 表示任务处理结果
type Result struct {
    TaskID int
    Output string
    Error  error
}
```

### 2. 实现任务处理器

实现 `ProcessTask` 函数，模拟任务处理（使用 `time.Sleep` 模拟耗时）：

```go
// ProcessTask 处理单个任务，返回结果
// 模拟处理时间：500ms
// 输出格式："Task [ID]: [Name] processed"
func ProcessTask(task Task) Result
```

### 3. 实现并发工作池

实现 `WorkerPool` 函数，使用固定数量的 worker goroutine 并发处理任务：

```go
// WorkerPool 使用 n 个 worker 并发处理任务
// tasks: 待处理任务队列
// n: worker 数量
// 返回：所有任务的结果切片
func WorkerPool(tasks []Task, n int) []Result
```

**要求：**
- 创建 n 个 worker goroutine
- 使用 channel 分发任务
- 使用 channel 收集结果
- 等待所有任务完成后返回结果

### 4. 实现生产者-消费者模式

实现 `ProducerConsumer` 函数：

```go
// ProducerConsumer 实现生产者-消费者模式
// producerCount: 生产者数量
// consumerCount: 消费者数量
// taskCount: 每个生产者产生的任务数
// 返回：所有消费的结果
func ProducerConsumer(producerCount, consumerCount, taskCount int) []Result
```

**要求：**
- 多个生产者 goroutine 向同一个 channel 发送任务
- 多个消费者 goroutine 从 channel 接收并处理任务
- 使用 `sync.WaitGroup` 等待所有生产者完成
- 正确关闭 channel，让消费者能够退出

### 5. 主函数测试

在 `main()` 中完成以下测试：

```go
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
```

---

## 运行要求

```bash
go run main.go
```

---

## 提示

- 使用 `make(chan Task)` 创建任务 channel
- 使用 `make(chan Result, bufferSize)` 创建带缓冲的结果 channel
- 使用 `sync.WaitGroup` 等待 goroutine 完成
- 生产者负责关闭 channel：`close(taskCh)`
- 使用 `for task := range taskCh` 循环接收任务
- 超时控制示例：
  ```go
  select {
  case result := <-resultCh:
      fmt.Println(result)
  case <-time.After(time.Second):
      fmt.Println("timeout")
  }
  ```

---

## 核心考察点

1. **Goroutine 创建与管理** - 正确启动和协调多个 goroutine
2. **Channel 通信** - 使用 channel 传递数据和同步
3. **Channel 关闭** - 生产者关闭，消费者检测关闭退出
4. **WaitGroup 使用** - 等待一组 goroutine 完成
5. **Select 超时控制** - 使用 select 实现非阻塞或超时操作
