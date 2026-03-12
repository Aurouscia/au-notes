# 指针基础

## 什么是指针

指针是一个变量，它存储的是另一个变量的**内存地址**。

## 核心操作符

### 1. `&` 取地址符

获取变量的内存地址：

```go
x := 42
addr := &x  // addr 存储的是 x 的内存地址
```

### 2. `*` 声明指针类型 / 解引用

- **声明指针类型**：`*T` 表示指向类型 T 的指针
- **解引用**：通过指针访问或修改指向的值

```go
var p *int    // 声明一个指向 int 的指针
x := 42
p = &x        // p 指向 x

fmt.Println(*p)  // 解引用，输出 42（x 的值）
*p = 100         // 通过指针修改 x 的值
fmt.Println(x)   // 输出 100
```

## 指针的零值

未初始化的指针值为 `nil`：

```go
var p *int
fmt.Println(p == nil)  // true

// 对 nil 指针解引用会导致 panic
// fmt.Println(*p)  // panic: invalid memory address
```

## new 函数

`new(T)` 分配内存并返回指向该内存的指针，内存被初始化为类型 T 的零值：

```go
p := new(int)    // *int，指向值为 0 的内存
*p = 42

person := new(Person)  // *Person，字段都是零值
person.Name = "Alice"
```

## 指针 vs 值传递

| 方式 | 特点 |
|------|------|
| 值传递 | 复制整个值，函数内修改不影响原变量 |
| 指针传递 | 只复制地址（8字节），函数内修改会影响原变量 |

```go
// 值传递 - 无法修改原变量
func modifyValue(x int) {
    x = 100
}

// 指针传递 - 可以修改原变量
func modifyPointer(x *int) {
    *x = 100
}

func main() {
    a := 10
    modifyValue(a)
    fmt.Println(a)  // 10（未改变）
    
    modifyPointer(&a)
    fmt.Println(a)  // 100（已改变）
}
```

## 方法接收者：值 vs 指针

```go
type Counter struct {
    value int
}

// 值接收者 - 不修改原对象
func (c Counter) Get() int {
    return c.value
}

// 指针接收者 - 修改原对象
func (c *Counter) Inc() {
    c.value++
}
```

## 常见使用场景

1. **修改函数参数**：需要在函数内修改外部变量
2. **避免大对象拷贝**：传递大型结构体时提高效率
3. **实现方法修改对象**：需要修改接收者的字段
4. **实现链表、树等数据结构**

## 注意事项

1. **nil 检查**：解引用前确保指针不为 nil
2. **不要返回局部变量的指针**：Go 会逃逸分析，但尽量避免
3. **指针不能进行算术运算**：Go 不支持 `p++` 这样的操作
