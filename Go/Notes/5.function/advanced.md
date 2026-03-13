# 函数高级特性

## 1. 变参函数（可变参数）

Go 语言支持可变数量的参数，使用 `...` 语法。

### 基本语法

```go
func 函数名(参数名 ...类型) 返回值
```

### 示例

```go
// 计算任意数量整数的和
func sum(nums ...int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3))      // 6
    fmt.Println(sum(10, 20))       // 30
    fmt.Println(sum())             // 0
}
```

### 变参的本质

变参在函数内部实际上是**切片（slice）**：

```go
func printArgs(args ...int) {
    fmt.Printf("类型: %T\n", args)  // 类型: []int
    fmt.Printf("长度: %d\n", len(args))
}
```

### 传递切片给变参函数

使用 `...` 展开切片：

```go
nums := []int{1, 2, 3, 4}
result := sum(nums...)  // 展开切片传递
```

### 混合参数

变参必须是最后一个参数：

```go
// 正确：变参在最后
func greet(prefix string, names ...string) {
    for _, name := range names {
        fmt.Printf("%s %s\n", prefix, name)
    }
}

// 错误：变参后面不能有其他参数
// func wrong(a ...int, b string)  // 编译错误
```

---

## 2. 匿名函数

匿名函数是没有名字的函数，可以直接定义和调用。

### 基本用法

```go
// 赋值给变量
add := func(a, b int) int {
    return a + b
}
fmt.Println(add(3, 4))  // 7

// 立即执行
result := func(a, b int) int {
    return a * b
}(3, 4)
fmt.Println(result)  // 12
```

### 作为参数传递

```go
func calculate(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func main() {
    // 传递匿名函数
    result := calculate(10, 5, func(x, y int) int {
        return x - y
    })
    fmt.Println(result)  // 5
}
```

### 作为返回值

```go
func makeMultiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    double := makeMultiplier(2)
    triple := makeMultiplier(3)
    
    fmt.Println(double(5))  // 10
    fmt.Println(triple(5))  // 15
}
```

---

## 3. 闭包（Closure）

闭包是引用了外部变量的匿名函数，这些变量会被闭包"捕获"并持续存在。

### 闭包的本质

```go
func makeCounter() func() int {
    count := 0  // 局部变量
    return func() int {
        count++  // 捕获并修改外部变量
        return count
    }
}

func main() {
    counter1 := makeCounter()
    counter2 := makeCounter()
    
    fmt.Println(counter1())  // 1
    fmt.Println(counter1())  // 2
    fmt.Println(counter1())  // 3
    
    fmt.Println(counter2())  // 1（独立的计数器）
}
```

> **关键点**：每个闭包都有自己的 `count` 变量副本，互不干扰。

### 闭包的注意事项

```go
func makeAdders() []func(int) int {
    var funcs []func(int) int
    for i := 0; i < 3; i++ {
        funcs = append(funcs, func(x int) int {
            return x + i  // 捕获的是变量 i 的引用
        })
    }
    return funcs
}

func main() {
    adders := makeAdders()
    fmt.Println(adders[0](10))  // 13（i 最终为 3）
    fmt.Println(adders[1](10))  // 13
    fmt.Println(adders[2](10))  // 13
}
```

**解决方法** - 使用局部变量：

```go
func makeAddersFixed() []func(int) int {
    var funcs []func(int) int
    for i := 0; i < 3; i++ {
        n := i  // 创建局部变量副本
        funcs = append(funcs, func(x int) int {
            return x + n  // 捕获 n 而不是 i
        })
    }
    return funcs
}
```

---

## 4. defer 延迟执行

`defer` 用于延迟函数的执行，直到外层函数返回前才执行。

### 基本用法

```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
    // 输出: hello
    //       world
}
```

### 多个 defer 的执行顺序

**后进先出**（LIFO - Last In First Out）：

```go
func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    fmt.Println("done")
    // 输出: done
    //       3
    //       2
    //       1
}
```

### defer 的典型用途

资源清理（确保释放）：

```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // 确保文件关闭
    
    // 处理文件...
    return nil
}
```

### defer 的参数求值

defer 语句中的参数在定义时立即求值：

```go
func main() {
    i := 0
    defer fmt.Println(i)  // 输出 0（defer 定义时的值）
    i++
    fmt.Println(i)        // 输出 1
}
```

### defer 与命名返回值

defer 可以修改命名返回值：

```go
func increment() (result int) {
    defer func() {
        result++  // 修改命名返回值
    }()
    return 0  // 实际返回 1
}
```

---

## 5. 小结

| 特性 | 说明 | 典型场景 |
|------|------|----------|
| 变参函数 | `...类型` 语法，本质是切片 | 格式化输出、求和/求平均 |
| 匿名函数 | 无函数名，可立即执行 | 回调函数、简单逻辑封装 |
| 闭包 | 捕获外部变量的匿名函数 | 计数器、缓存、工厂函数 |
| defer | 延迟执行，LIFO 顺序 | 资源释放、日志记录 |

---

## 练习题（思考）

1. 以下代码的输出是什么？

```go
func main() {
    for i := 0; i < 3; i++ {
        defer fmt.Print(i)
    }
}
```

2. 闭包和普通函数有什么区别？什么时候应该使用闭包？

3. 以下代码有什么问题？如何修复？

```go
func process(items []int) {
    for _, item := range items {
        defer func() {
            fmt.Println(item)
        }()
    }
}
```
