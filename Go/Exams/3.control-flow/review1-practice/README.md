# 复习实践题：文本分析器

## 需求

实现一个文本分析器，重点考察 **for-range**、**switch**、**defer** 和 **标签** 的综合运用。

## 功能要求

实现 `TextAnalyzer` 结构体和相关方法：

```go
type TextAnalyzer struct {
    Text string
}
```

### 1. 统计字符类型

```go
func (ta TextAnalyzer) CountTypes() map[string]int
```

遍历文本中的每个字符（使用 `for-range`），统计以下类型数量：
- `"chinese"`：中文字符（Unicode 范围 \u4e00-\u9fff）
- `"english"`：英文字母（a-z, A-Z）
- `"digit"`：数字（0-9）
- `"space"`：空白字符（空格、制表符、换行符）
- `"other"`：其他字符

使用 `switch` 实现分类逻辑。

### 2. 查找子串所有位置

```go
func (ta TextAnalyzer) FindAll(sub string) []int
```

返回子串在文本中出现的所有起始位置（字节索引）。

例如：`"abab"` 中查找 `"ab"`，返回 `[0, 2]`。

### 3. 逐行处理（使用 defer）

```go
func (ta TextAnalyzer) ProcessLines() []string
```

将文本按行分割，对每一行进行处理：
- 使用 `defer` 记录每行处理完成日志（`fmt.Println("处理完成:", line)`）
- 返回处理后的行列表（去除首尾空格）

### 4. 查找第一个满足条件的字符位置

```go
func (ta TextAnalyzer) FindFirst(predicate func(rune) bool) (int, rune)
```

遍历文本，返回第一个满足 `predicate` 的字符的**字节索引**和字符值。

如果没有找到，使用带标签的 `break` 或 `return` 处理。

## 主函数测试

在 `main()` 中完成以下测试：

1. 创建一个包含中英文混合、数字、空格的文本
2. 测试 `CountTypes()`，打印各类字符数量
3. 测试 `FindAll()`，查找某个子串的所有位置
4. 测试 `ProcessLines()`，观察 defer 的输出顺序
5. 测试 `FindFirst()`，查找第一个中文字符的位置

## 示例文本

```
Hello 世界！
Go语言 123
Test 测试
```

## 项目结构

```
review1-practice/
├── main.go
└── README.md
```

## 运行

```bash
go mod init text-analyzer
go run main.go
```
