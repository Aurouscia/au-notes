# Goroutine 协程基础

## 1. 什么是 Goroutine

Goroutine 是 Go 语言中的**轻量级线程**，由 Go 运行时（runtime）管理，而非操作系统直接管理。

### 特点

| 特性 | 说明 |
|------|------|
| 轻量级 | 初始栈仅 2KB，可动态增长和收缩 |
| 低成本 | 创建和切换开销远小于 OS 线程 |
| 多对多 | 多个 Goroutine 复用少量 OS 线程 |
| 内置调度 | Go 运行时自动调度，无需手动管理 |

### 启动 Goroutine

使用 `go` 关键字启动：

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    for i := 0; i < 3; i++ {
        fmt.Println("Hello from goroutine")
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    // 启动一个 goroutine
    go sayHello()
    
    // 主 goroutine 继续执行
    for i := 0; i < 3; i++ {
        fmt.Println("Hello from main")
        time.Sleep(100 * time.Millisecond)
    }
    
    // 等待一下，让 goroutine 有机会执行
    time.Sleep(500 * time.Millisecond)
}
```

### 匿名函数启动

```go
func main() {
    // 使用匿名函数
    go func() {
        fmt.Println("Anonymous goroutine")
    }()
    
    // 带参数的匿名函数
    msg := "Hello"
    go func(m string) {
        fmt.Println(m)
    }(msg)
    
    time.Sleep(time.Second)
}
```

---

## 2. Channel 通道

Channel 是 Goroutine 之间**通信和同步**的机制，遵循 CSP（Communicating Sequential Processes）模型。

> **核心理念**：不要通过共享内存来通信，而是通过通信来共享内存。

### 声明与创建

```go
// 声明
var ch chan int  // 声明一个传递 int 的 channel

// 创建（必须使用 make）
ch1 := make(chan int)      // 无缓冲 channel
ch2 := make(chan int, 5)   // 有缓冲 channel，容量为 5
```

### 基本操作

```go
ch := make(chan int)

// 发送数据
ch <- 42

// 接收数据
value := <-ch

// 接收并忽略值
<-ch

// 检查 channel 是否关闭
value, ok := <-ch  // ok 为 false 表示 channel 已关闭且无数据
```

---

## 3. 无缓冲 Channel（同步通道）

```go
ch := make(chan int)  // 容量为 0
```

### 特性

- **同步通信**：发送和接收必须同时就绪
- **阻塞行为**：
  - 发送方阻塞，直到有接收方接收
  - 接收方阻塞，直到有发送方发送

### 示例

```go
package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Println("Working...")
    time.Sleep(time.Second)
    fmt.Println("Done!")
    done <- true  // 通知主 goroutine
}

func main() {
    done := make(chan bool)
    
    go worker(done)
    
    <-done  // 等待 worker 完成（阻塞）
    fmt.Println("Main finished")
}
```

### ⚠️ 死锁陷阱

```go
func main() {
    ch := make(chan int)
    ch <- 42  // ❌ 死锁！没有接收方，永远阻塞
    fmt.Println(<-ch)
}
```

---

## 4. 有缓冲 Channel（异步通道）

```go
ch := make(chan int, 3)  // 容量为 3
```

### 特性

- **异步通信**：发送方在缓冲区未满时不阻塞
- **缓冲行为**：
  - 缓冲区未满：发送不阻塞
  - 缓冲区已满：发送阻塞，等待接收
  - 缓冲区非空：接收不阻塞
  - 缓冲区为空：接收阻塞，等待发送

### 示例

```go
package main

import "fmt"

func main() {
    ch := make(chan int, 2)
    
    // 发送两个值，不会阻塞（缓冲区未满）
    ch <- 1
    ch <- 2
    
    // 接收值
    fmt.Println(<-ch)  // 1
    fmt.Println(<-ch)  // 2
}
```

### 查看缓冲状态

```go
ch := make(chan int, 5)

fmt.Println(len(ch))  // 当前缓冲区元素个数：0
fmt.Println(cap(ch))  // 缓冲区容量：5

ch <- 1
ch <- 2
fmt.Println(len(ch))  // 2
```

---

## 5. 关闭 Channel

### 规则

- 只能由**发送方**关闭
- 关闭后**不能再发送**数据（会 panic）
- 关闭后**可以继续接收**已缓冲的数据
- 接收时通过 `ok` 值判断是否已关闭

```go
package main

import "fmt"

func producer(ch chan int) {
    for i := 0; i < 3; i++ {
        ch <- i
    }
    close(ch)  // 生产者关闭 channel
}

func main() {
    ch := make(chan int)
    
    go producer(ch)
    
    // 方式1：使用 range 遍历（推荐）channel 关闭时，循环会自动退出，无需检查 ok
    for v := range ch {
        fmt.Println(v)  // 0, 1, 2
    }
    
    // 方式2：使用 ok 检查
    for {
        v, ok := <-ch
        if !ok {
            break  // channel 已关闭
        }
        fmt.Println(v)
    }
}
```

---

## 6. 单向 Channel

用于限制 channel 的操作权限，提高代码安全性：

```go
// 只发送（Send-only）
func producer(ch chan<- int) {
    ch <- 42
    // <-ch  // ❌ 编译错误：不能从 send-only channel 接收
}

// 只接收（Receive-only）
func consumer(ch <-chan int) {
    v := <-ch
    // ch <- 1  // ❌ 编译错误：不能向 receive-only channel 发送
}

func main() {
    ch := make(chan int)
    go producer(ch)
    consumer(ch)
}
```

---

## 7. 核心要点总结

| 概念 | 要点 |
|------|------|
| Goroutine | `go` 关键字启动，轻量级并发执行单元 |
| Channel | 使用 `make` 创建，引用类型 |
| 无缓冲 Channel | 同步通信，发送接收必须配对 |
| 有缓冲 Channel | 异步通信，缓冲区满时发送阻塞 |
| 关闭 Channel | 发送方关闭，关闭后不能发送，可继续接收 |
| 单向 Channel | `chan<-` 只发送，`<-chan` 只接收 |

### 常见错误

```go
// ❌ 操作 nil channel → 永久阻塞
var ch chan int
<-ch  // 永远阻塞

// ❌ 向已关闭的 channel 发送 → panic
close(ch)
ch <- 1  // panic: send on closed channel

// ❌ 重复关闭 channel → panic
close(ch)
close(ch)  // panic: close of closed channel

// ❌ 无缓冲 channel 单 goroutine 收发 → 死锁
ch := make(chan int)
ch <- 1  // 阻塞，没有接收方
```

## 8.多个 goroutine 写入同一结果时的最佳实践

为什么使用一个 channel 接受结果？与“创建一个切片让每个 goroutine 往里写”有什么区别？

- 同步方式
    - 通信（CSP）
    - 共享内存 + 锁
- 代码复杂度
    - 简单，无锁
    - 需要手动加锁（否则可能会数据错乱）
- 安全性
    - 天然线程安全
    - 容易忘记加锁导致竞态
- 阻塞语义
    - 清晰的阻塞/等待
    - 锁的粒度难控制
- 可扩展性
    - 容易扩展为流水线
    - 锁竞争可能成为瓶颈

最佳选择：

- 结果数量不确定：Channel
- 需要流水线处理：Channel
- 需要超时/取消：Channel
- 结果数量固定，且每个 worker 写固定位置：共享切片（索引写入）
- 一般并发任务：Channel

Channel 是 Go 并发编程的精髓，优先使用 channel，只有在性能关键且能确保无竞态时才考虑共享内存。