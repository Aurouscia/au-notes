# 练习题 2：实践 - 创建并管理 Go 模块

## 任务

在这个练习中，你需要完成以下步骤：

### 步骤 1：初始化模块

1. 进入 ` Exams/1.module/exam2-practice/` 目录
2. 初始化一个名为 `mycalculator` 的 Go 模块

### 步骤 2：编写代码

创建 `main.go` 文件，实现一个简单的计算器，包含以下函数：

```go
// Add 返回两数之和
func Add(a, b int) int

// Subtract 返回两数之差
func Subtract(a, b int) int
```

在 `main()` 函数中调用这两个函数并打印结果。

### 步骤 3：添加依赖

1. 使用 `github.com/fatih/color` 包来彩色输出计算结果
2. 安装该依赖（提示：先写代码 import，再运行某个命令）

### 步骤 4：验证

1. 检查生成的 `go.mod` 文件内容
2. 确认 `go.sum` 文件已生成
3. 运行程序，确保能正常输出彩色结果

---

## 预期结果

- `go.mod` 文件包含模块名 `mycalculator` 和 `github.com/fatih/color` 依赖
- `go.sum` 文件存在且非空
- 运行 `go run main.go` 能输出彩色文字

完成后请告诉我"做完了，请检查"。
