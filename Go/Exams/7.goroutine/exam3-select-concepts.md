# Select、WaitGroup 与 Context 概念题

## 一、填空题

1. `select` 语句用于在多个 ______ 操作中进行选择。
    - channel 接收

2. 当 `select` 的多个 case 同时就绪时，会 ______ 选择一个执行。
    - 随机

3. `sync.WaitGroup` 的三个核心方法是：Add()、______ 和 Wait()。
    - Done()

4. `WaitGroup` 的 `Add()` 方法应该在 ______ goroutine 中调用，而不是在 goroutine 内部。
    - 外部（这样可以确保 goroutine 启动时计数器肯定自增了）

5. 创建带超时的 context 使用 ______ 函数。
    - WithTimeout()

6. 子 context 会继承父 context 的 ______ 信号和截止时间。
    - 取消

7. 调用 ______ 函数可以取消 context 并传播取消信号给所有子 context。
    - Cancel()

8. 如果不调用 `cancel()`，会导致 ______ 泄漏。
    - gorountine

## 二、判断题（正确的打√，错误的打×）

1. `select` 必须有 `default` 分支。（  ）
    - 错，可以没有

2. `select` 的 case 可以执行任意类型的语句。（  ）
    - 错，只能是 channel 读取

3. 向已关闭的 channel 发送数据会导致 panic。（  ）
    - 对

4. `WaitGroup` 可以重复使用，不需要重新创建。（  ）
    - 不知道

5. `WaitGroup` 应该传递指针给 goroutine，而不是值。（  ）
    - 对，否则会形成两个互不相干的 waitGroup

6. `context.WithValue` 返回的 context 也需要调用 cancel。（  ）
    - 不知道

7. 父 context 取消时，所有子孙 context 都会收到取消信号。（  ）
    - 对

8. 子 context 的超时会影响父 context。（  ）
    - 错，不会向上级影响

9. `WithTimeout` 即使超时自动触发，也需要调用 `cancel()` 释放资源。（  ）
    - 对

10. 使用 `defer cancel()` 可以确保即使函数提前返回也能释放资源。（  ）
    - 对

## 三、简答题

1. 请说明 `select` 语句的两种使用模式（阻塞模式和非阻塞模式），并给出示例。
    - 无 default 的是阻塞模式，select 会一直等待直到有一个 case 接收
        ```go
        select{
            case: v1 := <-ch1
            case: <-ch2
        }
        ```
    - 有 default 的是非阻塞模式，如果没有一个 case 可接收 select，会走 default 
        ```go
        select{
            case: <-ch1
            case: <-ch2
            default:
                fmt.Println("no case matched")
        }
        ```

2. 解释为什么 `WaitGroup` 的 `Add()` 方法要在启动 goroutine 之前调用，而不是在 goroutine 内部调用。
    - 如果放在内部，可能导致外层 goroutine 提前遇到 Wait（此时 goroutine 还未来得及执行，计数器为 0）引发问题

3. 请画出 Context 的树状结构示意图，并说明父子 context 之间的关系。
    ```md
    ---background
            \
            timeout1---timeout2
                    \
                    timeout3
    ```
    - 父级 context 取消会导致子孙 context 均取消，例如 timeout1 取消代表 timeout2 和 3 都取消了
    - 反之不会，子级 context 取消不影响父级，所以子级 timeout 一般短于父级的


4. 解释不调用 `cancel()` 会导致 goroutine 泄漏的原理。
    - 一般来说 goroutine 内部可能会等待着 context.Done() 才结束运行
    - 如果不取消，goroutine 会一直等待下去

5. 在实际项目中，什么时候应该使用 `context.WithTimeout`，什么时候使用 `context.WithCancel`？
    - 需要一定时间后自动取消的情况下使用前者
    - 需要手动控制取消的情况下使用后者

## 四、代码分析题

### 题目 1

```go
func main() {
    ch := make(chan int)
    
    select {
    case ch <- 1:
        fmt.Println("sent")
    case v := <-ch:
        fmt.Println("received:", v)
    }
}
```

问题：这段代码的输出是什么？会死锁吗？为什么？
- 什么都不会输出
- 因为 select 外没有向 ch 写入或读取的地方，这两个 case 都始终不会触发

### 题目 2

```go
func main() {
    var wg sync.WaitGroup
    
    for i := 0; i < 3; i++ {
        go func() {
            wg.Add(1)
            defer wg.Done()
            fmt.Println("worker")
        }()
    }
    
    wg.Wait()
    fmt.Println("done")
}
```

问题：这段代码有什么问题？如何修复？
- wg.Add(1) 不应该在 goroutine 内部调用，应该放在 go 的上一行，否则可能提前通过 wg.Wait()

### 题目 3

```go
func process() {
    ctx, _ := context.WithTimeout(context.Background(), 2*time.Second)
    result, err := queryDB(ctx)
    if err != nil {
        log.Println(err)
        return
    }
    fmt.Println(result)
}
```

问题：这段代码有什么问题？如何修复？
- 没有接收 cancel 函数
- 应该 ctx, cancel := ... ，并在下方 defer cancel()

### 题目 4

```go
func handler(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    result, err := serviceA(ctx)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    
    // 注意：这里复用了上面的 ctx
    result2, err := serviceB(ctx)
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    
    fmt.Fprintf(w, "%s, %s", result, result2)
}
```

问题：这段代码有什么问题？假设 serviceA 耗时 4 秒，serviceB 也需要 4 秒，会发生什么？如何修复？
- AB 顺序执行，A 结束时，B 才开始，B 开始 1 秒后 context 就会超时
- 应该把两个 service 都放在 goroutine 里面执行，并通过 waitGroup 等待它们完成