package main

import (
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
	// TODO: 实现
}

// download 模拟下载（内部方法）
func (m *MultiDownloader) download(url string, resultCh chan<- DownloadResult) {
	// TODO: 实现
	// 1. 随机睡眠 1-5 秒模拟下载
	// 2. 发送结果到 resultCh
}

// downloadWithTimeout 带超时的下载
func downloadWithTimeout(url string, timeout time.Duration) DownloadResult {
	// TODO: 实现
	// 使用 select 实现超时控制
	return DownloadResult{}
}

// DownloadAll 并发下载所有任务
func (m *MultiDownloader) DownloadAll(timeout time.Duration) []DownloadResult {
	// TODO: 实现
	// 1. 为每个 URL 启动 goroutine 下载
	// 2. 使用 WaitGroup 等待所有任务完成
	// 3. 收集所有结果并返回
	return nil
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
