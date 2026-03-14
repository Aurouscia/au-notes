# 接口基础 - 实践题

## 需求

请实现一个简单的图形计算程序，使用接口来统一处理不同类型的图形。

### 1. 定义接口

定义 `Shape` 接口，包含以下方法：

- `Area() float64` - 计算面积
- `Perimeter() float64` - 计算周长

### 2. 实现具体图形

实现以下图形类型，它们都满足 `Shape` 接口：

**Rectangle（矩形）**
```go
type Rectangle struct {
    Width  float64
    Height float64
}
```

**Circle（圆形）**
```go
type Circle struct {
    Radius float64
}
```

**Triangle（三角形）**
```go
type Triangle struct {
    A, B, C float64 // 三条边的长度
}
```

### 3. 实现工具函数

实现以下函数：

```go
// PrintShapeInfo 打印图形的详细信息
func PrintShapeInfo(s Shape) {
    // 输出格式：
    // 图形: [类型名]
    // 面积: [值]
    // 周长: [值]
}

// TotalArea 计算多个图形的总面积
func TotalArea(shapes ...Shape) float64

// TotalPerimeter 计算多个图形的总周长
func TotalPerimeter(shapes ...Shape) float64
```

### 4. 实现 fmt.Stringer 接口

为每个图形类型实现 `fmt.Stringer` 接口，使 `fmt.Println` 能输出友好的格式：

```go
// Rectangle 输出: "矩形(宽:10, 高:5)"
// Circle 输出: "圆形(半径:7)"
// Triangle 输出: "三角形(边A:3, 边B:4, 边C:5)"
```

### 5. 主函数测试

在 `main()` 中：

1. 创建几个不同类型的图形
2. 使用 `PrintShapeInfo` 分别打印它们的信息
3. 将所有图形放入一个 `[]Shape` 切片
4. 计算并输出总面积和总周长
5. 使用 `fmt.Println` 直接打印图形（测试 Stringer 接口）

---

## 运行要求

```bash
go run main.go
```

---

## 提示

- 圆的面积公式：`math.Pi * r * r`
- 圆的周长公式：`2 * math.Pi * r`
- 三角形面积可用海伦公式：
  - `s = (a + b + c) / 2`
  - `area = math.Sqrt(s * (s-a) * (s-b) * (s-c))`
- 使用 `fmt.Sprintf` 格式化字符串
- 使用 `reflect.TypeOf` 或 `fmt.Sprintf("%T", s)` 获取类型名
