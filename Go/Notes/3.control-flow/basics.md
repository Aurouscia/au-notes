# 流程控制基础

Go 语言的流程控制结构包括条件语句、循环语句和跳转语句。Go 强调代码简洁，没有 `while`、`do-while`，所有循环都用 `for` 实现。

---

## 一、if-else 条件语句

### 1.1 基本语法

```go
if condition {
    // 条件为 true 时执行
} else if anotherCondition {
    // 另一个条件为 true 时执行
} else {
    // 以上条件都不满足时执行
}
```

> ⚠️ **重要**：Go 要求 `if` 的左大括号 `{` 必须与 `if` 在同一行，不能换行。

### 1.2 特殊写法：if 中定义变量

```go
if age := 20; age >= 18 {
    fmt.Println("已成年")
} else {
    fmt.Println("未成年")
}
// age 只在 if-else 块内有效
```

这种写法常用于错误处理：

```go
if file, err := os.Open("data.txt"); err != nil {
    log.Fatal(err)
} else {
    defer file.Close()
    // 处理文件
}
```

---

## 二、for 循环语句

Go 只有 `for` 一种循环关键字，但支持多种形式。

### 2.1 标准 for 循环（C 风格）

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

### 2.2 条件循环（类似 while）

```go
n := 0
for n < 10 {
    fmt.Println(n)
    n++
}
```

### 2.3 无限循环

```go
for {
    // 无限循环，需要用 break、return 或 os.Exit 退出
}
```

### 2.4 for-range 遍历

用于遍历数组、切片、字符串、map、channel：

```go
// 遍历切片
nums := []int{10, 20, 30}
for index, value := range nums {
    fmt.Printf("索引: %d, 值: %d\n", index, value)
}

// 只需要值时，用下划线忽略索引
for _, v := range nums {
    fmt.Println(v)
}

// 遍历 map
scores := map[string]int{"Go": 100, "Python": 90}
for key, value := range scores {
    fmt.Printf("%s: %d\n", key, value)
}

// 遍历字符串（按 Unicode 字符）
for i, ch := range "Go语言" {
    fmt.Printf("索引: %d, 字符: %c\n", i, ch)
}
```

> 💡 **注意**：`range` 遍历时，如果对象是值类型（如数组），会复制整个对象；建议遍历切片或指针数组。

---

## 三、switch 选择语句

### 3.1 基本语法

```go
switch expression {
case value1:
    // 匹配 value1
    case value2, value3:
    // 匹配 value2 或 value3
default:
    // 默认情况
}
```

**Go 的 switch 特点**：
- 自动 `break`，匹配成功后不会继续执行后续 case
- case 可以是表达式，不限于常量

```go
score := 85
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")  // 输出 B
case score >= 60:
    fmt.Println("C")  // 不会执行，自动 break
default:
    fmt.Println("D")
}
```

### 3.2 fallthrough 穿透

使用 `fallthrough` 强制执行下一个 case：

```go
n := 1
switch n {
case 1:
    fmt.Println("1")
    fallthrough
case 2:
    fmt.Println("2")  // 也会执行
    fallthrough
case 3:
    fmt.Println("3")  // 也会执行
default:
    fmt.Println("default")
}
// 输出: 1 2 3 default
```

---

## 四、跳转语句

### 4.1 break 跳出循环

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break  // 跳出循环
    }
    fmt.Println(i)
}
```

**带标签的 break**（跳出指定循环）：

```go
OuterLoop:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i*j > 2 {
            break OuterLoop  // 跳出外层循环
        }
        fmt.Printf("%d*%d=%d\n", i, j, i*j)
    }
}
```

### 4.2 continue 跳过本次

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue  // 跳过 i=2
    }
    fmt.Println(i)  // 输出 0, 1, 3, 4
}
```

**带标签的 continue**：

```go
OuterLoop:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if j == 1 {
            continue OuterLoop  // 跳到外层循环的下一次迭代
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}
```

### 4.3 goto 无条件跳转

```go
func main() {
    i := 0
Loop:
    if i < 5 {
        fmt.Println(i)
        i++
        goto Loop  // 跳转到 Loop 标签
    }
}
```

> ⚠️ **建议**：尽量避免使用 `goto`，除非用于错误处理或跳出深层嵌套。

---

## 五、流程控制对比表

| 语句 | 用途 | 特点 |
|------|------|------|
| `if-else` | 条件分支 | 可在条件中定义变量 |
| `for` | 循环 | Go 唯一的循环关键字 |
| `for-range` | 遍历集合 | 自动处理索引和值 |
| `switch` | 多分支选择 | 自动 break，可用 fallthrough |
| `break` | 跳出循环/switch | 可带标签 |
| `continue` | 跳过本次循环 | 可带标签 |
| `goto` | 无条件跳转 | 谨慎使用 |

---

## 六、代码示例

```go
package main

import "fmt"

func main() {
    // if-else 示例
    score := 85
    if score >= 90 {
        fmt.Println("优秀")
    } else if score >= 80 {
        fmt.Println("良好")
    } else {
        fmt.Println("继续努力")
    }

    // for 循环示例
    sum := 0
    for i := 1; i <= 100; i++ {
        sum += i
    }
    fmt.Println("1到100的和:", sum)

    // for-range 示例
    fruits := []string{"苹果", "香蕉", "橙子"}
    for i, fruit := range fruits {
        fmt.Printf("%d: %s\n", i, fruit)
    }

    // switch 示例
    day := "Monday"
    switch day {
    case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
        fmt.Println("工作日")
    case "Saturday", "Sunday":
        fmt.Println("周末")
    default:
        fmt.Println("未知")
    }
}
```
