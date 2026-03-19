# 变量与常量基础

## 一、变量（Variable）

变量是存储数据的基本单元，Go 是**静态类型语言**，变量必须先声明后使用。

### 1.1 变量声明方式

#### 方式一：标准声明

```go
var name string
var age int
```

> 注意：Go 的类型声明位于变量名**之后**，与 C++/Java 不同。

#### 方式二：声明并初始化

```go
var name string = "Go"
var age int = 14
```

#### 方式三：类型推导

当提供初始值时，Go 会自动推断类型：

```go
var name = "Go"    // 推导为 string
var age = 14       // 推导为 int
var pi = 3.14      // 推导为 float64
```

#### 方式四：短变量声明（推荐，仅函数内使用）

```go
name := "Go"
age := 14
```

> ⚠️ `:=` 只能在**函数内部**使用，不能用于声明全局变量。

### 1.2 多变量声明

```go
// 同类型多变量
var a, b, c int = 1, 2, 3

// 不同类型多变量（类型推导）
var name, age = "Go", 14

// 短变量声明多变量
x, y := 10, 20

// 分组声明（常用于全局变量）
var (
    name    string = "Go"
    version int    = 1
    isOpen  bool   = true
)
```

### 1.3 零值（Zero Value）

未初始化的变量会自动赋予**零值**：

| 类型 | 零值 |
|------|------|
| 数值类型（int, float等） | `0` |
| 布尔类型 | `false` |
| 字符串类型 | `""`（空字符串）|
| 指针、切片、map、channel、函数 | `nil` |

```go
var i int      // i = 0
var b bool     // b = false
var s string   // s = ""
```

### 1.4 变量使用规则

1. **先声明后使用**：未声明的变量无法使用
2. **必须使用**：声明的变量必须使用，否则编译报错
3. **不能重复声明**：同一作用域内不能重复声明同名变量
4. **`:=` 的特殊规则**：左侧至少有一个新变量，否则视为赋值（⚠️易错）

```go
x := 10
x := 20      // 错误！x 已声明，应使用 x = 20

x, y := 10, 20   // x 已声明，y 是新变量 → 合法（x 被赋值，y 被声明）
```

---

## 二、常量（Constant）

常量是编译期确定且不可修改的值。

### 2.1 常量声明

```go
const Pi = 3.14159
const MaxSize = 100
```

关键点：`const` 不指定类型得到的是无类型常量（untyped constant），可以灵活赋值（类似于隐式转换）

### 2.2 多常量声明

```go
const (
    Monday    = 1
    Tuesday   = 2
    Wednesday = 3
)
```

### 2.3 iota 枚举

`iota` 是 Go 的常量计数器，在 `const` 块中从 0 开始递增：

```go
const (
    Red   = iota  // 0
    Green         // 1
    Blue          // 2
)
```

⚠️ const 后没有等号（因为等号在里面有了）
⚠️ const 块内部没有逗号（区别于 struct 字面量）

常见用法：

```go
// 使用下划线跳过某些值
const (
    _     = iota      // 0，跳过
    KB    = 1 << iota // 1 << 1 = 2 
    MB                // 1 << 2 = 4 
    GB                // 1 << 3 = 8
)

// 实际正确的写法
const (
    _  = iota
    KB = 1 << (10 * iota) // 1 << 10 = 1024
    MB = 1 << (10 * iota) // 1 << 20 = 1048576
    GB = 1 << (10 * iota) // 1 << 30 = 1073741824
)
```

### 2.4 常量规则

1. **必须初始化**：常量声明时必须赋值
2. **值不可变**：运行时不能修改
3. **只能是基本类型**：数值、字符串、布尔值
4. **不能用变量或函数返回值初始化**

```go
const a = 10           // ✓ 合法，可在编译期确定值
const b = getValue()   // ✗ 非法，不能用函数返回值
const c = someVar      // ✗ 非法，不能用变量
```

---

## 三、变量 vs 常量

| 特性 | 变量 | 常量 |
|------|------|------|
| 关键字 | `var` / `:=` | `const` |
| 值是否可变 | ✓ 可变 | ✗ 不可变 |
| 是否必须初始化 | ✗ 否（有零值） | ✓ 是 |
| 作用域 | 包级/局部 | 包级/局部 |
| 能否用函数初始化 | ✓ 能 | ✗ 不能 |

---

## 四、代码示例

```go
package main

import "fmt"

// 全局变量
var globalName = "Go"

// 全局常量
const Version = "1.21"

// iota 枚举
const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main() {
    // 短变量声明
    name := "Gopher"
    age := 25
    
    // 多变量声明
    x, y := 10, 20
    
    // 交换变量
    x, y = y, x
    
    fmt.Println(name, age)
    fmt.Println("星期:", Sunday, Monday, Tuesday)
    fmt.Println("版本:", Version)
}
```
