# 考试 3：流程控制进阶概念

## 一、填空题

1. 以下代码的输出中，`i` 的值依次是 ______：
   ```go
   s := "Hello, 世界"
   for i, r := range s {
       fmt.Printf("%d:%c ", i, r)
   }
   ```
   （写出前 3 个 i 的值即可）

2. `switch` 语句省略表达式时，case 中应该写 ______。

3. 使用 `fallthrough` 时，下一个 case 的条件 ______（会被/不会被）判断。

4. 以下代码的输出是 ______：
   ```go
   for i := 0; i < 5; i++ {
       if i == 2 {
           continue
       }
       if i == 4 {
           break
       }
       fmt.Print(i)
   }
   ```

5. 带标签的 `break` 可以用于跳出 ______、______ 或 ______ 语句。

---

## 二、判断题（正确填✓，错误填✗）

1. （ ）`for-range` 遍历 map 时，每次迭代的顺序是固定的。

2. （ ）以下代码的输出是 `0 1 2`：
   ```go
   for i := 0; i < 3; i++ {
       defer fmt.Print(i)
   }
   ```

3. （ ）`switch` 的 case 可以是范围表达式，如 `case x > 0 && x < 10:`。

4. （ ）`fallthrough` 可以连续使用，让多个 case 依次执行。

5. （ ）以下代码会输出 `0 1 2 3 4`：
   ```go
   i := 0
   for i < 5 {
       fmt.Print(i)
       i++
       if i == 3 {
           continue
           i++  // 这行不会执行
       }
   }
   ```

---

## 三、代码分析题

### 第 1 题

```go
package main

import "fmt"

func main() {
    s := "Go语言"
    for i, r := range s {
        fmt.Printf("%d:%c ", i, r)
    }
}
```

输出是：________

### 第 2 题

```go
package main

import "fmt"

func main() {
    n := 5
    switch {
    case n > 0 && n < 3:
        fmt.Print("小 ")
        fallthrough
    case n > 3 && n < 6:
        fmt.Print("中 ")
        fallthrough
    case n > 6:
        fmt.Print("大 ")
    default:
        fmt.Print("其他")
    }
}
```

输出是：________

### 第 3 题

```go
package main

import "fmt"

func main() {
Outer:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i*j >= 2 {
                continue Outer
            }
            fmt.Printf("%d ", i*j)
        }
    }
}
```

输出是：________

### 第 4 题

```go
package main

import "fmt"

func main() {
    nums := []int{1, 2, 3, 4, 5}
    for i := range nums {
        if i == 2 {
            nums = append(nums, 6)
        }
        fmt.Print(nums[i], " ")
    }
}
```

输出是：________

---

## 四、简答题

1. 解释 `for-range` 遍历字符串和遍历切片的区别（从返回值的类型和含义说明）。

2. 以下代码有什么问题？如何修复？
   ```go
   func main() {
       for i := 0; i < 3; i++ {
           go func() {
               fmt.Println(i)
           }()
       }
   }
   ```

3. 使用 `switch` 语句重写以下 `if-else` 代码：
   ```go
   score := 85
   if score >= 90 {
       grade = "A"
   } else if score >= 80 {
       grade = "B"
   } else if score >= 60 {
       grade = "C"
   } else {
       grade = "D"
   }
   ```
