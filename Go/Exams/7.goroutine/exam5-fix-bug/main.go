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
	ch := make(chan int)
	
	select {
	case ch <- 1:
		fmt.Println("sent")
		v := <-ch
		fmt.Println("received:", v)
	}
}

// ========== 题目 2：WaitGroup 计数错误 ==========
func worker2(id int, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func problem2() {
	fmt.Println("\n=== 题目 2 ===")
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		go worker2(i, wg)
	}
	
	wg.Wait()
	fmt.Println("All workers finished")
}

// ========== 题目 3：Context 泄漏 ==========
func process3(ctx context.Context) error {
	// 创建子 context，但忘记 cancel
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	
	select {
	case <-time.After(3 * time.Second):
		return fmt.Errorf("timeout")
	case <-ctx.Done():
		return ctx.Err()
	}
}

func problem3() {
	fmt.Println("\n=== 题目 3 ===")
	for i := 0; i < 1000; i++ {
		ctx := context.Background()
		process3(ctx)
	}
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
	var wg sync.WaitGroup
	
	wg.Add(1)
	go producer4(ch, &wg)
	
	wg.Add(1)
	go consumer4(ch, &wg)
	
	wg.Wait()
	close(ch)
	
	fmt.Println("Done")
}

// ========== 题目 5：Context 值传递错误 ==========
func handler5(ctx context.Context) {
	// 设置用户 ID
	ctx = context.WithValue(ctx, "userID", "12345")
	
	process5(ctx)
}

func process5(ctx context.Context) {
	if userID := ctx.Value("userID"); userID != nil {
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
	
	// problem1()
	// problem2()
	// problem3()
	// problem4()
	// problem5()
	
	fmt.Println("请取消注释运行各个题目")
}
