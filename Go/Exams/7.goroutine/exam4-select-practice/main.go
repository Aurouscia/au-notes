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
	m.download(url, resultCh)

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

	resultCh := make(chan DownloadResult, len(m.urls))

	wg := sync.WaitGroup{}
	for _, u := range m.urls {
		wg.Add(1)
		go func() {
			resultCh <- m.downloadWithTimeout(u, timeout)
			// wg.Done() // 错误：不应该在这里done，而是应该在收集结果后 done，否则可能收集不全
		}()
	}

	results := []DownloadResult{}
	go func() {
		for res := range resultCh {
			// fmt.Printf("接收到结果：%s，是否有错误：%v\n", res.URL, res.Err != nil)
			results = append(results, res)
			wg.Done()
		}
	}()

	wg.Wait()
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
