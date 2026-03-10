package main

import "fmt"

// ReverseArray 反转数组（不修改原数组）
func ReverseArray(arr [5]int) [5]int {
	// 请在此实现
	return [5]int{}
}

// Unique 切片去重（保持顺序）
func Unique(s []int) []int {
	// 请在此实现
	return nil
}

// Transpose 矩阵转置
func Transpose(matrix [][]int) [][]int {
	// 请在此实现
	return nil
}

func main() {
	// 验证任务1：数组反转
	arr := [5]int{1, 2, 3, 4, 5}
	reversed := ReverseArray(arr)
	fmt.Println("原数组:", arr)
	fmt.Println("反转后:", reversed)

	// 验证任务2：切片去重
	s := []int{1, 2, 2, 3, 3, 3, 4, 2, 1}
	unique := Unique(s)
	fmt.Println("原切片:", s)
	fmt.Println("去重后:", unique)

	// 验证任务3：矩阵转置
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	transposed := Transpose(matrix)
	fmt.Println("原矩阵:")
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println("转置后:")
	for _, row := range transposed {
		fmt.Println(row)
	}
}
