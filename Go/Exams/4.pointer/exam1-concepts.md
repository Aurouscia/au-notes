# 指针概念题

## 一、填空题

1. 在 Go 中，`&x` 表示获取变量 x 的________。

2. 如果 `p` 是一个指针变量，`*p` 表示________。

3. 声明一个指向 `int` 类型的指针变量，应该写为________。

4. 指针的零值是________。

5. 函数 `new(int)` 返回的类型是________。

## 二、判断题（正确填✓，错误填✗）

1. Go 语言支持指针运算（如 `p++`）。

2. 对 nil 指针进行解引用操作会导致编译错误。

3. 使用指针作为函数参数可以避免大结构体的拷贝，提高效率。

4. 方法接收者为指针类型时，方法内部可以修改原对象的字段。

5. `*int` 和 `*float64` 是相同的指针类型。

## 三、代码阅读题

请写出以下代码的输出结果：

### 题目 1

```go
package main

import "fmt"

func main() {
    x := 10
    p := &x
    *p = 20
    fmt.Println(x)
}
```

输出：________

### 题目 2

```go
package main

import "fmt"

func modify(y *int) {
    *y = *y + 10
}

func main() {
    x := 5
    modify(&x)
    fmt.Println(x)
}
```

输出：________

### 题目 3

```go
package main

import "fmt"

func main() {
    var p *int
    if p == nil {
        fmt.Println("nil")
    }
    fmt.Println("end")
}
```

输出：________

### 题目 4

```go
package main

import "fmt"

func main() {
    a := 1
    b := 2
    p1 := &a
    p2 := &b
    p1 = p2
    *p1 = 100
    fmt.Println(a)
    fmt.Println(b)
}
```

输出：
- a = ________
- b = ________

## 四、简答题

1. 简述在什么情况下应该使用指针作为函数参数？



2. 值接收者方法和指针接收者方法有什么区别？什么时候应该选择指针接收者？



3. 以下代码有什么问题？请说明原因。

```go
func getPointer() *int {
    x := 10
    return &x
}
```


