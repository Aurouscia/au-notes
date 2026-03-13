package main

import "fmt"

// 请在此实现所有函数

func Logf(format string, args ...interface{}) {
	fmt.Print("[LOG] ")
	fmt.Println(fmt.Sprintf(format, args...))
}

func MakeAccumulator(initial float64) func(float64) float64 {
	sum := 0.0
	return func(x float64) float64 {
		sum += x
		return sum
	}
}

func MakeMemoizedFib() func(int) int {
	var m = make(map[int]int)
	return func(x int) int {
		if m[x] != 0 {

		}
		return 0
	}
}

func main() {
	fmt.Println("=== 函数高级特性实践题 ===")
	// 请在此编写测试代码
}
