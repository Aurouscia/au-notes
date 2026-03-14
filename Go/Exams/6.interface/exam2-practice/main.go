package main

import (
	"fmt"
	"math"
)

// 请在此定义 Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 请在此定义 Rectangle、Circle、Triangle 结构体
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	A, B, C float64 // 三条边的长度
}

// 请实现各图形类型的 Area 和 Perimeter 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

// 请实现 PrintShapeInfo 函数
func PrintShapeInfo(s Shape) {
	// 输出格式：
	// 图形: [类型名]
	// 面积: [值]
	// 周长: [值]
	fmt.Printf("图形: %T\n", s)
	fmt.Printf("面积: %.2f\n", s.Area())
	fmt.Printf("周长: %.2f\n", s.Perimeter())
}

// 请实现 TotalArea 函数
func TotalArea(shapes ...Shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

// 请实现 TotalPerimeter 函数
func TotalPerimeter(shapes ...Shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.Perimeter()
	}
	return sum
}

// 请为各图形类型实现 fmt.Stringer 接口
// Rectangle 输出: "矩形(宽:10, 高:5)"
// Circle 输出: "圆形(半径:7)"
// Triangle 输出: "三角形(边A:3, 边B:4, 边C:5)"
func (r Rectangle) String() string {
	return fmt.Sprintf("矩形(宽%.2f, 高%.2f)", r.Width, r.Height)
}

func (c Circle) String() string {
	return fmt.Sprintf("圆形(半径%.2f)", c.Radius)
}

func (t Triangle) String() string {
	return fmt.Sprintf("三角形(边A%.2f, 边B%.2f, 边C%.2f)", t.A, t.B, t.C)
}

func main() {
	fmt.Println("=== 图形计算程序 ===")
	// 请在此编写测试代码

	r := Rectangle{Width: 3, Height: 5}
	c := Circle{Radius: 10}
	t := Triangle{A: 3, B: 4, C: 5}
	PrintShapeInfo(r)
	PrintShapeInfo(c)
	PrintShapeInfo(t)
	shapes := []Shape{r, c, t} // 使用 []T{...} 初始化语法
	fmt.Println("总面积：", TotalArea(shapes...))
	fmt.Println("总周长：", TotalPerimeter(shapes...))
	fmt.Println(r)
	fmt.Println(c)
	fmt.Println(t)
}
