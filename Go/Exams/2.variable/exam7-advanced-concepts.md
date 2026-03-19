# 考试 7：变量与常量进阶概念

## 一、填空题

1. 以下代码中，`a` 的类型是 ______：
   ```go
   const a = 10
   var b = a + 3.14
   ```
    - 无类型常量（可自动适应各种使用情况，自动转换）

2. 以下代码中，`c` 的类型是 ______：
   ```go
   var c = 10 + 3.14
   ```
    - float64（因为 c 是变量）

3. 短变量声明 `x, y := 1, 2` 后，再执行 `x, z := 3, 4`，此时 `x` 的值是 ______，`y` 的值是 ______。
    - 3
    - 2

4. `iota` 在下一个 `const` 块中会 ______（继续递增/重置为0）。
    - 重置为 0（同一个 const 块才是继续递增）

5. 以下代码的输出是 ______：
   ```go
   const (
       _ = iota
       KB = 1 << (10 * iota)
       MB
       GB
       TB
   )
   ```
   `KB` = ______，`MB` = ______，`GB` = ______
    - 1024
    - 1024 * 1024
    - 1024 * 1024 * 1024

---

## 二、判断题（正确填✓，错误填✗）

1. （ ）无类型常量可以赋值给任何兼容类型的变量。
    - 对

2. （ ）以下代码是合法的：
   ```go
   var x int = 10
   x := 20
   ```
    - 错（:= 左侧至少有一个新变量，此处没有）

3. （ ）以下代码会编译错误：
   ```go
   x, y := 1, 2
   x, y := 3, 4
   ```
    - 对（:= 左侧没有新变量）

4. （ ）`const Pi float64 = 3.14` 是有类型常量，不能赋值给 `float32` 变量。
    - 对

5. （ ）以下代码的输出是 `0 1 2`：
   ```go
   const (
       A = iota
       B = iota
       C = iota
   )
   ```
    - 对

---

## 三、代码分析题

### 第 1 题

```go
package main

import "fmt"

func main() {
    x := 10
    x, y := 20, 30
    fmt.Println(x, y)
}
```

输出是：________
- 20, 30

### 第 2 题

```go
package main

import "fmt"

func main() {
    const Pi = 3.14159
    var f32 float32 = Pi
    var f64 float64 = Pi
    fmt.Println(f32, f64)
}
```

这段代码会：________（编译错误/正常运行）
- 正常运行

### 第 3 题

```go
package main

import "fmt"

const (
    Monday = iota + 1
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
    Sunday
)

func main() {
    fmt.Println(Monday, Sunday)
}
```

输出是：________
- 1, 7

### 第 4 题

```go
package main

import "fmt"

func main() {
    var a = 10
    if true {
        a, b := 20, 30
        fmt.Println(a, b)
    }
    fmt.Println(a)
}
```

输出是：________
- 20, 30
- 10（内外不是一个作用域，互不影响）

---

## 四、简答题

1. 解释无类型常量（untyped constant）和有类型常量的区别，并各举一个例子。
    - 无类型常量是定义时没有指定类型的常量，可以赋值给任意兼容变量
        - const x = 1
        - var x0 int = x
    - 有类型常量只能赋值给同类型的变量，否则需要显式转换
        - const y int = 2
        - var y0 int = y

2. 以下代码有什么问题？如何修复？
   ```go
   func main() {
       x := 10
       if x > 5 {
           x, y := 20, 30
           fmt.Println(x, y)
       }
       fmt.Println(y)  // 错误
   }
   ```
    - y 在 if 语句块内部定义，外部无法使用
    - 应该把 y 的定义提到外面，或者把打印放到里面

3. 使用 `iota` 定义一个位掩码（bitmask），表示文件权限：Read(4)、Write(2)、Execute(1)。
    ```go
    const = (
        Execute = 1 << iota
        Write
        Read
    )
    ```
    ❌ const 后是没有等于号的，行尾也没有逗号