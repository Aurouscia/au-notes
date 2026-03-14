package main

import (
	"fmt"
	"sync"
)

// Counter 接口 - 只读操作
type Counter interface {
	Get() int
	String() string
}

// MutableCounter 接口 - 可修改操作
type MutableCounter interface {
	Counter
	Add(n int)
	Reset()
}

// 请实现 BaseCounter 类型

type BaseCounter struct {
	// 请补充字段
}

// 请实现构造函数 NewBaseCounter

// 请使用指针接收者实现 MutableCounter 接口的所有方法

// SafeCounter 线程安全计数器（可选挑战）
type SafeCounter struct {
	mu    sync.Mutex
	value int
	name  string
}

// 请同样实现 MutableCounter 接口（使用指针接收者）

// CounterFactory 工厂函数类型
type CounterFactory func(name string) MutableCounter

// Registry 类型注册表
type Registry struct {
	factories map[string]CounterFactory
}

// 请实现 Registry 的 Register 和 Create 方法

// ProcessCounter 请实现此函数
// 使用类型断言判断 c 是否是 MutableCounter
func ProcessCounter(c Counter) {
	// 请实现
}

// SumCounters 请实现此函数
func SumCounters(counters ...Counter) int {
	// 请实现
	return 0
}

// FindByName 请实现此函数
func FindByName(counters []Counter, name string) (Counter, bool) {
	// 请实现
	return nil, false
}

func main() {
	fmt.Println("=== 接口进阶练习 ===")
	// 请在此编写测试代码
}
