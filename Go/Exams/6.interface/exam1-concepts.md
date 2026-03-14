# 接口基础 - 概念题

## 一、填空题

1. Go 的接口实现是 ______ 的，不需要显式声明 `implements`。

2. 空接口 `interface{}` 可以存储 ______ 类型的值。

3. 类型断言的语法是 `i.(T)`，其中 `i` 必须是 ______ 类型，`T` 是 ______ 类型。

4. 安全类型断言使用 `value, ok := i.(T)`，当断言失败时 `ok` 为 ______，`value` 为 ______。

5. 类型选择（type switch）的语法是 `switch v := i.(type)`，其中 `.(type)` 只能在 ______ 中使用。

---

## 二、判断题（请写出判断结果和理由）

1. 以下代码可以正确编译和运行：
   ```go
   type Speaker interface {
       Speak() string
   }
   
   type Dog struct{}
   
   func (d Dog) Speak() string {
       return "汪汪"
   }
   
   func main() {
       var d Dog
       var s Speaker = &d
       fmt.Println(s.Speak())
   }
   ```

2. 以下代码可以正确编译：
   ```go
   type Saver interface {
       Save()
   }
   
   type File struct{}
   
   func (f *File) Save() {}
   
   func main() {
       var f File
       var s Saver = f
       _ = s
   }
   ```

3. 以下代码会输出 `true`：
   ```go
   var p *int = nil
   var i interface{} = p
   fmt.Println(i == nil)
   ```

4. 以下类型断言是安全的，不会 panic：
   ```go
   var i interface{} = "hello"
   n := i.(int)
   _ = n
   ```

5. 以下代码可以正确编译：
   ```go
   type Reader interface {
       Read([]byte) (int, error)
   }
   
   type ReadWriter interface {
       Reader
       Write([]byte) (int, error)
   }
   ```

---

## 三、简答题

1. 请解释以下代码的输出，并说明原因：
   ```go
   var i interface{} = 42
   str, ok := i.(string)
   fmt.Printf("str=%q, ok=%v\n", str, ok)
   ```

2. 什么是接口的"隐式实现"？与 Java/C# 的显式实现相比有什么优缺点？

3. 以下代码有什么问题？如何修复？
   ```go
   type Stringer interface {
       String() string
   }
   
   type Person struct {
       Name string
   }
   
   func (p *Person) String() string {
       return p.Name
   }
   
   func main() {
       p := Person{Name: "Alice"}
       var s Stringer = p
       fmt.Println(s.String())
   }
   ```

4. 请解释 `值接收者` 和 `指针接收者` 在接口实现时的区别，以及 Go 的自动解引用机制。
