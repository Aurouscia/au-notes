package main

import "fmt"

// 请在此实现所有函数

func Logf(format string, args ...interface{}) {
	fmt.Print("[LOG] ")
	fmt.Println(fmt.Sprintf(format, args...))
}

func MakeAccumulator(initial float64) func(float64) float64 {
	sum := initial
	return func(x float64) float64 {
		sum += x
		return sum
	}
}

func MakeMemoizedFib() func(int) int {
	var m = make(map[int]int)
	var fib func(int) int
	fib = func(x int) int {
		var cached int = m[x]
		if cached != 0 {
			return cached
		}
		// fib 数列里没有 0，0 可认为 cache miss
		// ❌有误：实际上 fib 数列第一项是 0，如果需要判断 cache miss，可使用 map[int]*int，使用 nil 表示“无值”
		var res int
		if x <= 1 {
			res = 1
		} else {
			res = fib(x-1) + fib(x-2)
		}
		m[x] = res
		return res
	}
	return fib
}

func WithResource(name string, fn func()) {
	fmt.Println("Opening resource:", name)
	defer fmt.Println("Closing resource:", name)
	fn()
}

func Pipeline(funcs ...func(int) int) func(int) int {
	return func(x int) int {
		for _, f := range funcs {
			x = f(x)
		}
		return x
	}
}

func main() {
	fmt.Println("=== 函数高级特性实践题 ===")
	// 请在此编写测试代码

	Logf("Hello %s", "World") // 输出: [LOG] Hello World
	Logf("Simple message")    // 输出: [LOG] Simple message

	acc := MakeAccumulator(10)
	fmt.Println(acc(5))  // 返回 15
	fmt.Println(acc(3))  // 返回 18
	fmt.Println(acc(-8)) // 返回 10

	fib := MakeMemoizedFib()
	fmt.Println("fib(3) =", fib(3))
	fmt.Println("fib(5) =", fib(5))

	WithResource("database", func() {
		fmt.Println("doing something with that database")
	})

	add1 := func(x int) int { return x + 1 }
	mul2 := func(x int) int { return x * 2 }
	sub3 := func(x int) int { return x - 3 }
	pipe := Pipeline(add1, mul2, sub3)
	fmt.Println(pipe(5)) // ((5 + 1) * 2) - 3 = 9
}
