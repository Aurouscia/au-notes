# 函数高级特性 - 概念题

## 一、填空题

1. 变参函数使用 ______ 符号定义，变参在函数内部实际上是 ______ 类型。

2. 将切片传递给变参函数时，需要使用 ______ 操作符展开切片。

3. 闭包是引用了 ______ 变量的匿名函数，这些变量会被闭包 ______ 并持续存在。

4. `defer` 语句在函数 ______ 时才执行，多个 `defer` 按照 ______ 的顺序执行。

5. `defer` 语句中的参数在 ______ 时立即求值，而不是在 ______ 时。

---

## 二、判断题（请在每题后写出判断结果和理由）

1. 以下代码可以正确计算任意数量整数的和：
   ```go
   func sum(nums ...int) int {
       total := 0
       for _, n := range nums {
           total += n
       }
       return total
   }
   ```

2. 变参函数可以有多个变参参数，例如 `func foo(a ...int, b ...string)`。

3. 以下代码会输出 `1 2 3`：
   ```go
   func main() {
       for i := 1; i <= 3; i++ {
           defer fmt.Print(i, " ")
       }
   }
   ```

4. 闭包捕获的外部变量在闭包执行完毕后会被垃圾回收。

5. 以下代码中，`counter1` 和 `counter2` 共享同一个计数器：
   ```go
   func makeCounter() func() int {
       count := 0
       return func() int {
           count++
           return count
       }
   }
   
   func main() {
       counter1 := makeCounter()
       counter2 := makeCounter()
       // ...
   }
   ```

---

## 三、简答题

1. 请解释以下代码的输出，并说明原因：
   ```go
   func main() {
       i := 0
       defer fmt.Println(i)
       i++
       fmt.Println(i)
   }
   ```

2. 闭包有什么实际应用场景？请举例说明。

3. 以下代码有什么问题？如何修复？
   ```go
   func makeFunctions() []func() {
       var funcs []func()
       for i := 0; i < 3; i++ {
           funcs = append(funcs, func() {
               fmt.Println(i)
           })
       }
       return funcs
   }
   
   func main() {
       for _, f := range makeFunctions() {
           f()
       }
   }
   ```
