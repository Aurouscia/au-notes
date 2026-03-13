# 函数基础 - 实践题

## 需求

请实现一个计算器程序，包含以下函数：

### 1. 基础运算函数

实现四个基础运算函数，每个函数接收两个 `float64` 参数，返回运算结果：

- `Add(a, b float64) float64` - 加法
- `Subtract(a, b float64) float64` - 减法
- `Multiply(a, b float64) float64` - 乘法
- `Divide(a, b float64) (float64, error)` - 除法（需要处理除数为0的情况）

### 2. 多返回值函数

实现 `CalculateAll(a, b float64)` 函数，同时返回加减乘除四个运算的结果。

返回值顺序为：加、减、乘、除。除法运算出错时，其他运算结果仍应正常返回。

### 3. 命名返回值函数

实现 `Rectangle(width, height float64)` 函数，使用命名返回值返回矩形的面积和周长。

### 4. 指针交换函数

实现 `Swap(a, b *int)` 函数，通过指针交换两个整数的值。

### 5. 主函数测试

在 `main()` 函数中：

1. 测试四个基础运算函数
2. 测试 `CalculateAll` 函数，并演示如何忽略不需要的返回值
3. 测试 `Rectangle` 函数
4. 测试 `Swap` 函数，验证交换是否成功

---

## 文件结构

```
exam2-practice/
├── go.mod
├── main.go    # 你的代码写在这里
└── README.md  # 本文件
```

---

## 运行要求

```bash
go run main.go
```

程序应输出各函数的测试结果。
