# 复习 1：变量与常量概念题（针对薄弱环节）

请阅读题目后直接写出答案。

---

## 第1题：无类型常量

以下代码中，`a` 和 `b` 的类型分别是什么？`c` 的类型又是什么？

```go
const a = 10          // a 的类型是？
var b = a + 3.14      // b 的类型是？
const c float64 = 10  // c 的类型是？
```

- `a` 的类型：
- `b` 的类型：
- `c` 的类型：

---

## 第2题：短变量声明的复用规则

以下代码是否合法？如果合法，`x` 和 `y` 的最终值是多少？

```go
func main() {
    x, y := 10, 20
    x, z := 30, 40
    fmt.Println(x, y, z)
}
```

- 是否合法：
- 最终输出：

---

## 第3题：iota 与位掩码

使用 `iota` 定义文件权限位掩码，要求：
- `Execute` = 1（可执行）
- `Write` = 2（可写）
- `Read` = 4（可读）

请写出 `const` 块的完整代码：

```go
const (
    // 请在这里填写
)
```

---

## 第4题：值接收者 vs 指针接收者

以下代码的输出是什么？为什么？

```go
type Counter struct {
    count int
}

func (c Counter) Increment() {
    c.count++
}

func (c *Counter) Add(n int) {
    c.count += n
}

func main() {
    c := Counter{count: 5}
    c.Increment()
    fmt.Println(c.count)
    c.Add(3)
    fmt.Println(c.count)
}
```

- 第一个 `Println` 输出：
- 第二个 `Println` 输出：
- 原因：

---

## 第5题：结构体比较

以下代码能否编译通过？如果不能，原因是什么？

```go
type User struct {
    Name    string
    Age     int
    Tags    []string
}

func main() {
    u1 := User{Name: "Alice", Age: 25, Tags: []string{"a", "b"}}
    u2 := User{Name: "Alice", Age: 25, Tags: []string{"a", "b"}}
    fmt.Println(u1 == u2)
}
```

- 能否编译通过：
- 原因：

---

## 第6题：嵌入结构体

以下代码的输出是什么？

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println(a.Name, "says hello")
}

type Dog struct {
    Animal
    Breed string
}

func main() {
    d := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden"}
    fmt.Println(d.Name)
    d.Speak()
}
```

- `fmt.Println(d.Name)` 输出：
- `d.Speak()` 输出：

---

## 第7题：new() 与指针

以下代码的输出是什么？

```go
func main() {
    p := new(int)
    *p = 100
    fmt.Println(*p)
    
    q := new(Counter)
    q.count = 10
    fmt.Println(q.count)
}
```

- 第一个 `Println` 输出：
- 第二个 `Println` 输出：
- `new(int)` 返回的类型是：
