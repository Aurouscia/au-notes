# Goroutine 进阶：Select、WaitGroup 与 Context

## 1. Select 多路复用

`select` 语句用于在多个 channel 操作中进行选择，类似于 `switch`，但专门用于 channel。

### 基本语法

```go
select {
case v1 := <-ch1:      // 从 ch1 接收
    fmt.Println("ch1:", v1)
case v2 := <-ch2:      // 从 ch2 接收
    fmt.Println("ch2:", v2)
case ch3 <- 100:       // 向 ch3 发送
    fmt.Println("sent to ch3")
default:               // 默认分支（可选）
    fmt.Println("no channel ready")
}
```

### 特性

| 特性 | 说明 |
|------|------|
| 阻塞等待 | 没有 `default` 时，select 阻塞直到某个 case 可执行 |
| 随机选择 | 多个 case 同时就绪时，**随机**选择一个执行 |
| 非阻塞 | 有 `default` 时，若无可执行的 case，执行 default |

### 示例：多数据源合并

```go
package main

import (
    "fmt"
    "time"
)

func producer1(ch chan<- string) {
    for i := 0; i < 3; i++ {
        ch <- fmt.Sprintf("producer1: %d", i)
        time.Sleep(100 * time.Millisecond)
    }
    close(ch)
}

func producer2(ch chan<- string) {
    for i := 0; i < 3; i++ {
        ch <- fmt.Sprintf("producer2: %d", i)
        time.Sleep(150 * time.Millisecond)
    }
    close(ch)
}

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go producer1(ch1)
    go producer2(ch2)
    
    // 使用 select 接收两个 channel 的数据
    for ch1 != nil || ch2 != nil {
        select {
        case v, ok := <-ch1:
            if !ok {
                ch1 = nil  // channel 关闭后置为 nil，不再选择
                continue
            }
            fmt.Println(v)
        case v, ok := <-ch2:
            if !ok {
                ch2 = nil
                continue
            }
            fmt.Println(v)
        }
    }
}
```

### 超时控制

```go
package main

import (
    "fmt"
    "time"
)

func slowOperation(ch chan<- string) {
    time.Sleep(2 * time.Second)
    ch <- "result"
}

func main() {
    ch := make(chan string)
    go slowOperation(ch)
    
    select {
    case result := <-ch:
        fmt.Println(result)
    case <-time.After(1 * time.Second):  // 超时 1 秒
        fmt.Println("timeout!")
    }
}
```

### 非阻塞操作

```go
ch := make(chan int, 1)

// 非阻塞发送
select {
case ch <- 42:
    fmt.Println("sent successfully")
default:
    fmt.Println("channel full, skip")
}

// 非阻塞接收
select {
case v := <-ch:
    fmt.Println("received:", v)
default:
    fmt.Println("no data available")
}
```

---

## 2. sync.WaitGroup

`WaitGroup` 用于等待**一组 goroutine** 全部完成。

### 核心方法

| 方法 | 作用 |
|------|------|
| `Add(n int)` | 添加 n 个待等待的任务计数 |
| `Done()` | 任务完成，计数减 1 |
| `Wait()` | 阻塞，直到计数为 0 |

### 基本示例

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // 确保在函数退出时调用 Done()
    
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)        // 为每个 worker 添加计数
        go worker(i, &wg)
    }
    
    wg.Wait()  // 等待所有 worker 完成
    fmt.Println("All workers finished")
}
```

### ⚠️ 常见错误

```go
// ❌ 错误：在 goroutine 内部 Add
var wg sync.WaitGroup
for i := 0; i < 3; i++ {
    go func() {
        wg.Add(1)      // 可能 Wait() 时还没执行到 Add
        defer wg.Done()
        // ...
    }()
}
wg.Wait()  // 可能提前返回！

