# 接口基础 - 概念题

## 一、填空题

1. Go 的接口实现是 ______ 的，不需要显式声明 `implements`。
    - 隐式

2. 空接口 `interface{}` 可以存储 ______ 类型的值。
    - 任意

3. 类型断言的语法是 `i.(T)`，其中 `i` 必须是 ______ 类型，`T` 是 ______ 类型。
    - 接口
    - 具体

4. 安全类型断言使用 `value, ok := i.(T)`，当断言失败时 `ok` 为 ______，`value` 为 ______。
    - false
    - nil
    - **❌有误，正确答案：T 的零值（如 int 是 0，string 是 ""，不是 nil）**

5. 类型选择（type switch）的语法是 `switch v := i.(type)`，其中 `.(type)` 只能在 ______ 中使用。
    - 类型选择

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
   - 对

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
    - 对
    - **❌有误，正确答案：错。`File` 没有实现 `Save()`，只有 `*File` 实现了。值类型不能自动取地址。**

3. 以下代码会输出 `true`：
   ```go
   var p *int = nil
   var i interface{} = p
   fmt.Println(i == nil)
   ```
    - 错，i 虽然装着空指针，但是它本身不为空（自身仍有一个类型指针和一个数据指针）

4. 以下类型断言是安全的，不会 panic：
   ```go
   var i interface{} = "hello"
   n := i.(int)
   _ = n
   ```
    - 错，有两个返回值的版本才是安全的，只有一个的版本，失败会 panic

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
    - 错，接口没有这种组合机制
    - **❌有误，正确答案：对。Go 支持接口组合（embedding），`ReadWriter` 自动包含 `Reader` 的所有方法。**

---

## 三、简答题

1. 请解释以下代码的输出，并说明原因：
   ```go
   var i interface{} = 42
   str, ok := i.(string)
   fmt.Printf("str=%q, ok=%v\n", str, ok)
   ```
    - str=nil, ok=false
    - 因为 i 内是数字，类型断言为字符串会失败
    - **❌有误，正确答案：str="", ok=false。断言失败时 `str` 是 string 类型的零值 `""`，不是 nil（string 不能为 nil）。**

2. 什么是接口的"隐式实现"？与 Java/C# 的显式实现相比有什么优缺点？
    - 只要某个类型实现某接口要求的所有方法，它即实现了该接口，无需 implements 指定
    - 优点：灵活，可在不修改原类型的情况下添加接口，接口和数据各自演进
    - 缺点：查找“实现某接口的类型”较为麻烦

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
    - 值变量无法使用指针接收者接收者方法
    - 将 String 方法的 *Person 的星号去掉，反正这个方法不需要修改 p

4. 请解释 `值接收者` 和 `指针接收者` 在接口实现时的区别，以及 Go 的自动解引用机制。
    - 值变量和指针变量都可以调用值接收者方法
    - 仅指针变量能调用指针接收者方法（值变量不行，因为“自动取地址”可能引起问题）
    - 自动解引用：当指针变量 p 调用值接收者方法时，无需 (*p)，go 会自动使用指针指向的值调用值接收者方法