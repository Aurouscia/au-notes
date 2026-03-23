package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// ========== 题目 1：Select 死锁 ==========
func problem1() {
	fmt.Println("=== 题目 1 ===")
	// ch := make(chan int)
	ch := make(chan int, 1) // 解决方法：添加缓冲区

	select {
	case ch <- 1:
		fmt.Println("sent")
		v := <-ch
		fmt.Println("received:", v)
	}
}

// ========== 题目 2：WaitGroup 计数错误 ==========
func worker2(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func problem2() {
	fmt.Println("\n=== 题目 2 ===")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)          // 解决方法：在 go 之前 wg.Add(1)
		go worker2(i, &wg) // 解决方法：传入 wg 的指针而不是值
	}

	wg.Wait()
	fmt.Println("All workers finished")
}

// ========== 题目 3：Context 泄漏 ==========
func process3(ctx context.Context, wg *sync.WaitGroup) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel() // 解决方法：defer cancel
	defer wg.Done()

	select {
	case <-time.After(3 * time.Second):
		return fmt.Errorf("timeout")
	case <-ctx.Done():
		return ctx.Err()
	}
}

func problem3() {
	fmt.Println("\n=== 题目 3 ===")
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1) // ❌忘记了 wg.Add(1)，导致 wg.Wait 直接离开
		ctx := context.Background()
		go process3(ctx, &wg)
	}
	wg.Wait()
	fmt.Println("Done")
}

// ========== 题目 4：Channel 关闭顺序错误 ==========
func producer4(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ch <- i
	}
}

func consumer4(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

func problem4() {
	fmt.Println("\n=== 题目 4 ===")
	ch := make(chan int)
	var wgSend sync.WaitGroup
	var wgRecieve sync.WaitGroup

	wgSend.Add(1)
	go producer4(ch, &wgSend)

	wgRecieve.Add(1)
	go consumer4(ch, &wgRecieve)

	wgSend.Wait()
	close(ch)
	wgRecieve.Wait() // 修复方法：分两个 waitGroup，并把接收的 wait 放到 close 后，否则 consumer 会无限等待下去，永远不释放 waitGroup

	fmt.Println("Done")
}

// ========== 题目 5：Context 值传递错误 ==========

// 定义 key 类型
type contextKey string

const userIDKey contextKey = "userID"

func handler5(ctx context.Context) {
	// 设置用户 ID
	ctx = context.WithValue(ctx, userIDKey, "12345") // 这里需要使用专用的 contextKey 类型（虽然底层是 string，但是它的值不能被字符串字面量代替）

	process5(ctx)
}

func process5(ctx context.Context) {
	// 注意：这里使用了 userIDKey 来获取值
	if userID := ctx.Value(userIDKey); userID != nil {
		fmt.Println("UserID:", userID)
	} else {
		fmt.Println("UserID not found")
	}
}

func problem5() {
	fmt.Println("\n=== 题目 5 ===")
	ctx := context.Background()
	handler5(ctx)
}

func main() {
	// 依次运行每个题目，观察问题
	// 注意：部分题目可能导致死锁或 panic

	problem1()
	problem2()
	problem3()
	problem4()
	problem5()

	// fmt.Println("请取消注释运行各个题目")
}
