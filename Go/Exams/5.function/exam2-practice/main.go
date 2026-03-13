package main

import (
	"errors"
	"fmt"
)

// 请在此实现所有函数

func Add(a, b float64) float64 {
	return a + b
}
func Subtract(a, b float64) float64 {
	return a - b
}
func Multiply(a, b float64) float64 {
	return a * b
}
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func CalculateAll(a, b float64) (addRes, subRes, mulRes, divRes float64, divErr error) {
	addRes = Add(a, b)
	subRes = Subtract(a, b)
	mulRes = Multiply(a, b)
	divRes, divErr = Divide(a, b)
	return
}

func Rectangle(width, height float64) (area, circumference float64) {
	area = width * height
	circumference = 2*width + 2*height
	return
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println("=== 函数基础实践题 ===")
	// 请在此编写测试代码

	fmt.Println("3+2=", Add(3, 2))
	fmt.Println("3-2=", Subtract(3, 2))
	fmt.Println("3*2=", Multiply(3, 2))

	var divRes, err = Divide(3, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("3/2=", divRes)
	}

	divRes, err = Divide(3, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("3/0=", divRes)
	}

	var r1, r2, r3, _, _ = CalculateAll(3, 2)
	fmt.Println("CalculateAll(3, 2) 前三个返回值：", r1, r2, r3)

	var area, cir = Rectangle(3, 4)
	fmt.Println("3*4 矩形的面积和周长：", area, cir)

	var pa = new(int)
	*pa = 3
	var pb = new(int)
	*pb = 4
	fmt.Println("交换前：", *pa, *pb)
	Swap(pa, pb)
	fmt.Println("交换后：", *pa, *pb)
}