// ✅ 正确：在启动 goroutine 前 Add
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        // ...
    }()
}
wg.Wait()
```

```go
// ❌ 错误：传递 WaitGroup 值（复制）
go worker(&wg)  // 正确：传递指针

func worker(wg sync.WaitGroup) {  // 错误：值拷贝，Done() 不影响原 wg
    defer wg.Done()
}
```

---

## 3. Context 上下文

`context` 包用于在 goroutine 之间传递**取消信号、超时、截止时间**。

### 核心类型

| 类型 | 说明 |
|------|------|
| `context.Context` | 接口类型，可获取取消信号、截止时间、键值对 |
| `context.CancelFunc` | 取消函数，调用后传播取消信号 |

### 创建 Context

```go
// 1. 根 context（通常作为函数参数传入）
ctx := context.Background()

// 2. 带取消的 context
ctx, cancel := context.WithCancel(parent)
defer cancel()  // 确保调用，防止 goroutine 泄漏

// 3. 带超时的 context
ctx, cancel := context.WithTimeout(parent, 2*time.Second)
defer cancel()

// 4. 带截止时间的 context
ctx, cancel := context.WithDeadline(parent, time.Now().Add(2*time.Second))
defer cancel()

// 5. 带值的 context
ctx := context.WithValue(parent, key, value)
```

### Context 的树状结构

Context 在实际使用中一般**链式创建**，形成**树状结构**：

```
Background()
    └── WithTimeout(5s)  // HTTP 请求 context
            ├── WithCancel()  // 业务逻辑 A
            │       └── WithValue()  // 传递用户 ID
            └── WithTimeout(2s)  // 业务逻辑 B（更短的超时）
```

```go
// 实际使用示例：HTTP 处理器中链式创建
func handler(w http.ResponseWriter, r *http.Request) {
    // r.Context() 是 HTTP 框架传入的 context（可能已有超时设置）
    ctx := r.Context()
    
    // 基于传入的 context 创建子 context，添加更短的超时
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()
    
    // 继续传递 ctx 给下游函数...
    result, err := processOrder(ctx, orderID)
    // ...
}

func processOrder(ctx context.Context, orderID string) (*Order, error) {
    // 基于传入的 ctx 再创建子 context，用于数据库操作
    ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
    defer cancel()
    
    return db.QueryOrder(ctx, orderID)
}
```

**关键点**：
- 子 context 继承父 context 的取消信号和截止时间
- 父 context 取消时，**所有子孙 context** 都会收到取消信号
- 子 context 的超时/取消**不会**影响父 context

### 取消信号传播

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func worker(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():  // 监听取消信号
            fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Worker %d working...\n", id)
            time.Sleep(300 * time.Millisecond)
        }
    }
}

func main() {
    ctx, cancel := context.WithCancel(context.Background())
    
    for i := 1; i <= 3; i++ {
        go worker(ctx, i)
    }
    
    time.Sleep(1 * time.Second)
    fmt.Println("Cancelling all workers...")
    cancel()  // 取消所有子 goroutine
    
    time.Sleep(500 * time.Millisecond)  // 等待 goroutine 退出
}
```

### 超时控制

```go
package main

import (
    "context"
    "fmt"
    "time"
)

func slowQuery(ctx context.Context) (string, error) {
    select {
    case <-time.After(2 * time.Second):
        return "result", nil
    case <-ctx.Done():
        return "", ctx.Err()  // context.DeadlineExceeded
    }
}

func main() {
    // 设置 1 秒超时
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    
    result, err := slowQuery(ctx)
    if err != nil {
        fmt.Println("Error:", err)  // Error: context deadline exceeded
        return
    }
    fmt.Println(result)
}
```

### 传递请求值（谨慎使用）

```go
// 定义 key 类型，防止冲突
type key string

const userIDKey key = "userID"

func handler(ctx context.Context) {
    // 存储值
    ctx = context.WithValue(ctx, userIDKey, "12345")
    
    // 获取值
    if userID, ok := ctx.Value(userIDKey).(string); ok {
        fmt.Println("UserID:", userID)
    }
}
```

