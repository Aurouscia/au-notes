# 结构体（struct）

## 什么是结构体

结构体是一种**自定义数据类型**，用于将多个不同类型的字段组合成一个整体，用于描述复杂的实体。

## 定义结构体

```go
type Person struct {
    Name string
    Age  int
}
```

## 初始化方式

### 1. 零值初始化

```go
var p Person  // Name="", Age=0
```

### 2. 字段名初始化（推荐）

```go
p := Person{
    Name: "Alice",
    Age:  30,
}
```

### 3. 位置初始化（需按字段顺序）

```go
p := Person{"Bob", 25}
```

### 4. 指针初始化

```go
// 使用 new
p := new(Person)  // *Person，字段为零值
p.Name = "Charlie"

// 使用取地址符
p := &Person{Name: "David", Age: 35}
```

## 访问和修改字段

使用点号 `.` 访问：

```go
p := Person{Name: "Alice", Age: 30}
fmt.Println(p.Name)  // 读取
p.Age = 31           // 修改
```

## 结构体方法

### 值接收者方法

不修改原结构体：

```go
func (p Person) GetInfo() string {
    return p.Name + " - " + strconv.Itoa(p.Age)
}
```

### 指针接收者方法

可以修改原结构体：

```go
func (p *Person) Birthday() {
    p.Age++
}
```

## 结构体嵌入（组合）

Go 没有继承，使用嵌入实现代码复用：

```go
type Address struct {
    City   string
    Street string
}

type Employee struct {
    Name    string
    Address // 嵌入结构体（匿名字段）
}

func main() {
    e := Employee{
        Name: "Tom",
        Address: Address{
            City:   "Beijing",
            Street: "Main St",
        },
    }
    
    // 可以直接访问嵌入字段的成员
    fmt.Println(e.City)        // Beijing
    fmt.Println(e.Address.City) // 也可以这样写
}
```

## 结构体比较

- 如果所有字段都是可比较类型，结构体可以使用 `==` 比较
- 包含切片、map、函数等不可比较类型的结构体不能用 `==`

```go
type Point struct {
    X, Y int
}

p1 := Point{1, 2}
p2 := Point{1, 2}
fmt.Println(p1 == p2)  // true
```

## 注意事项

1. **字段导出**：大写字母开头的字段是导出的（public），小写是未导出的（private）
2. **标签（Tag）**：可用于反射，如 JSON 序列化
   ```go
   type User struct {
       Name string `json:"name"`
       Age  int    `json:"age"`
   }
   ```
3. **空结构体**：`struct{}` 不占用内存，常用于信号通道
