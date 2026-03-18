# 实践题：多路复用下载器

## 需求

实现一个并发下载器，从多个 URL 下载数据，使用 `select` 实现以下功能：

1. **多数据源合并**：同时从多个 URL 下载，哪个先返回就处理哪个
2. **超时控制**：单个下载超时时间为 3 秒
3. **优雅退出**：当所有下载完成或超时时，汇总结果

## 要求

1. 实现 `download(url string, resultCh chan<- DownloadResult)` 函数：
   - 模拟下载耗时（随机 1-5 秒）
   - 将结果发送到 resultCh
   - 如果超时，发送超时错误

2. 实现 `downloadWithTimeout(url string, timeout time.Duration) DownloadResult` 函数：
   - 使用 `select` 实现超时控制
   - 返回下载结果或超时错误

3. 实现 `MultiDownloader` 结构体和方法：
   - `Add(url string)`：添加下载任务
   - `DownloadAll(timeout time.Duration) []DownloadResult`：并发下载所有任务

4. 使用 `sync.WaitGroup` 等待所有下载 goroutine 完成

## 数据结构

```go
type DownloadResult struct {
    URL      string
    Data     string
    Duration time.Duration
    Err      error
}

type MultiDownloader struct {
    urls []string
}
```

## 示例输出

```
开始下载 3 个任务...
https://api1.example.com/data - 成功 (1.2s)
https://api2.example.com/data - 超时
https://api3.example.com/data - 成功 (2.5s)

下载完成：2 成功, 1 失败
```

## 提示

1. 使用 `time.After()` 或 `context.WithTimeout()` 实现超时
2. 使用 `select` 监听多个 channel
3. 使用 `sync.WaitGroup` 等待所有任务完成
4. 注意处理 goroutine 泄漏问题