> **注意**：Context 值用于传递请求作用域的数据（如 traceID、userID），不要用于传递可选参数。

### 惯例

函数如果有 context 参数，应该是第一个参数

---

## 4. 综合示例：任务队列

结合 WaitGroup、Channel 和 Context 实现一个可取消的并发任务处理器：

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Task 表示一个任务
type Task struct {
    ID   int
    Data string
}

// Result 表示任务结果
type Result struct {
    TaskID int
    Output string
    Err    error
}

// worker 处理任务
func worker(ctx context.Context, id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for {
        select {
        case task, ok := <-tasks:
            if !ok {
                fmt.Printf("Worker %d: task channel closed\n", id)
                return
            }
            
            // 模拟处理
            select {
            case <-time.After(500 * time.Millisecond):
                results <- Result{
                    TaskID: task.ID,
                    Output: fmt.Sprintf("Processed: %s", task.Data),
                }
            case <-ctx.Done():
                results <- Result{
                    TaskID: task.ID,
                    Err:    ctx.Err(),
                }
                return
            }
            
        case <-ctx.Done():
            fmt.Printf("Worker %d: cancelled\n", id)
            return
        }
    }
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    tasks := make(chan Task, 10)
    results := make(chan Result, 10)
    
    var wg sync.WaitGroup
    
    // 启动 3 个 worker
    for i := 1; i <= 3; i++ {
        wg.Add(1)
        go worker(ctx, i, tasks, results, &wg)
    }
    
    // 发送任务
    go func() {
        for i := 1; i <= 10; i++ {
            select {
            case tasks <- Task{ID: i, Data: fmt.Sprintf("task-%d", i)}:
                fmt.Printf("Sent task %d\n", i)
            case <-ctx.Done():
                fmt.Println("Task sending cancelled")
                close(tasks)
                return
            }
        }
        close(tasks)
    }()
    
    // 等待所有 worker 完成，然后关闭 results channel
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // 收集结果
    for r := range results {
        if r.Err != nil {
            fmt.Printf("Task %d failed: %v\n", r.TaskID, r.Err)
        } else {
            fmt.Printf("Task %d success: %s\n", r.TaskID, r.Output)
        }
    }
    
    fmt.Println("All done")
}
```

---

## 5. 核心要点总结

| 概念 | 用途 | 关键 API |
|------|------|----------|
| `select` | 多 channel 操作多路复用 | `case <-ch:`, `default` |
| `WaitGroup` | 等待多个 goroutine 完成 | `Add()`, `Done()`, `Wait()` |
| `Context` | 传递取消信号、超时、截止时间和值 | `WithCancel()`, `WithTimeout()`, `WithValue()` |

### 选择指南

| 场景 | 推荐方案 |
|------|----------|
| 等待单个 goroutine 完成 | Channel |
| 等待多个 goroutine 完成 | `sync.WaitGroup` |
| 多 channel 操作 | `select` |
| 超时控制 | `context.WithTimeout` + `select` |
| 主动取消 goroutine | `context.WithCancel` |
| 传递请求元数据 | `context.WithValue` |

### 不调用 cancel() 的后果

#### 1. Goroutine 泄漏

```go
// ❌ 错误示例：忘记 cancel 导致泄漏
func doSomething() {
    ctx, cancel := context.WithCancel(context.Background())
    // defer cancel()  // 忘记调用！
    
    go worker(ctx)  // worker 会一直阻塞等待 ctx.Done()
    // 函数返回，但 worker goroutine 永远不会退出
}
```

**泄漏原理**：

```
WithCancel 内部实现简图：

parentCtx ──┬──► childCtx (返回的 ctx)
            │
            └──► goroutine: 监听 parent 的 Done()
                     │
                     ▼
                当 parent 取消时，关闭 child 的 done channel
                当 cancel() 被调用时，关闭 child 的 done channel
