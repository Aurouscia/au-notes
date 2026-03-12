package main

import "fmt"

// ==================== 任务 1：图书 ====================

// Book 结构体定义
// 请在此定义 Book 结构体

// String 方法：返回图书信息字符串
// func (b Book) String() string

// Discount 方法：应用折扣
// func (b *Book) Discount(rate float64)

// ==================== 任务 2：矩形和圆形 ====================

// Rectangle 结构体定义
// 请在此定义 Rectangle 结构体

// Area 方法：计算矩形面积
// func (r Rectangle) Area() float64

// Circle 结构体定义
// 请在此定义 Circle 结构体

// Area 方法：计算圆形面积
// func (c Circle) Area() float64

// ==================== 任务 3：嵌入结构体 ====================

// Address 结构体定义
// 请在此定义 Address 结构体

// Company 结构体定义（嵌入 Address）
// 请在此定义 Company 结构体

// Info 方法：返回公司信息
// func (c Company) Info() string

// ==================== main 函数 ====================

func main() {
	// 任务 1 验证
	fmt.Println("=== 任务 1：图书 ===")
	// book := Book{...}
	// fmt.Println(book.String())
	// book.Discount(0.8)
	// fmt.Println(book.String())

	// 任务 2 验证
	fmt.Println("\n=== 任务 2：面积计算 ===")
	// rect := Rectangle{...}
	// circle := Circle{...}
	// fmt.Printf("矩形面积: %.2f\n", rect.Area())
	// fmt.Printf("圆形面积: %.2f\n", circle.Area())

	// 任务 3 验证
	fmt.Println("\n=== 任务 3：公司信息 ===")
	// company := Company{...}
	// fmt.Println(company.Info())
}
