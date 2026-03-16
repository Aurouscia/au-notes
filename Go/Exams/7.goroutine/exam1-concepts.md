# Goroutine 与 Channel - 概念题

## 一、填空题

1. 启动一个 goroutine 需要使用 ______ 关键字。

2. Channel 必须使用 ______ 函数进行初始化，否则是 nil。

3. 无缓冲 channel 的容量是 ______，有缓冲 channel 使用 `make(chan T, n)` 创建，其中 n 是 ______。

4. 向已关闭的 channel 发送数据会导致 ______，重复关闭 channel 也会导致 ______。

5. 单向 channel 中，`chan<- int` 表示只能______，`<-chan int` 表示只能______。

6. 从 channel 接收数据时，使用 `value, ok := <-ch` 的形式，当 channel 已关闭且无数据时，`ok` 的值为 ______。

---

## 二、判断题（请写出判断结果和理由）

1. 以下代码可以正确编译和运行：
   ```go
   func main() {
       ch := make(chan int)
       ch <- 42
       fmt.Println(<-ch)
   }
   ```

2. 以下代码可以正确编译和运行：
   ```go
   func main() {
       ch := make(chan int, 1)
       ch <- 42
       close(ch)
       fmt.Println(<-ch)
   }
   ```

3. 以下代码会输出 `true`：
   ```go
   func main() {
       var ch chan int
       fmt.Println(ch == nil)
   }
   ```

4. 以下代码可以正确编译：
   ```go
   func send(ch chan<- int) {
       ch <- 42
       fmt.Println(<-ch)
   }
   ```

5. 以下代码会 panic：
   ```go
   func main() {
       ch := make(chan int, 1)
       close(ch)
       ch <- 1
   }
   ```

---

## 三、简答题

1. 请解释无缓冲 channel 和有缓冲 channel 的区别，以及各自适用的场景。

2. 以下代码有什么问题？如何修复？
   ```go
   func main() {
       ch := make(chan int)
       go func() {
           ch <- 1
           ch <- 2
           ch <- 3
       }()
       fmt.Println(<-ch)
   }
   ```

3. 请解释为什么 Go 提倡 "通过通信来共享内存，而不是通过共享内存来通信"。

4. 以下代码的输出是什么？请解释原因。
   ```go
   func main() {
       ch := make(chan int, 2)
       ch <- 1
       ch <- 2
       close(ch)
       
       for v := range ch {
           fmt.Println(v)
       }
       
       v, ok := <-ch
       fmt.Printf("v=%d, ok=%v\n", v, ok)
   }
   ```

5. 以下代码的输出是什么？请解释 goroutine 的执行顺序。
   ```go
   func main() {
       done := make(chan bool)
       
       go func() {
           fmt.Println("Goroutine start")
           time.Sleep(100 * time.Millisecond)
           fmt.Println("Goroutine end")
           done <- true
       }()
       
       fmt.Println("Main waiting")
       <-done
       fmt.Println("Main done")
   }
   ```
