# 函数基础

## 1. 函数定义

Go 语言使用 `func` 关键字定义函数。

### 基本语法

```go
func 函数名(参数列表) 返回值类型 {
    // 函数体
    return 返回值
}
```

### 示例

```go
// 无参数、无返回值
func sayHello() {
    fmt.Println("Hello, Go!")
}

// 有参数、有返回值
func add(a int, b int) int {
    return a + b
}

// 相邻参数类型相同可简写
func multiply(x, y int) int {
    return x * y
}
```

---

## 2. 参数传递

### 值传递

Go 语言默认使用**值传递**，函数内对参数的修改不会影响原变量。

```go
func modifyValue(x int) {
    x = 100  // 只修改了副本
}

func main() {
    a := 10
    modifyValue(a)
    fmt.Println(a)  // 输出: 10（原值未变）
}
```

### 指针传递

通过指针可以修改原变量的值。

```go
func modifyPointer(x *int) {
    *x = 100  // 通过指针修改原值
}

func main() {
    a := 10
    modifyPointer(&a)
    fmt.Println(a)  // 输出: 100（原值被修改）
}
```

---

## 3. 返回值

### 单返回值

```go
func square(x int) int {
    return x * x
}
```

### 多返回值

Go 语言支持函数返回多个值，这是 Go 的特色之一。

```go
// 返回商和余数
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    q, r := divide(17, 5)
    fmt.Printf("商: %d, 余数: %d\n", q, r)  // 商: 3, 余数: 2
}
```

### 命名返回值

可以给返回值命名，在函数体中直接使用，且 `return` 时可省略变量名（裸返回）。

```go
func rectangle(width, height float64) (area, perimeter float64) {
    area = width * height
    perimeter = 2 * (width + height)
    return  // 裸返回，自动返回 area 和 perimeter
}

func main() {
    a, p := rectangle(3, 4)
    fmt.Printf("面积: %.2f, 周长: %.2f\n", a, p)
}
```

> **注意**：命名返回值会被初始化为该类型的零值。

---

## 4. 忽略返回值

使用空白标识符 `_` 忽略不需要的返回值。

```go
func getCoordinates() (x, y, z int) {
    return 1, 2, 3
}

func main() {
    x, _, z := getCoordinates()  // 忽略 y
    fmt.Println(x, z)  // 输出: 1 3
}
```

---

## 5. 函数作为一等公民

Go 语言中函数是一等公民，可以：
- 赋值给变量
- 作为参数传递
- 作为返回值

```go
// 函数赋值给变量
add := func(a, b int) int {
    return a + b
}
result := add(3, 4)  // result = 7
```

---

## 6. 小结

| 特性 | 说明 |
|------|------|
| 定义 | `func` 关键字 |
| 参数 | 支持简写、值传递、指针传递 |
| 返回值 | 支持单返回值、多返回值、命名返回值 |
| 裸返回 | 命名返回值时可省略 return 后的变量 |
| 忽略返回值 | 使用 `_` 空白标识符 |

---

## 练习题（思考）

1. 以下代码输出什么？为什么？

```go
func swap(a, b int) (int, int) {
    a, b = b, a
    return a, b
}

func main() {
    x, y := 1, 2
    swap(x, y)
    fmt.Println(x, y)
}
```

2. 命名返回值和普通返回值各有什么优缺点？
