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
	value int
	Name  string
}

// 请实现构造函数 NewBaseCounter

func NewBaseCounter(name string) *BaseCounter {
	var res = BaseCounter{value: 0, Name: name}
	return &res
}

// 请使用指针接收者实现 MutableCounter 接口的所有方法

func (c BaseCounter) Get() int {
	return c.value
}

func (c BaseCounter) String() string {
	return fmt.Sprintf("Counter[%s]: %d", c.Name, c.value)
}

func (c *BaseCounter) Add(n int) {
	c.value += n
}

func (c *BaseCounter) Reset() {
	c.value = 0
}

// SafeCounter 线程安全计数器（可选挑战）
type SafeCounter struct {
	mu sync.Mutex
	BaseCounter
}

// 请同样实现 MutableCounter 接口（使用指针接收者）
func (c *SafeCounter) Add(n int) {
	c.mu.Lock()
	c.value += n
	c.mu.Unlock()
}

func (c *SafeCounter) Reset() {
	c.mu.Lock()
	c.value = 0
	c.mu.Unlock()
}

func (c *SafeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func (c *SafeCounter) String() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return fmt.Sprintf("Counter[%s]: %d", c.Name, c.value)
}

// CounterFactory 工厂函数类型
type CounterFactory func(name string) MutableCounter

// Registry 类型注册表
type Registry struct {
	factories map[string]CounterFactory
}

// 请实现 Registry 的 Register 和 Create 方法
func (r *Registry) Register(typeName string, factory CounterFactory) {
	r.factories[typeName] = factory
}
func (r *Registry) Create(typeName string, name string) (MutableCounter, error) {
	c, ok := r.factories[typeName]
	if !ok {
		return nil, fmt.Errorf("未注册该类型的 counter")
	}
	return c(name), nil
}

// ProcessCounter 请实现此函数
// 使用类型断言判断 c 是否是 MutableCounter
func ProcessCounter(c Counter) {
	mc, isMutable := c.(MutableCounter)
	if isMutable {
		fmt.Println("是 mutableCounter，操作前的值：", c.Get())
		mc.Add(10)
		fmt.Println("Add(10) 操作后的值：", c.Get())
		mc.Reset()
		fmt.Println("Reset() 后的值：", c.Get())
	}
}

// SumCounters 请实现此函数
func SumCounters(counters ...Counter) int {
	sum := 0
	for _, c := range counters {
		sum += c.Get()
	}
	return 0
}

// FindByName 请实现此函数
func FindByName(counters []Counter, name string) (Counter, bool) {
	fmt.Println()
	fmt.Println("find by name: ")
	for _, c := range counters {
		bc, isCounter := c.(BaseCounter)
		fmt.Printf("%T, isCounter=%v\n", bc, isCounter)
		if isCounter {
			if bc.Name == name {
				return bc, true
			}
		}
	}
	fmt.Println()
	return nil, false
}

func main() {
	fmt.Println("=== 接口进阶练习 ===")
	bc := BaseCounter{Name: "counter1"}
	log := func() {
		fmt.Println("bc 的值：", bc.Get())
	}
	log()
	bc.Add(10)
	log()
	bc.Reset()
	log()
	fmt.Println(bc)

	var counter Counter = &bc
	// var mutableCounter MutableCounter = bc // BaseCounter does not implement MutableCounter (method Add has pointer receiver)
	var mutableCounter MutableCounter = &BaseCounter{Name: "counter2"}
	ProcessCounter(mutableCounter)
	ProcessCounter(counter)

	fmt.Println("sum:", SumCounters(counter, mutableCounter))
	someCounter, found := FindByName([]Counter{counter, mutableCounter}, "counter2")
	if found {
		mc := someCounter.(MutableCounter)
		mc.Add(50)
	}

	var baseFac CounterFactory = func(name string) MutableCounter {
		return &BaseCounter{Name: name}
	}
	var safeFac CounterFactory = func(name string) MutableCounter {
		return &SafeCounter{BaseCounter: BaseCounter{Name: name}}
	}
	r := Registry{factories: map[string]CounterFactory{}}
	r.Register("base", baseFac)
	r.Register("safe", safeFac)
	safe1, _ := r.Create("safe", "safe1")
	safe2, _ := r.Create("safe", "safe2")
	base1, _ := r.Create("base", "base1")

	counters := []Counter{safe1, safe2, base1}
	for _, c := range counters {
		switch c.(type) {
		case *BaseCounter:
			fmt.Println("是baseCounter")
		case *SafeCounter:
			fmt.Println("是safeCounter")
		}
	}
}
