# 接口进阶 - 指针接收者与薄弱环节强化练习

## 需求

实现一个**可修改的计数器系统**，重点考察**指针接收者**的实际应用和之前薄弱环节的知识点。

### 1. 定义接口

```go
// Counter 接口 - 只读操作
type Counter interface {
    Get() int           // 获取当前值
    String() string     // 返回格式化字符串
}

// MutableCounter 接口 - 可修改操作
type MutableCounter interface {
    Counter             // 嵌入 Counter 接口（接口组合）
    Add(n int)          // 增加 n
    Reset()             // 重置为 0
}
```

### 2. 实现基础计数器

实现 `BaseCounter` 类型，使用**指针接收者**实现 `MutableCounter` 接口：

```go
type BaseCounter struct {
    value int
    name  string
}

// 构造函数
func NewBaseCounter(name string) *BaseCounter
```

**要求：**
- `Add(n int)` - 将 n 加到 value 上
- `Reset()` - 将 value 重置为 0
- `Get() int` - 返回当前 value
- `String() string` - 返回格式 `"Counter[name]: value"`

### 3. 实现线程安全的计数器（可选挑战）

实现 `SafeCounter`，使用 `sync.Mutex` 保护内部状态：

```go
type SafeCounter struct {
    mu    sync.Mutex
    value int
    name  string
}
```

同样实现 `MutableCounter` 接口。

### 4. 实现工具函数

```go
// ProcessCounter 接收 Counter 接口，打印其信息
func ProcessCounter(c Counter) {
    // 使用类型断言判断 c 是否是 MutableCounter
    // 如果是，调用 Add(10) 和 Reset()
    // 打印操作前后的值
}

// SumCounters 计算多个 Counter 的总和
func SumCounters(counters ...Counter) int

// FindByName 在计数器列表中查找指定名称的计数器
// 返回 Counter 和 bool（是否找到）
func FindByName(counters []Counter, name string) (Counter, bool)
```

### 5. 实现类型注册表（考察 type switch 和 map）

```go
// CounterFactory 是一个工厂函数类型
type CounterFactory func(name string) MutableCounter

// Registry 用于注册和创建不同类型的计数器
type Registry struct {
    factories map[string]CounterFactory
}

// Register 注册一个工厂函数
func (r *Registry) Register(typeName string, factory CounterFactory)

// Create 根据类型名创建计数器
func (r *Registry) Create(typeName string, name string) (MutableCounter, error)
```

### 6. 主函数测试

在 `main()` 中完成以下测试：

```go
func main() {
    // 1. 创建 BaseCounter，测试 MutableCounter 接口
    //    - 创建、Add、Get、Reset、String
    
    // 2. 将 *BaseCounter 赋值给 Counter 接口变量（测试自动解引用）
    //    注意：这里用指针赋值给只读接口
    
    // 3. 尝试将 BaseCounter（值类型）赋值给 MutableCounter
    //    观察编译错误，理解为什么
    
    // 4. 使用 ProcessCounter 函数
    //    - 传入 MutableCounter（应该能识别并修改）
    //    - 传入只读的 Counter（类型断言失败）
    
    // 5. 创建多个计数器，使用 SumCounters 计算总和
    
    // 6. 使用 FindByName 查找计数器
    //    - 使用类型断言将找到的 Counter 转回 MutableCounter
    //    - 修改其值
    
    // 7. 测试 Registry
    //    - 注册 BaseCounter 和 SafeCounter 的工厂函数
    //    - 通过 Registry 创建计数器
    
    // 8. 使用 type switch 处理不同类型的 Counter
    //    遍历计数器列表，根据实际类型打印不同信息
}
```

---

## 核心考察点

### 薄弱环节强化

1. **接口组合** - `MutableCounter` 嵌入 `Counter`
2. **类型断言零值** - 安全断言处理失败情况
3. **标准库接口** - 正确实现 `fmt.Stringer`（方法名是 `String`）

### 指针接收者应用

4. **为什么用指针接收者？**
   - `Add()` 和 `Reset()` 需要修改内部状态 `value`
   - 值接收者会复制结构体，修改的是副本

5. **值类型 vs 指针类型的接口实现**
   - `*BaseCounter` 实现了 `MutableCounter`
   - `BaseCounter`（值）**不实现** `MutableCounter`

6. **自动解引用**
   - `*BaseCounter` 可以赋值给 `Counter`（调用值接收者方法 `Get()`）

---

## 运行要求

```bash
go run main.go
```

---

## 提示

- 指针接收者语法：`func (c *BaseCounter) Add(n int)`
- 类型断言：`mc, ok := c.(MutableCounter)`
- 类型 switch：`switch v := c.(type) { case *BaseCounter: ... }`
- 接口组合：直接写接口名嵌入，不需要写方法
- `fmt.Stringer` 接口方法名是 `String`，不是 `Stringer`
