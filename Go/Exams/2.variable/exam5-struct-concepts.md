# 结构体概念题

## 一、填空题

1. 定义结构体使用关键字________。
    - type 和 struct

2. 结构体字段名以大写字母开头表示该字段是________（导出/未导出）的。
    - 导出

3. 值接收者方法接收的是结构体的________，指针接收者方法接收的是结构体的________。
    - 值（副本）
    - 指针（指向其本身）

4. 使用 `new(Person)` 返回的是________类型（值/指针）。
    - 指针

5. Go 语言通过________（继承/嵌入）实现代码复用。
    - 嵌入

## 二、判断题（正确填✓，错误填✗）

1. 结构体中所有字段类型必须相同。
    - 错

2. 指针接收者方法可以修改原结构体的字段值。
    - 对

3. 嵌入结构体的字段可以直接通过外层结构体变量访问。
    - 对

4. 包含切片的结构体可以使用 `==` 进行比较。
    - 错

5. `struct{}` 空结构体不占用内存空间。
    - 对

## 三、代码阅读题

请写出以下代码的输出结果：

### 题目 1

```go
package main

import "fmt"

type Rectangle struct {
    Width  int
    Height int
}

func (r Rectangle) Area() int {
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor int) {
    r.Width *= factor
    r.Height *= factor
}

func main() {
    rect := Rectangle{Width: 2, Height: 3}
    fmt.Println(rect.Area())
    rect.Scale(2)
    fmt.Println(rect.Width, rect.Height)
}
```

输出：
- Area() = ________
    - 6
- Width, Height = ________
    - 4, 6

### 题目 2

```go
package main

import "fmt"

type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
    Animal
    Breed string
}

func main() {
    d := Dog{
        Animal: Animal{Name: "Buddy"},
        Breed:  "Golden Retriever",
    }
    fmt.Println(d.Name)
    d.Speak()
}
```

输出：
- d.Name = ________
    - Buddy
- d.Speak() 输出：________
    - Buddy makes a sound

### 题目 3

```go
package main

import "fmt"

type Counter struct {
    count int
}

func (c Counter) Value() int {
    return c.count
}

func (c Counter) Increment() {
    c.count++
}

func main() {
    c := Counter{count: 0}
    c.Increment()
    c.Increment()
    fmt.Println(c.Value())
}
```

输出：________
- 2

## 四、简答题

1. 简述值接收者方法和指针接收者方法的区别，以及在什么情况下应该选择指针接收者？
    - 值接收者方法会在调用时复制整个结构体，方法内无法改变原结构体的字段，只能修改副本
    - 指针接收者方法会在调用时传入原结构体的指针，方法内可以修改原结构体的字段
    - 需要修改原结构体字段时应该选择指针接收者


2. 结构体嵌入（embedding）有什么作用？请举例说明如何使用。
    - 可以将一个结构体的字段放入另一个结构体内，实现类似“继承”的效果
    - 例如：
        ```go
        type Dog struct {
            Animal
            Breed string
        }
        ```
        即可让 Dog 拥有 Animal 的字段


3. 以下代码有什么问题？如何修复？

```go
type User struct {
    Name    string
    Hobbies []string
}

func main() {
    u1 := User{Name: "Alice", Hobbies: []string{"reading", "coding"}}
    u2 := User{Name: "Alice", Hobbies: []string{"reading", "coding"}}
    fmt.Println(u1 == u2)
}
```
- 结构体内有切片字段，切片无法进行比较，所以结构体无法进行比较
