# Select、WaitGroup 与 Context 概念题

## 一、填空题

1. `select` 语句用于在多个 ______ 操作中进行选择。

2. 当 `select` 的多个 case 同时就绪时，会 ______ 选择一个执行。

3. `sync.WaitGroup` 的三个核心方法是：Add()、______ 和 Wait()。

4. `WaitGroup` 的 `Add()` 方法应该在 ______ goroutine 中调用，而不是在 goroutine 内部。

5. 创建带超时的 context 使用 ______ 函数。

6. 子 context 会继承父 context 的 ______ 信号和截止时间。

7. 调用 ______ 函数可以取消 context 并传播取消信号给所有子 context。

8. 如果不调用 `cancel()`，会导致 ______ 泄漏。

## 二、判断题（正确的打√，错误的打×）

1. `select` 必须有 `default` 分支。（  ）

2. `select` 的 case 可以执行任意类型的语句。（  ）

3. 向已关闭的 channel 发送数据会导致 panic。（  ）

4. `WaitGroup` 可以重复使用，不需要重新创建。（  ）

5. `WaitGroup` 应该传递指针给 goroutine，而不是值。（  ）

6. `context.WithValue` 返回的 context 也需要调用 cancel。（  ）

7. 父 context 取消时，所有子孙 context 都会收到取消信号。（  ）

8. 子 context 的超时会影响父 context。（  ）

9. `WithTimeout` 即使超时自动触发，也需要调用 `cancel()` 释放资源。（  ）

10. 使用 `defer cancel()` 可以确保即使函数提前返回也能释放资源。（  ）

## 三、简答题

1. 请说明 `select` 语句的两种使用模式（阻塞模式和非阻塞模式），并给出示例。

2. 解释为什么 `WaitGroup` 的 `Add()` 方法要在启动 goroutine 之前调用，而不是在 goroutine 内部调用。

3. 请画出 Context 的树状结构示意图，并说明父子 context 之间的关系。

4. 解释不调用 `cancel()` 会导致 goroutine 泄漏的原理。

5. 在实际项目中，什么时候应该使用 `context.WithTimeout`，什么时候使用 `context.WithCancel`？

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
