# 结构体概念题

## 一、填空题

1. 定义结构体使用关键字________。

2. 结构体字段名以大写字母开头表示该字段是________（导出/未导出）的。

3. 值接收者方法接收的是结构体的________，指针接收者方法接收的是结构体的________。

4. 使用 `new(Person)` 返回的是________类型（值/指针）。

5. Go 语言通过________（继承/嵌入）实现代码复用。

## 二、判断题（正确填✓，错误填✗）

1. 结构体中所有字段类型必须相同。

2. 指针接收者方法可以修改原结构体的字段值。

3. 嵌入结构体的字段可以直接通过外层结构体变量访问。

4. 包含切片的结构体可以使用 `==` 进行比较。

5. `struct{}` 空结构体不占用内存空间。

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
- Width, Height = ________

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
- d.Speak() 输出：________

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

## 四、简答题

1. 简述值接收者方法和指针接收者方法的区别，以及在什么情况下应该选择指针接收者？



2. 结构体嵌入（embedding）有什么作用？请举例说明如何使用。



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


