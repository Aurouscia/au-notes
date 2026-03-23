# 考试 4：指针进阶概念

## 一、填空题

1. 以下代码的输出是 ______：
   ```go
   a := 10
   b := 20
   p1 := &a
   p2 := &b
   p1 = p2
   *p1 = 100
   fmt.Println(a, b)
   ```
    - 10 100 （p1 p2 都指向 b，修改 p1 只修改 b）

2. 函数 `new(int)` 返回的是 ______ 类型，它分配的内存位于 ______（栈/堆）。
    - *int
    - 堆  ❌有误：取决于编译器的逃逸分析。如果指针只在函数内用，一般分配在栈上；如果逃逸出函数（如被返回）则分配在堆上。

3. 以下代码的输出是 ______：
   ```go
   var p *int
   fmt.Println(p == nil)
   ```
    - true（未初始化的指针是 nil）

4. 以下代码会 ______（正常运行/panic）：
   ```go
   var p *int
   fmt.Println(*p)
   ```
    - panic（试图对一个没有指向的指针取值）

5. 以下代码的输出是 ______：
   ```go
   x := 5
   p := &x
   fmt.Println(*p == x)
   fmt.Println(p == &x)
   ```
    - true
    - true

---

## 二、判断题（正确填✓，错误填✗）

1. （ ）Go 语言中，指针的指针（`**int`）是合法的。
    - 对

2. （ ）以下代码会编译错误：
   ```go
   func getPtr() *int {
       x := 10
       return &x
   }
   ```
    - 错（这样没有问题，常用于构造函数）

3. （ ）`&` 运算符只能用于变量，不能用于常量或字面量。
    - 错（构造函数中常写 &SomeStruct{...}） ❌有误：`&`确实只能用于变量，`SomeStruct{...}`其实是创建了一个临时匿名变量（`&10`是非法的，但可用于复合字面量）

4. （ ）以下代码的输出是 `10 20`：
   ```go
   a, b := 10, 20
   swap(&a, &b)
   fmt.Println(a, b)
   
   func swap(x, y *int) {
       x, y = y, x
   }
   ```
    - 错（这个交换缺少 * 符号，仅仅交换了指针本身，没有修改底层值）

5. （ ）`unsafe.Pointer` 可以进行任意的指针类型转换。
    - 不知道 ❌有误：`unsafe.Pointer`可以表示任意类型的指针，允许进行任意的指针类型转换（但使用时需要非常小心）

---

## 三、代码分析题

### 第 1 题

```go
package main

import "fmt"

func modify(p *int) {
    *p = 100
    p = nil
}

func main() {
    x := 10
    p := &x
    modify(p)
    fmt.Println(x, p == nil)
}
```

输出是：________
- 100 false

### 第 2 题

```go
package main

import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    p := &arr[0]
    fmt.Println(*p)
    p++
    fmt.Println(*p)
}
```

这段代码会：________（编译错误/正常运行，输出是什么）
- 编译错误（不能进行指针运算）

### 第 3 题

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Birthday() {
    p.Age++
}

func (p *Person) RealBirthday() {
    p.Age++
}

func main() {
    person := Person{Name: "Alice", Age: 20}
    person.Birthday()
    fmt.Println(person.Age)
    person.RealBirthday()
    fmt.Println(person.Age)
}
```

输出是：________
- 20（值传递，没有修改原 struct）
- 21

### 第 4 题

```go
package main

import "fmt"

func main() {
    s := []int{1, 2, 3}
    p := &s[1]
    *p = 100
    fmt.Println(s)
    s = append(s, 4)
    *p = 200
    fmt.Println(s)
}
```

输出是：________
- 1 100 3
- 1 100 3 4

---

## 四、简答题

1. 解释以下代码中 `p` 的变化，并说明最终输出：
   ```go
   x, y := 1, 2
   p := &x
   q := &y
   r := p
   p = q
   *r = 10
   fmt.Println(x, y, *p)
   ```
    - 一开始 p 指向 x
    - 然后 p 改为指向 y
    - 输出为 10 2 2

2. 以下代码有什么问题？如何修复？
   ```go
   func appendValue(s []int, v int) {
       s = append(s, v)
   }
   
   func main() {
       s := []int{1, 2, 3}
       appendValue(s, 4)
       fmt.Println(s)
   }
   ```
    - appendValue 函数中 s 被替换，没有修改 main 中的 s
    - 修复：改为 `appendValue(s *[]int, v int)`，调用时 `appendValue(&s, 4)`

3. 编写一个函数 `swap`，交换两个 `int` 变量的值，要求使用指针实现。
    ```go
    func swap(x *int, y *int){
        *x, *y = *y, *x
    }
    ```