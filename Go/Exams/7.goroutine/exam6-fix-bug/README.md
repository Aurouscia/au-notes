# 纠错题：找出并修复代码中的问题

## 说明

以下代码包含多个与 `select`、`WaitGroup` 和 `context` 相关的错误，请找出并修复。

## 题目 1：Select 死锁

```go
package main

import "fmt"

func main() {
	ch := make(chan int)
	
	select {
	case ch <- 1:
		fmt.Println("sent")
		v := <-ch
		fmt.Println("received:", v)
	}
}
```

问题：这段代码会死锁，请解释原因并修复。

---

## 题目 2：WaitGroup 计数错误

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	
	for i := 1; i <= 3; i++ {
		go worker(i, wg)
	}
	
	wg.Wait()
	fmt.Println("All workers finished")
}
```

问题：这段代码有什么问题？如何修复？

---

## 题目 3：Context 泄漏

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func process(ctx context.Context) error {
	// 创建子 context，但忘记 cancel
	ctx, _ = context.WithTimeout(ctx, 2*time.Second)
	
	select {
	case <-time.After(3 * time.Second):
		return fmt.Errorf("timeout")
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	for i := 0; i < 1000; i++ {
		ctx := context.Background()
		process(ctx)
	}
	fmt.Println("Done")
}
```

问题：这段代码有什么问题？会导致什么后果？如何修复？

---

## 题目 4：Channel 关闭顺序错误

```go
package main

import (
	"fmt"
	"sync"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		ch <- i
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	
	wg.Add(1)
	go producer(ch, &wg)
	
	wg.Add(1)
	go consumer(ch, &wg)
	
	wg.Wait()
	close(ch)
	
	fmt.Println("Done")
}
```

问题：这段代码有什么问题？如何修复？

---

## 题目 5：Context 值传递错误

```go
package main

import (
	"context"
	"fmt"
)

func handler(ctx context.Context) {
	// 设置用户 ID
	ctx = context.WithValue(ctx, "userID", "12345")
	
	process(ctx)
}

func process(ctx context.Context) {
	if userID := ctx.Value("userID"); userID != nil {
		fmt.Println("UserID:", userID)
	} else {
		fmt.Println("UserID not found")
	}
}

func main() {
	ctx := context.Background()
	handler(ctx)
}
```

问题：这段代码的输出是什么？有什么问题？如何修复？
