package main

import "fmt"

// ==================== 任务 1：图书 ====================

// Book 结构体定义
// 请在此定义 Book 结构体

type Book struct {
	Title  string
	Author string
	Price  float64
}

// String 方法：返回图书信息字符串
// func (b Book) String() string

func (b Book) String() string {
	return fmt.Sprintf("《%s》 by %s - ￥%.2f", b.Title, b.Author, b.Price)
}

// Discount 方法：应用折扣
// func (b *Book) Discount(rate float64)

func (b *Book) Discount(rate float64) {
	b.Price *= rate
}

// ==================== 任务 2：矩形和圆形 ====================

// Rectangle 结构体定义
// 请在此定义 Rectangle 结构体

type Rectangle struct {
	Width  float64
	Height float64
}

// Area 方法：计算矩形面积
// func (r Rectangle) Area() float64

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle 结构体定义
// 请在此定义 Circle 结构体

type Circle struct {
	Radius float64
}

// Area 方法：计算圆形面积
// func (c Circle) Area() float64

func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// ==================== 任务 3：嵌入结构体 ====================

// Address 结构体定义
// 请在此定义 Address 结构体

type Address struct {
	City   string
	Street string
}

// Company 结构体定义（嵌入 Address）
// 请在此定义 Company 结构体

type Company struct {
	Address
	Name string
}

// Info 方法：返回公司信息
// func (c Company) Info() string

func (c Company) Info() string {
	return fmt.Sprintf("公司名称: %s, 地址: %s市%s街道", c.Name, c.City, c.Street)
}

// ==================== main 函数 ====================

func main() {
	// 任务 1 验证
	fmt.Println("=== 任务 1：图书 ===")
	book := Book{Title: "Harry Potter", Author: "J. K. Rowling", Price: 100}
	fmt.Println(book.String())
	book.Discount(0.8)
	fmt.Println(book.String())

	// 任务 2 验证
	fmt.Println("\n=== 任务 2：面积计算 ===")
	rect := Rectangle{Width: 3, Height: 4}
	circle := Circle{Radius: 5}
	fmt.Printf("矩形面积: %.2f\n", rect.Area())
	fmt.Printf("圆形面积: %.2f\n", circle.Area())

	// 任务 3 验证
	fmt.Println("\n=== 任务 3：公司信息 ===")
	company := Company{Name: "字节跳动", Address: Address{City: "上海", Street: "江湾城路"}}
	fmt.Println(company.Info())
}
