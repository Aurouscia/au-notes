# Go 接口 (Interface) - 基础

## 1. 什么是接口

接口是**方法签名的集合**，定义了一组行为契约。任何类型只要实现了接口中的所有方法，就**隐式地**实现了该接口。

```go
package main

import "fmt"

// 定义接口
type Speaker interface {
    Speak() string
}

// Dog 类型
type Dog struct {
    Name string
}

// Dog 实现 Speak 方法
func (d Dog) Speak() string {
    return "汪汪！"
}

// Cat 类型
type Cat struct {
    Name string
}

// Cat 实现 Speak 方法
func (c Cat) Speak() string {
    return "喵喵~"
}

func main() {
    // 接口变量可以存储任何实现了该接口的类型
    var s Speaker
    
    s = Dog{Name: "大黄"}
    fmt.Println(s.Speak()) // 输出: 汪汪！
    
    s = Cat{Name: "小白"}
    fmt.Println(s.Speak()) // 输出: 喵喵~
}
```

## 2. 隐式实现

Go 的接口实现是**隐式**的，不需要显式声明 `implements`：

```go
// 只要类型有 Speak() string 方法，它就自动实现了 Speaker 接口
// 不需要写: type Dog struct implements Speaker
```

**好处**：
- 解耦：接口定义和实现分离
- 灵活：可以在包外为已有类型添加接口实现

## 3. 空接口 interface{}

空接口不包含任何方法，因此**所有类型都实现了空接口**：

```go
var anything interface{}

anything = 42
anything = "hello"
anything = true
anything = struct{ X, Y int }{1, 2}

// 常用场景：存储任意类型的值
func Printf(format string, args ...interface{})  // fmt 包
func println(args ...interface{})                // 内置函数
```

## 4. 类型断言

**类型断言**的语法是 `i.(T)`，用于从接口值 `i` 中提取类型为 `T` 的具体值。

### 4.1 语法详解 `i.(T)`

| 组成部分 | 含义 |
|---------|------|
| `i` | **接口类型的变量**，必须是一个接口（如 `interface{}` 或自定义接口） |
| `.` | 类型断言运算符 |
| `(T)` | 要断言的目标类型 `T`，可以是具体类型或接口类型 |

**两种使用形式：**

```go
// 形式1：安全断言（推荐）
value, ok := i.(T)  // ok 为 true 表示断言成功，false 表示失败

// 形式2：强制断言（危险）
value := i.(T)      // 断言失败会直接 panic
```

### 4.2 工作原理

接口变量内部存储了 `(类型, 值)` 两部分信息：

```go
var i interface{} = "hello"
// i 内部: (类型: string, 值: "hello")

str, ok := i.(string)
// 1. 检查 i 存储的类型是否是 string → 是
// 2. 提取值 "hello" 赋给 str
// 3. ok = true

num, ok := i.(int)
// 1. 检查 i 存储的类型是否是 int → 否（实际是 string）
// 2. num 得到 int 的零值 0
// 3. ok = false
```

### 4.3 安全类型断言示例

```go
var i interface{} = "hello"

// 安全断言：即使失败也不会 panic
str, ok := i.(string)
if ok {
    fmt.Println("是字符串:", str)  // 输出: 是字符串: hello
} else {
    fmt.Println("不是字符串")
}

// 尝试断言为 int
num, ok := i.(int)
if ok {
    fmt.Println("是整数:", num)
} else {
    fmt.Println("不是整数")  // 输出: 不是整数
}
```

### 4.4 强制类型断言（不推荐）

```go
str := i.(string)  // 如果 i 不是 string，会 panic

// 示例：
var i interface{} = 42
str := i.(string)  // panic: interface conversion: interface {} is int, not string
```

**什么时候会 panic？**
- 接口值 `i` 为 `nil` 时
- 接口值 `i` 存储的实际类型不是 `T` 时

### 4.5 断言为接口类型

`T` 也可以是接口类型，用于检查 `i` 是否实现了另一个接口：

```go
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type ReadWriter interface {
    Reader
    Writer
}

var rw ReadWriter = // ...

// 检查 rw 是否实现了 Writer 接口
w, ok := rw.(Writer)  // ok = true，因为 ReadWriter 嵌入了 Writer
```

### 4.3 类型选择 (type switch)

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("整数: %d\n", v)
    case string:
        fmt.Printf("字符串: %s\n", v)
    case bool:
        fmt.Printf("布尔: %t\n", v)
    default:
        fmt.Printf("未知类型: %T\n", v)
    }
}

