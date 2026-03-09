# 考试 1：流程控制概念题

## 一、填空题

1. Go 语言中唯一的循环关键字是 ______。
    - for

2. `for-range` 遍历切片时，第一个返回值是 ______，第二个返回值是 ______。
    - 索引
    - 值

3. Go 的 `switch` 语句中，默认每个 case 结束后会自动 ______（执行/跳出）。
    - 跳出

4. 使用 `______` 关键字可以让 switch 的 case 执行完后继续执行下一个 case。
    - fallthrough

5. `break` 语句后面可以添加 ______，用于跳出指定的外层循环。
    - 标签（for前面的标签）

6. 在 if 语句中定义的变量，其作用域是 ______。
    - if 的判断条件和大括号内

## 二、判断题（正确填✓，错误填✗）

1. （ ）Go 的 `for` 循环有三种形式：标准形式、条件形式、无限循环。
    - 对

2. （ ）以下代码是合法的：
   ```go
   if x > 0
   {
       fmt.Println("正数")
   }
   ```
    - 错：if 的左大括号必须与 if 关键字一行

3. （ ）`for-range` 遍历字符串时，返回的是字节索引和 Unicode 码点。
    - 错：索引应该是 unicode 字符的索引 ❌ **正确：返回的是字节索引（UTF-8 编码的字节位置）和 rune（Unicode 码点）**

4. （ ）`switch` 的 case 只能是常量值。
    - 对 ❌ **错，Go 的 switch case 可以是表达式，不限于常量**

5. （ ）以下代码会输出 0, 1, 2, 3：
   ```go
   for i := 0; i < 3; i++ {
       if i == 2 {
           continue
       }
       fmt.Println(i)
   }
   ```
    - 错，应该是 0，1，3 ❌ **应该是 0, 1（i=2 时 continue 跳过，i=3 时循环条件 i<3 不满足，循环结束）**

6. （ ）`goto` 语句可以跳转到任意位置，包括跳转到循环内部。
    - 错

## 三、代码分析题

### 第 1 题

```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30}
    for i, v := range nums {
        if i == 1 {
            break
        }
        fmt.Print(v, " ")
    }
}
```

输出结果是：
- 10 

### 第 2 题

```go
package main

import "fmt"

func main() {
    n := 5
    switch n {
    case 1, 3, 5:
        fmt.Print("奇数 ")
        fallthrough
    case 2, 4, 6:
        fmt.Print("偶数 ")
    default:
        fmt.Print("其他")
    }
}
```

输出结果是：
- 奇数 其他 ❌ **正确答案是：奇数 偶数（fallthrough 会强制执行下一个 case，不会跳到 default）**

### 第 3 题

```go
package main

import "fmt"

func main() {
OuterLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if i*j >= 2 {
                break OuterLoop
            }
            fmt.Printf("%d ", i*j)
        }
    }
}
```

输出结果是：
- 0 0 0 0 1 ✅ **正确！**

### 第 4 题

```go
package main

import "fmt"

func main() {
    sum := 0
    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            continue
        }
        sum += i
    }
    fmt.Println(sum)
}
```

输出结果是：
- 25

## 四、简答题

1. 简述 `for-range` 遍历切片时，如果只想要值而不需要索引，应该如何写？
    - for _, value range xxx，使用 `_` 丢弃值 ❌ **语法有误，正确写法：`for _, value := range xxx`**

2. 解释为什么以下代码会编译错误，并给出修正方法：
   ```go
   if score := 85; score >= 60 {
       fmt.Println("及格")
   }
   fmt.Println(score)  // 错误
   ```
    - 因为 score 在 if 中定义，而最后一个 score 超出其作用域

3. 简述带标签的 `break` 和 `continue` 的作用，并说明使用场景。
    - break 和 continue 带标签可以跳出指定循环，或者对指定循环开启下一轮
    - 例子：循环较深时，发生错误，逻辑无法继续进行，使用带标签的 break 直接跳出最外层循环
    - **补充：标签必须定义在 for、switch 或 select 代码块的上一行**