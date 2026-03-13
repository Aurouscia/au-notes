# 函数高级特性 - 实践题

## 需求

请实现以下功能：

### 1. 变参函数 - 日志格式化

实现 `Logf(format string, args ...interface{})` 函数，功能如下：

- 在输出内容前添加 `[LOG]` 前缀
- 使用 `fmt.Sprintf` 格式化输出
- 如果没有参数，直接输出 format 字符串

示例：
```go
Logf("Hello %s", "World")  // 输出: [LOG] Hello World
Logf("Simple message")     // 输出: [LOG] Simple message
```

### 2. 闭包 - 累加器工厂

实现 `MakeAccumulator(initial float64) func(float64) float64` 函数：

- 返回一个闭包，该闭包接收一个 float64 参数
- 每次调用将参数累加到初始值上
- 返回累加后的结果

示例：
```go
acc := MakeAccumulator(10)
acc(5)   // 返回 15
acc(3)   // 返回 18
acc(-8)  // 返回 10
```

### 3. 闭包 - 缓存函数（Memoization）

实现 `MakeMemoizedFib() func(int) int` 函数：

- 返回一个带缓存的斐波那契数列计算函数
- 使用 map 缓存已计算的结果
- 避免重复计算

### 4. defer - 资源管理器

实现 `WithResource(name string, fn func())` 函数：

- 打印 `Opening resource: name`
- 使用 `defer` 确保在函数结束时打印 `Closing resource: name`
- 执行传入的 `fn` 函数

示例输出：
```
Opening resource: database
(执行 fn 中的代码)
Closing resource: database
```

### 5. 综合 - 函数管道

实现 `Pipeline(funcs ...func(int) int) func(int) int` 函数：

- 接收多个函数作为参数
- 返回一个新函数，该函数将输入依次通过所有函数处理
- 即 `Pipeline(f, g, h)(x)` 等价于 `h(g(f(x)))`

示例：
```go
add1 := func(x int) int { return x + 1 }
mul2 := func(x int) int { return x * 2 }
sub3 := func(x int) int { return x - 3 }

pipe := Pipeline(add1, mul2, sub3)
pipe(5)  // ((5 + 1) * 2) - 3 = 9
```

### 6. 主函数测试

在 `main()` 函数中测试以上所有功能，输出测试结果。

---

## 文件结构

```
exam4-advanced-practice/
├── go.mod
├── main.go    # 你的代码写在这里
└── README.md  # 本文件
```

---

## 运行要求

```bash
go run main.go
```
