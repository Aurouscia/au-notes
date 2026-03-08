package main

import (
	"fmt"

	"example.com/myapp/config"
)

func main() {
	currentLogLevel := config.LevelInfo
	var connectionCount int
	fmt.Printf("应用：%s %s\n", config.AppName, config.Version)
	fmt.Printf("日志级别：%d\n", currentLogLevel)
	fmt.Printf("当前连接：%d / %d\n", connectionCount, config.MaxConnections)

	a, b := 10, 20
	fmt.Printf("交换前：a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("交换前：a=%d, b=%d\n", a, b)
}
