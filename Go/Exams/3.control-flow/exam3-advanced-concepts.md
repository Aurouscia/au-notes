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
    - 0, 1, 2

2. `switch` 语句省略表达式时，case 中应该写 ______。
    - 布尔表达式 **✓ 正确**

3. 使用 `fallthrough` 时，下一个 case 的条件 ______（会被/不会被）判断。
    - 不会被（会无条件执行下一个 case 的语句体） **✓ 正确**

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
    - 0, 1, 3（遇到 2 时被跳过这次循环，遇到 4 时终止循环） **❌有误，正确答案：013（输出是连续的字符串，不是逗号分隔的列表）**

5. 带标签的 `break` 可以用于跳出 ______、______ 或 ______ 语句。
    - for **✓ 正确**
    - switch **✓ 正确**
    - select **✓ 正确**

---

## 二、判断题（正确填✓，错误填✗）

1. （ ）`for-range` 遍历 map 时，每次迭代的顺序是固定的。
    - 错，map 不保证顺序 **✓ 正确**

2. （ ）以下代码的输出是 `0 1 2`：
   ```go
   for i := 0; i < 3; i++ {
       defer fmt.Print(i)
   }
   ```
    - 对 **❌有误，正确答案：✗（错误）。defer 是在整个函数最后，后进先出，输出是 `210` 而不是 `0 1 2`**

3. （ ）`switch` 的 case 可以是范围表达式，如 `case x > 0 && x < 10:`。
    - 对 **✓ 正确（当 switch 省略表达式时，case 可以是布尔表达式）**

4. （ ）`fallthrough` 可以连续使用，让多个 case 依次执行。
    - 对 **✓ 正确**

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
    - 对 **❌有误，正确答案：✗（错误）。实际输出是 `01234`（没有空格）**

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
- 0:G
- 1:o
- 2:语
- 4:言（unicode 字节索引） **❌有误，正确答案：0:G 2:语 4:言（"Go语言"中，'G'占1字节，'o'占1字节，'语'占3字节，'言'占3字节，所以索引是 0, 1, 2, 5，而不是 0, 1, 2, 4）**

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
- 中 大（中输出了，大肯定也输出）

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
- 0 0 0 0 1 0
- （i 为 0，j 为 0 1 2，输出 3 个 0；i 为 1，j 为 0 1，输出 0 1；i 为 2，j 为 0）

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
- 1 2 3 4 5（中途修改不会改变遍历） **❌正确，但原因不完整。分析：在 Go 中，for-range 遍历切片时，range 表达式在循环开始前只被求值一次，确定遍历的长度，这里确定为 5 了，后面添加的元素没用**
---

## 四、简答题

1. 解释 `for-range` 遍历字符串和遍历切片的区别（从返回值的类型和含义说明）。
    - 遍历字符串时，第一返回值是字符在字符串中的字节索引 **✓ 正确**
    - 遍历切片时，第一返回值就是连续的序列索引 **✓ 正确**

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
    - 循环体内捕获的循环变量 i 其实是引用，三个 goroutine 输出的可能都是最终值 2 **✓ 正确。修复方法：将 i 作为参数传递给 goroutine：go func(i int) { fmt.Println(i) }(i)**

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
    - 
        ```go
        score := 85
        switch {
            case score >= 90:
                grade = "A"
            case score >= 80:
                grade = "B"
            case score >= 60:
                grade = "C"
            default:
                grade = "D"
        }
        ```
        **✓ 正确**