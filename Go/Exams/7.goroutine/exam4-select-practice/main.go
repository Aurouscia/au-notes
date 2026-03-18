package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// DownloadResult 表示下载结果
type DownloadResult struct {
	URL      string
	Data     string
	Duration time.Duration
	Err      error
}

// MultiDownloader 多任务下载器
type MultiDownloader struct {
	urls []string
}

// Add 添加下载任务
func (m *MultiDownloader) Add(url string) {
	m.urls = append(m.urls, url)
}

// download 模拟下载（内部方法）
func (m *MultiDownloader) download(url string, resultCh chan<- DownloadResult) {
	// 1. 随机睡眠 1-5 秒模拟下载
	// 2. 发送结果到 resultCh
	secondCount := rand.Int31n(5) + 1
	fmt.Printf("对于 %s 的随机时长：%ds\n", url, secondCount)
	dur := time.Duration(secondCount) * time.Second
	time.Sleep(dur)
	result := DownloadResult{
		URL:      url,
		Data:     "Data",
		Duration: dur,
		Err:      nil,
	}
	resultCh <- result
}

// downloadWithTimeout 带超时的下载
func (m *MultiDownloader) downloadWithTimeout(url string, timeout time.Duration) DownloadResult {
	// 使用 select 实现超时控制
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resultCh := make(chan DownloadResult, 1)
	// m.download(url, resultCh) // ❌有误：这里不应该是同步调用！否则会阻塞在这里等待，下面的 select 毫无意义
	go m.download(url, resultCh) // 修复：在 goroutine 中执行下载，实现真正的并发超时控制

	select {
	case result := <-resultCh:
		return result
	case <-ctx.Done():
		var errRes = DownloadResult{
			URL:      url,
			Err:      fmt.Errorf("timeout"),
			Duration: timeout,
		}
		return errRes
	}
}

// DownloadAll 并发下载所有任务
func (m *MultiDownloader) DownloadAll(timeout time.Duration) []DownloadResult {
	// 1. 为每个 URL 启动 goroutine 下载
	// 2. 使用 WaitGroup 等待所有任务完成
	// 3. 收集所有结果并返回

	resultCh := make(chan DownloadResult, len(m.urls)) // ❌ 虽然无缓冲区不会死锁，但可能影响并发性能，最好设置缓冲区

	wg := sync.WaitGroup{}
	for _, u := range m.urls {
		wg.Add(1)
		go func() {
			resultCh <- m.downloadWithTimeout(u, timeout)
			// wg.Done() // ❌ 不应该在这里done，而是应该在收集结果后 done，否则可能收集不全（竞态条件）
		}()
	}

	// results := []DownloadResult{}
	// go func() {
	// 	  for res := range resultCh {
	// 		  results = append(results, res)
	// 		  wg.Done()
	// 	  }
	//    close(resultCh)
	// }()
	// wg.Wait() // ❌ for range resultCh 会一直阻塞，直到 resultCh 被关闭。但 close(resultCh) 在循环后面，循环不结束就执行不到

	// 以下是正确的最佳实践：
	// 1. 在专门的 goroutine 中 Wait，Wait 通过后说明任务完成，这时再关闭 channel
	// 2. 在主 goroutine 中收集结果，杜绝线程不安全问题（确保切片和逻辑在同一个 goroutine 中）
	go func() {
		wg.Wait()
		close(resultCh)
	}()
	results := []DownloadResult{}
	for res := range resultCh {
		results = append(results, res)
		wg.Done()
	}

	return results
}

func main() {
	downloader := &MultiDownloader{}

	// 添加下载任务
	downloader.Add("https://api1.example.com/data")
	downloader.Add("https://api2.example.com/data")
	downloader.Add("https://api3.example.com/data")

	fmt.Println("开始下载", len(downloader.urls), "个任务...")

	results := downloader.DownloadAll(3001 * time.Millisecond)

	// 打印结果
	success := 0
	failed := 0
	for _, r := range results {
		if r.Err != nil {
			fmt.Printf("%s - 失败: %v\n", r.URL, r.Err)
			failed++
		} else {
			fmt.Printf("%s - 成功 (%v)\n", r.URL, r.Duration)
			success++
		}
	}

	fmt.Printf("\n下载完成：%d 成功, %d 失败\n", success, failed)
}
