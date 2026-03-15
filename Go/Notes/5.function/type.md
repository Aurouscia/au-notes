# Go 函数类型详解

## 1. 函数是一等公民

在 Go 中，函数是一等公民，可以：
- 作为参数传递
- 作为返回值
- 赋值给变量
- 存储在数据结构中

```go
// 函数类型变量
var add func(int, int) int
add = func(a, b int) int { return a + b }

// 作为参数
func apply(a, b int, op func(int, int) int) int {
    return op(a, b)
}

// 作为返回值
func makeMultiplier(n int) func(int) int {
    return func(x int) int { return x * n }
}
```

## 2. 函数类型的不变性（Invariance）

Go 的函数类型是**不变的（invariant）**，不是**协变的（covariant）**或**逆变的（contravariant）**。

这意味着：即使类型 `A` 实现了接口 `B`，`func() A` 也不能直接赋值给 `func() B`。

### 2.1 实际例子

```go
package main

import "fmt"

// 定义接口
type Animal interface {
    Speak() string
}

// Dog 实现了 Animal
type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "汪汪"
}

// 工厂函数类型
type AnimalFactory func(name string) Animal

// Dog 的构造函数 - 返回 *Dog
func NewDog(name string) *Dog {
    return &Dog{Name: name}
}

func main() {
    // ❌ 编译错误！
    // var factory AnimalFactory = NewDog
    // 错误: cannot use NewDog (type func(string) *Dog) as type AnimalFactory
    
    // 虽然 *Dog 实现了 Animal 接口
    // 但 func(string) *Dog 不能直接用做 func(string) Animal
    // 因为 Go 函数类型是不变的
    
    // ✅ 正确做法：使用匿名函数包装
    var factory AnimalFactory = func(name string) Animal {
        return NewDog(name)
    }
    
    animal := factory("大黄")
    fmt.Println(animal.Speak()) // 输出: 汪汪
}
```

### 2.2 为什么函数类型是不变的？

考虑以下场景：

```go
// 如果函数类型是协变的，以下代码应该能编译：
type AnimalFactory func() Animal
type DogFactory func() *Dog

var dogFactory DogFactory = func() *Dog { return &Dog{} }
var animalFactory AnimalFactory = dogFactory  // 假设这是合法的

// 那么我们可以这样做：
animal := animalFactory()  // 返回的是 *Dog，没问题

// 但如果反过来呢？
type Cat struct{}
func (c Cat) Speak() string { return "喵喵" }

// 如果允许协变，下面的类型转换就危险了：
animalFactory = func() Animal { return &Cat{} }  // 返回 Cat
dogFactory = animalFactory  // 如果允许，这就出问题了！
dog := dogFactory()  // 期望 *Dog，实际得到 *Cat
```

**结论**：为了保证类型安全，Go 选择让函数类型是不变的。

### 2.3 解决方案

#### 方案1：使用匿名函数包装（推荐）

```go
var factory AnimalFactory = func(name string) Animal {
    return NewDog(name)
}
```

#### 方案2：修改构造函数返回接口类型

```go
// 修改返回类型为接口
func NewDog(name string) Animal {
    return &Dog{Name: name}
}

// 现在可以直接赋值
var factory AnimalFactory = NewDog
```

#### 方案3：使用泛型（Go 1.18+）

```go
// 泛型工厂
func MakeFactory[T Animal](constructor func(string) T) AnimalFactory {
    return func(name string) Animal {
        return constructor(name)
    }
}
```

## 3. 函数类型的比较

两个函数类型相同，当且仅当：
- 参数个数相同
- 对应参数类型相同
- 返回值个数相同
- 对应返回值类型相同
- 参数名和返回值名不影响类型

```go
// 以下函数类型是相同的
func(a int, b string) bool
func(x int, y string) (ok bool)

// 以下函数类型是不同的
func(int, string) bool        // vs
func(string, int) bool        // 参数顺序不同

func(int) string              // vs
func(int) (string, error)     // 返回值数量不同
```

## 4. nil 函数

函数类型的零值是 `nil`，调用 nil 函数会 panic：

```go
var f func()
f()  // panic: runtime error: invalid memory address or nil pointer dereference

// 安全调用
if f != nil {
    f()
}
```

## 5. 函数值的大小

函数值在内部是一个指针（8 字节，64位系统），指向函数代码的入口地址。

```go
fmt.Println(unsafe.Sizeof(func(){}))  // 输出: 8
```
