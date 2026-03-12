package main

import "fmt"

// Swap 交换两个指针指向的值
func Swap(a, b *int) {
	var temp = *a
	*a = *b
	*b = temp
}

// SumArray 计算数组元素之和（通过指针）
// ❌有误：原实现 for x := range *arr 获取的是索引而非值
// 已修正为 for _, x := range *arr
func SumArray(arr *[5]int) int {
	sum := 0
	for _, x := range *arr {
		sum += x
	}
	return sum
}

// Person 结构体
type Person struct {
	Name string
	Age  int
}

// NewPerson 使用 new 创建并初始化 Person
func NewPerson(name string, age int) *Person {
	var p = new(Person)
	(*p).Name = name
	(*p).Age = age
	return p
}

func main() {
	// 验证任务1：交换两个整数
	x, y := 10, 20
	fmt.Printf("交换前: x=%d, y=%d\n", x, y)
	Swap(&x, &y)
	fmt.Printf("交换后: x=%d, y=%d\n", x, y)

	// 验证任务2：数组求和
	arr := [5]int{1, 2, 3, 4, 5}
	sum := SumArray(&arr)
	fmt.Printf("数组: %v, 和: %d\n", arr, sum)

	// 验证任务3：创建 Person
	p := NewPerson("Alice", 30)
	fmt.Printf("Person: Name=%s, Age=%d\n", p.Name, p.Age)
}
