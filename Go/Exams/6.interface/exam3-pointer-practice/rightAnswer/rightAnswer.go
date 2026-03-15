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

// BaseCounter 基础计数器
type BaseCounter struct {
	value int
	name  string // 小写，通过方法访问
}

// NewBaseCounter 构造函数
func NewBaseCounter(name string) *BaseCounter {
	// 最佳实践：直接返回 &BaseCounter{...}，不需要中间变量
	return &BaseCounter{value: 0, name: name}
}

// Get 值接收者 - Counter 接口要求
// 值接收者和指针接收者都满足 Counter 接口
func (c BaseCounter) Get() int {
	return c.value
}

// String 值接收者 - fmt.Stringer 接口要求
// 最佳实践：方法名是 String，不是 Stringer
func (c BaseCounter) String() string {
	return fmt.Sprintf("Counter[%s]: %d", c.name, c.value)
}

// Add 指针接收者 - 需要修改内部状态
// 最佳实践：会修改接收者的方法必须用指针接收者
func (c *BaseCounter) Add(n int) {
	c.value += n
}

// Reset 指针接收者 - 需要修改内部状态
func (c *BaseCounter) Reset() {
	c.value = 0
}

// Name 获取名称
func (c BaseCounter) Name() string {
	return c.name
}

// SafeCounter 线程安全计数器
type SafeCounter struct {
	mu    sync.Mutex
	value int
	name  string
}

// NewSafeCounter 构造函数
func NewSafeCounter(name string) *SafeCounter {
	return &SafeCounter{name: name}
}

// Get 使用 defer 确保解锁
// 最佳实践：简单函数用 defer，复杂函数手动 unlock 提高性能
func (c *SafeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// String 线程安全的 String 方法
func (c *SafeCounter) String() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return fmt.Sprintf("SafeCounter[%s]: %d", c.name, c.value)
}

// Add 线程安全的 Add
func (c *SafeCounter) Add(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += n
}

// Reset 线程安全的 Reset
func (c *SafeCounter) Reset() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value = 0
}

// Name 获取名称
func (c *SafeCounter) Name() string {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.name
}

// CounterFactory 工厂函数类型
type CounterFactory func(name string) MutableCounter

// Registry 类型注册表
type Registry struct {
	// 最佳实践：使用 make 初始化 map，或在构造函数中初始化
	factories map[string]CounterFactory
}

// NewRegistry 构造函数
// 最佳实践：提供构造函数确保内部 map 已初始化
func NewRegistry() *Registry {
	return &Registry{
		factories: make(map[string]CounterFactory),
	}
}

// Register 注册工厂函数
// 最佳实践：检查 registry 是否为 nil，避免 panic
func (r *Registry) Register(typeName string, factory CounterFactory) {
	if r.factories == nil {
		r.factories = make(map[string]CounterFactory)
	}
	r.factories[typeName] = factory
}

// Create 根据类型名创建计数器
// 最佳实践：返回具体的错误信息
func (r *Registry) Create(typeName string, name string) (MutableCounter, error) {
	factory, ok := r.factories[typeName]
	if !ok {
		return nil, fmt.Errorf("unknown counter type: %s", typeName)
	}
	return factory(name), nil
}

// ProcessCounter 处理计数器
// 最佳实践：使用类型断言检查能力，而不是具体类型
func ProcessCounter(c Counter) {
	// 类型断言：检查 c 是否也是 MutableCounter
	if mc, ok := c.(MutableCounter); ok {
		fmt.Printf("MutableCounter detected: %s\n", c)
		fmt.Println("  Before Add(10):", c.Get())
		mc.Add(10)
		fmt.Println("  After Add(10):", c.Get())
		mc.Reset()
		fmt.Println("  After Reset():", c.Get())
	} else {
		fmt.Printf("Read-only Counter: %s\n", c)
	}
}

