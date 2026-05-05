# 复习 1：流程控制概念题（针对薄弱环节）

请阅读题目后直接写出答案。

---

## 第1题：for-range 遍历字符串

以下代码的输出是什么？请写出 `i` 和 `r` 的每一组值。

```go
func main() {
    s := "Hi世界"
    for i, r := range s {
        fmt.Printf("%d:%c ", i, r)
    }
}
```

- 输出：
- 解释为什么 `i` 的值不连续：

---

## 第2题：switch 与 fallthrough

以下代码的输出是什么？

```go
func main() {
    score := 85
    grade := ""
    switch {
    case score >= 90:
        grade = "A"
    case score >= 80:
        grade = "B"
        fallthrough
    case score >= 60:
        grade = "C"
    default:
        grade = "D"
    }
    fmt.Println(grade)
}
```

- 输出：
- 解释 `fallthrough` 的作用：

---

## 第3题：defer 执行顺序

以下代码的输出是什么？

```go
func main() {
    for i := 0; i < 3; i++ {
        defer fmt.Print(i, " ")
    }
    fmt.Println("done")
}
```

- 输出：
- 解释 defer 的执行时机和顺序：

---

## 第4题：标签与 break

以下代码的输出是什么？

```go
func main() {
Outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i == 1 && j == 1 {
                break Outer
            }
            fmt.Printf("(%d,%d) ", i, j)
        }
    }
}
```

- 输出：
- 如果不使用 `break Outer`，只写 `break`，输出会有什么不同？

---

## 第5题：for-range 与 goroutine（陷阱题）

以下代码有什么问题？如何修复？

```go
func main() {
    nums := []int{1, 2, 3}
    for _, n := range nums {
        go func() {
            fmt.Println(n)
        }()
    }
    time.Sleep(time.Second)
}
```

- 问题：
- 修复后的代码：

---

## 第6题：switch case 表达式

以下代码是否合法？如果合法，输出是什么？

```go
func main() {
    x := 5
    switch x {
    case 1, 3, 5:
        fmt.Println("奇数")
    case 2, 4, 6:
        fmt.Println("偶数")
    default:
        fmt.Println("其他")
    }
}
```

- 是否合法：
- 输出：

---

## 第7题：continue 与标签

以下代码的输出是什么？

```go
func main() {
    sum := 0
Outer:
    for i := 1; i <= 3; i++ {
        for j := 1; j <= 3; j++ {
            if i*j == 4 {
                continue Outer
            }
            sum += i * j
        }
    }
    fmt.Println(sum)
}
```

- 输出：
- 解释 `continue Outer` 的作用：
