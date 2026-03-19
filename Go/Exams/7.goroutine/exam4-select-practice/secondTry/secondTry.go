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
	secCount := rand.Int31n(5) + 1
	fmt.Printf("开始任务 %s，随机秒数：%d\n", url, secCount)
	dur := time.Duration(secCount) * time.Second // 先使用类型转换，创建指定纳秒的 Duration，然后再通过乘“秒常量”转换为秒
	time.Sleep(dur)
	resultCh <- DownloadResult{
		URL:      url,
		Data:     "Data",
		Duration: dur,
		Err:      nil,
	}
}

// downloadWithTimeout 带超时的下载
func (m *MultiDownloader) downloadWithTimeout(url string, timeout time.Duration) DownloadResult {
	// 使用 select 实现超时控制

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel() // 虽然时间到了会自动 cancel，但最好是在任务提前完成时 cancel，避免定时器堆积或 worker 多余工作

	resultCh := make(chan DownloadResult, 1)
	// defer close(resultCh)
	// ❌ 这里不应该有 defer close，否则超时的情况下会发生“已关闭channel写入”的问题
	// 直接不 close 就离开，只要 resultCh 有缓冲区，下载 goroutine 就能完成，不会阻塞
	// 最终 channel 会由 GC 自动回收（但如果 channel 没有缓冲区的话，goroutine 不会！）
	// 改进方法：使用 context 取消

	go m.download(url, resultCh) // 单独启动一个 routine，避免阻塞，才能让 select 生效

	select {
	case result := <-resultCh:
		return result
	case <-ctx.Done():
		return DownloadResult{
			URL:      url,
			Data:     "",
			Duration: timeout,
			Err:      fmt.Errorf("timeout"),
		}
	}
}

// DownloadAll 并发下载所有任务
func (m *MultiDownloader) DownloadAll(timeout time.Duration) []DownloadResult {
	// 1. 为每个 URL 启动 goroutine 下载
	// 2. 使用 WaitGroup 等待所有任务完成
	// 3. 收集所有结果并返回

	resultCh := make(chan DownloadResult, len(m.urls))
	// defer close(resultCh) // ❌ 不能 defer close，否则 for-range 会卡死（它会等到 channel 关闭才脱离）
	wg := sync.WaitGroup{}
	for _, u := range m.urls {
		wg.Add(1)
		go func() {
			resultCh <- m.downloadWithTimeout(u, timeout)
		}()
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	results := []DownloadResult{}
	for r := range resultCh {
		results = append(results, r)
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

	results := downloader.DownloadAll(3 * time.Second)

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