// SumCounters 计算总和
// 最佳实践：使用 int 而不是 float64，除非需要小数
func SumCounters(counters ...Counter) int {
	sum := 0
	for _, c := range counters {
		sum += c.Get()
	}
	return sum
}

// FindByName 根据名称查找计数器
// 最佳实践：使用接口方法获取名称，而不是类型断言
// 这样任何实现了 Name() 方法的 Counter 都可以被查找

type NamedCounter interface {
	Counter
	Name() string
}

func FindByName(counters []Counter, name string) (Counter, bool) {
	for _, c := range counters {
		// 类型断言为 NamedCounter 接口，而不是具体类型
		if nc, ok := c.(NamedCounter); ok && nc.Name() == name {
			return c, true
		}
	}
	return nil, false
}

func main() {
	fmt.Println("=== 接口进阶练习（正确答案）===")

	// 1. 创建 BaseCounter，测试 MutableCounter 接口
	fmt.Println("\n1. BaseCounter 基本操作:")
	bc := NewBaseCounter("counter1")
	fmt.Println("Initial:", bc)
	bc.Add(10)
	fmt.Println("After Add(10):", bc)
	bc.Reset()
	fmt.Println("After Reset():", bc)

	// 2. 将 *BaseCounter 赋值给 Counter 接口变量（自动解引用）
	fmt.Println("\n2. 自动解引用测试:")
	var c Counter = bc // *BaseCounter 自动解引用调用 Get() 和 String()
	fmt.Println("Counter interface:", c)

	// 3. 值类型不能赋值给 MutableCounter
	// var mc MutableCounter = *bc // 编译错误：BaseCounter 没有 Add 方法
	// 正确做法：
	var mc MutableCounter = bc // *BaseCounter 实现了 MutableCounter
	mc.Add(100)
	fmt.Println("After mc.Add(100):", mc)

	// 4. ProcessCounter 测试
	fmt.Println("\n3. ProcessCounter 测试:")
	ProcessCounter(mc)          // 是 MutableCounter
	ProcessCounter(c)           // 也是 MutableCounter（因为 c 和 mc 指向同一个对象）

	// 5. SumCounters 测试
	fmt.Println("\n4. SumCounters 测试:")
	c1 := NewBaseCounter("c1")
	c1.Add(10)
	c2 := NewBaseCounter("c2")
	c2.Add(20)
	sum := SumCounters(c1, c2)
	fmt.Printf("Sum of %s and %s = %d\n", c1, c2, sum)

	// 6. FindByName 测试
	fmt.Println("\n5. FindByName 测试:")
	counters := []Counter{c1, c2}
	if found, ok := FindByName(counters, "c2"); ok {
		fmt.Println("Found:", found)
		// 类型断言转回 MutableCounter 进行修改
		if mc, ok := found.(MutableCounter); ok {
			mc.Add(50)
			fmt.Println("After Add(50):", found)
		}
	}

	// 7. Registry 测试
	fmt.Println("\n6. Registry 测试:")
	registry := NewRegistry()
	registry.Register("base", NewBaseCounter)
	registry.Register("safe", NewSafeCounter)

	safe1, err := registry.Create("safe", "safe-counter-1")
	if err != nil {
		panic(err)
	}
	safe1.Add(42)
	fmt.Println("Created and modified:", safe1)

	base1, err := registry.Create("base", "base-counter-1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Created:", base1)

	// 8. type switch 测试
	fmt.Println("\n7. Type switch 测试:")
	testCounters := []Counter{safe1, base1}
	for _, c := range testCounters {
		switch v := c.(type) {
		case *SafeCounter:
			fmt.Printf("SafeCounter: %s (thread-safe)\n", v)
		case *BaseCounter:
			fmt.Printf("BaseCounter: %s (not thread-safe)\n", v)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}

	// 9. 测试未注册的类型
	fmt.Println("\n8. 错误处理测试:")
	if _, err := registry.Create("unknown", "test"); err != nil {
		fmt.Println("Expected error:", err)
	}
}