describe(42)      // 整数: 42
describe("hi")    // 字符串: hi
describe(3.14)    // 未知类型: float64
```

## 5. 接口值的数据结构

接口变量内部存储两个指针：
- **类型指针**：指向具体类型的元数据
- **数据指针**：指向具体值的内存地址

```go
var s Speaker = Dog{Name: "大黄"}
// s 内部: (类型: Dog, 数据: Dog{Name: "大黄"})
```

**重要**：接口值为 `nil` 和接口存储的值为 `nil` 是不同的！

```go
var p *Dog = nil
var s Speaker = p

fmt.Println(s == nil)  // false！
// s 内部: (类型: *Dog, 数据: nil)
// 接口本身不是 nil，只是存储了 nil 指针
```

## 6. 值接收者 vs 指针接收者

```go
type Printer interface {
    Print()
}

type Book struct {
    Title string
}

// 值接收者
func (b Book) Print() {
    fmt.Println(b.Title)
}

// 指针接收者
func (b *Book) Save() {
    // 保存到文件...
}
```

| 接收者类型 | 值变量 | 指针变量 |
|-----------|--------|---------|
| 值接收者 | ✅ 实现 | ✅ 实现（自动解引用） |
| 指针接收者 | ❌ 不实现 | ✅ 实现 |

```go
b := Book{Title: "Go 语言"}
p := &b

var printer Printer

printer = b  // ✅ Book 实现了 Print()
printer = p  // ✅ *Book 也实现了 Print()
```

### 6.1 自动解引用机制

当指针类型赋值给接口时，Go 会自动解引用：

```go
// Book 实现了 Print() 方法（值接收者）
// 这意味着：
// - Book 实现了 Printer
// - *Book 也实现了 Printer（自动解引用）

p := &Book{Title: "Go"}
var printer Printer = p  // ✅ 合法
// 内部逻辑：(*p).Print() 被调用时会自动解引用
```

**但反过来不行**：

```go
// *File 实现了 Save() 方法（指针接收者）
// 这意味着：
// - *File 实现了 Saver
// - File 不实现 Saver（无法自动取地址）

type Saver interface {
    Save()
}

type File struct {
    Data []byte
}

func (f *File) Save() {  // 指针接收者
    // 保存...
}

var f File = File{}
var s Saver = f  // ❌ 编译错误：File 没有实现 Save()
                 // 因为 f 是值，无法自动获取其地址
```

**为什么值不能自动取地址？**
- 值变量可能存储在只读的内存区域（如常量）
- 自动取地址会引入不确定性
- Go 语言设计上避免这种隐式操作

### 6.2 方法集总结

| 类型 | 方法集 |
|------|--------|
| `T`（值） | 所有 `(t T)` 接收者的方法 |
| `*T`（指针） | 所有 `(t T)` 和 `(t *T)` 接收者的方法 |

```go
type MyInt int

func (i MyInt) Add(n int) int   // 值接收者
func (i *MyInt) Mul(n int)      // 指针接收者

var a MyInt = 10
var b *MyInt = &a

a.Add(5)   // ✅
a.Mul(2)   // ❌ 编译错误

b.Add(5)   // ✅ 自动解引用
b.Mul(2)   // ✅
```

## 7. 常见标准库接口

### 7.1 error 接口

```go
type error interface {
    Error() string
}

// 自定义错误
type MyError struct {
    Msg string
    Code int
}

func (e MyError) Error() string {
    return fmt.Sprintf("[%d] %s", e.Code, e.Msg)
}
```

### 7.2 fmt.Stringer 接口

```go
type Stringer interface {
    String() string
}

// 实现后 fmt.Println 会调用它
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (%d岁)", p.Name, p.Age)
}

fmt.Println(Person{"Alice", 30})  // Alice (30岁)
```

## 8. 接口设计最佳实践

1. **接口要小**：方法越少越好，遵循"单一职责"
2. **隐式实现**：不要显式声明实现关系
3. **接受接口，返回具体类型**：函数参数用接口，返回值用具体类型
4. **避免空接口滥用**：`interface{}` 会丢失类型信息

```go
// 好的设计：小接口
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

// 接口组合
type ReadWriter interface {
    Reader
    Writer
}
```

## 9. 练习要点

- 定义一个接口并让一个结构体实现它
- 使用空接口存储不同类型的值
- 使用类型断言和 type switch 处理空接口
- 理解值接收者和指针接收者的区别
- 实现 `error` 或 `fmt.Stringer` 接口