```

- `WithCancel`/`WithTimeout`/`WithDeadline` 内部会创建一个新的 goroutine
- 这个 goroutine 负责监听父 context 的取消信号，并传播给子 context
- **如果不调用 `cancel()`，这个内部 goroutine 会一直阻塞等待，无法被 GC 回收**

#### 2. 定时器资源泄漏

```go
// ❌ 错误示例：即使超时自动触发，也需要 cancel
func query() {
    ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
    // 超时后返回，但内部的 time.Timer 资源未释放！
    
    db.Query(ctx, sql)
}  // 函数返回，但定时器还在后台运行
```

**原理**：`WithTimeout` 内部使用 `time.AfterFunc` 创建定时器，需要调用 `cancel()` 来停止定时器，释放资源。

#### 3. 正确的做法

```go
// ✅ 正确示例：使用 defer 确保 cancel 被调用
func doSomething() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()  // 确保调用，防止泄漏
    
    go worker(ctx)
    // ... 其他逻辑
}  // 函数返回时自动调用 cancel()

// ✅ 对于可能提前返回的函数，defer 尤其重要
func process(ctx context.Context) error {
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()  // 无论哪个分支返回，都能执行到
    
    if err := step1(ctx); err != nil {
        return err  // 这里也会调用 cancel()
    }
    if err := step2(ctx); err != nil {
        return err  // 这里也会调用 cancel()
    }
    return nil
}
```

#### 4. 特殊情况：可以不调用 cancel 吗？

```go
// 理论上可以不调用 cancel 的情况（但不推荐）：
// 1. 程序即将退出（os.Exit 或 main 函数结束）
// 2. 明确知道 parent context 会很快被取消

// ❌ 即便如此，仍建议始终调用 cancel()：
// - 代码可能被复用，无法保证 parent 一定会被取消
// - 保持一致的编码习惯，避免遗漏
// - 资源及时释放，减少不必要的开销
```

### 常见错误

```go
// ❌ 忘记调用 cancel() 导致 goroutine 泄漏
ctx, cancel := context.WithCancel(parent)
// defer cancel()  // 忘记调用！

// ❌ 在 select 中随机选择导致饥饿
// 如果某个 case 总是就绪，其他 case 可能永远得不到执行

// ❌ WaitGroup 计数错误
wg.Add(1)
go func() {
    // 某些路径没有调用 Done()
    if someCondition {
        return  // 忘记 Done()！
    }
    wg.Done()
}()

// ✅ 总是使用 defer wg.Done()
wg.Add(1)
go func() {
    defer wg.Done()
    // ...
}()
```

### ⚠️ Select 发送后接收的死锁陷阱

这是一个**极具迷惑性**的易错点：

```go
ch := make(chan int)

select {
case ch <- 1:        // case A: 发送成功
    <-ch             // 然后在 case 内部接收
}
```

**为什么会死锁？**

1. `select` 选择 `case ch <- 1` 意味着**已经有接收方准备好**接收数据（但实际上并没有）
2. case 内的 <-ch 永远不会执行
3. 由于这是单 goroutine，没有其他 goroutine 能向 ch 发送数据 → **死锁**

**正确的做法**：发送和接收应该在**不同的 goroutine** 中进行：

```go
ch := make(chan int)

// 接收方
go func() {
    <-ch
}()

// 发送方
ch <- 1
```

或者使用**有缓冲 channel**：

```go
ch := make(chan int, 1)  // 容量为 1

select {
case ch <- 1:    // 不会阻塞，因为有缓冲
    v := <-ch    // 可以正常接收
    fmt.Println(v)  // 1
}
```

**核心原则**：`select` 的 case 只是**选择哪个操作可以执行**，不代表 case 内部是一个独立的执行上下文。case 被选中后，代码会**顺序执行**。
