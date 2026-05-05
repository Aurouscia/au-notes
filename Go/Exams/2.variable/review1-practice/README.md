# 复习实践题：购物车系统

## 需求

实现一个简单的购物车系统，重点考察**指针接收者**、**结构体嵌入**和**常量定义**的综合运用。

## 数据模型

### Product（商品）
```go
type Product struct {
    ID    uint
    Name  string
    Price float64
}
```

### CartItem（购物车项）
```go
type CartItem struct {
    Product  // 嵌入 Product
    Quantity int
}
```

### Cart（购物车）
```go
type Cart struct {
    Items []CartItem
}
```

## 功能要求

实现以下方法（注意接收者类型的选择）：

### 1. 添加商品到购物车
```go
func (c *Cart) AddItem(p Product, qty int)
```
- 如果购物车中已有相同商品，增加数量
- 如果没有，添加新项

### 2. 计算总价
```go
func (c Cart) TotalPrice() float64
```
- 返回购物车中所有商品的总价（Price * Quantity 之和）

### 3. 应用折扣
```go
func (c *Cart) ApplyDiscount(rate float64)
```
- 对所有商品的 Price 应用折扣（如 0.9 表示 9 折）
- **注意**：这里需要修改原购物车中的价格

### 4. 打印购物车
```go
func (c Cart) String() string
```
- 实现 `fmt.Stringer` 接口
- 返回格式示例：
  ```
  购物车（共2种商品）：
  - 苹果 x3, 单价: ￥5.00
  - 香蕉 x2, 单价: ￥3.00
  总价: ￥21.00
  ```

## 常量定义

定义以下常量：
- `DefaultDiscount = 1.0`（无折扣）
- `MemberDiscount = 0.9`（会员 9 折）
- `VipDiscount = 0.8`（VIP 8 折）

使用 `const` 定义，并思考：这些常量应该是有类型还是无类型？

## 主函数测试

在 `main()` 中完成以下测试：

1. 创建几个商品
2. 创建购物车，添加商品
3. 打印购物车（测试 `String()` 方法）
4. 应用会员折扣，再次打印
5. 验证总价计算是否正确

## 项目结构

```
review1-practice/
├── main.go
└── README.md
```

## 运行

```bash
go mod init shopping-cart
go run main.go
```
