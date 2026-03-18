# 实践题：可取消的任务队列

## 需求

实现一个支持优雅关闭的任务队列处理器，使用 `context` 实现取消信号传播：

1. **任务队列**：支持提交任务到队列
2. **Worker 池**：固定数量的 worker 并发处理任务
3. **优雅关闭**：收到取消信号后，等待正在处理的任务完成，但不再接受新任务
4. **超时控制**：单个任务处理超时时间为 2 秒

## 要求

1. 实现 `TaskQueue` 结构体：
   - `NewTaskQueue(workerCount int) *TaskQueue`：创建任务队列
   - `Submit(task Task) error`：提交任务（队列满时返回错误）
   - `Start(ctx context.Context)`：启动 worker 池
   - `Stop()`：优雅关闭

2. 实现 `Task` 结构体和处理逻辑：
   - 每个任务有 ID 和处理耗时
   - 使用 `select` 监听 `ctx.Done()` 实现取消
   - 超时返回错误

3. 使用 `context` 传递取消信号：
   - 主程序可以取消所有任务
   - 单个任务可以设置超时

## 数据结构

```go
type Task struct {
    ID       int
    Duration time.Duration // 模拟处理耗时
}

type TaskResult struct {
    TaskID int
    Output string
    Err    error
}

type TaskQueue struct {
    tasks   chan Task
    results chan TaskResult
    workers int
    wg      sync.WaitGroup
}
```

## 示例代码

```go
func main() {
    queue := NewTaskQueue(3) // 3 个 worker
    
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()
    
    // 启动 worker
    go queue.Start(ctx)
    
    // 提交任务
    for i := 1; i <= 10; i++ {
        task := Task{
            ID:       i,
            Duration: time.Duration(i%5+1) * time.Second,
        }
        if err := queue.Submit(task); err != nil {
            log.Printf("提交任务 %d 失败: %v", i, err)
        }
    }
    
    // 3 秒后取消所有任务
    time.Sleep(3 * time.Second)
    cancel()
    
    // 等待关闭
    queue.Stop()
}
```

## 提示

1. 使用 `context.WithTimeout` 为单个任务设置超时
2. 使用 `select` 监听 `ctx.Done()` 实现取消
3. 使用 `sync.WaitGroup` 等待所有 worker 退出
4. 注意处理 channel 的关闭顺序，避免死锁
