# 指针实践题

## 题目要求

请完成以下 3 个函数的实现，所有代码写在 `main.go` 中。

---

### 任务 1：交换两个整数

实现函数 `Swap`，接收两个指向整数的指针，交换它们指向的值。

```go
func Swap(a, b *int)
```

示例：
- 输入：a=10, b=20
- 调用后：a=20, b=10

---

### 任务 2：计算数组元素之和（指针遍历）

实现函数 `SumArray`，接收一个指向整数数组的指针，返回数组所有元素之和。

```go
func SumArray(arr *[5]int) int
```

示例：
- 输入：`&[5]int{1, 2, 3, 4, 5}`
- 输出：`15`

---

### 任务 3：使用 new 创建并初始化结构体

实现函数 `NewPerson`，使用 `new` 函数创建一个 `Person` 结构体指针，并设置其字段值，最后返回该指针。

```go
type Person struct {
    Name string
    Age  int
}

func NewPerson(name string, age int) *Person
```

示例：
- 输入：`"Alice", 30`
- 返回：`&Person{Name: "Alice", Age: 30}`

---

## 验证

在 `main` 函数中调用以上函数并打印结果，验证实现是否正确。
