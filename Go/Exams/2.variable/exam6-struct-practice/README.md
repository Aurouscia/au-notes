# 结构体实践题

## 题目要求

请完成以下 3 个任务，所有代码写在 `main.go` 中。

---

### 任务 1：定义图书结构体

定义一个 `Book` 结构体，包含以下字段：
- `Title`（书名，string）
- `Author`（作者，string）
- `Price`（价格，float64）

实现方法：
1. `String()` 方法：返回格式为 `"《书名》 by 作者 - ￥价格"` 的字符串
2. `Discount(rate float64)` 方法：指针接收者，按给定折扣率（如 0.8 表示 8 折）修改价格

---

### 任务 2：定义矩形和圆形

定义两个结构体：

```go
type Rectangle struct {
    Width  float64
    Height float64
}

type Circle struct {
    Radius float64
}
```

为它们实现 `Area() float64` 方法，计算面积：
- 矩形面积 = 宽 × 高
- 圆形面积 = π × 半径²（π 取 3.14159）

---

### 任务 3：使用嵌入结构体

定义 `Address` 结构体：

```go
type Address struct {
    City   string
    Street string
}
```

定义 `Company` 结构体，嵌入 `Address`：

```go
type Company struct {
    Name string
    Address  // 嵌入结构体
}
```

实现 `Info() string` 方法，返回格式为 `"公司名称: XXX, 地址: XX市XX街道"`

---

## 验证

在 `main` 函数中：
1. 创建一本图书，打印信息，应用 8 折折扣后再打印
2. 创建一个矩形（宽 3，高 4）和一个圆形（半径 5），打印它们的面积
3. 创建一个公司，打印公司信息
