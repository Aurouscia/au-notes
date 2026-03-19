# 错题集：流程控制

## 一、for-range 遍历字符串的返回值

**错误题目**：`for-range` 遍历字符串时，返回的是字节索引和 Unicode 码点。

**错误答案**：错：索引应该是 unicode 字符的索引

**正确答案**：返回的是**字节索引**（UTF-8 编码的字节位置）和 **rune**（Unicode 码点）

**知识点**：
- `for-range` 遍历字符串时，第一个返回值是**字节索引**（不是字符索引）
- 第二个返回值是 `rune` 类型（Unicode 码点）
- 由于 UTF-8 变长编码，中文字符占 3 个字节，所以索引会跳跃

**示例**：
```go
s := "Hello, 世界"
for i, r := range s {
    fmt.Printf("索引 %d: %c\n", i, r)
}
// 输出：
// 索引 0: H
// 索引 1: e
// ...
// 索引 7: 世  (跳过了 8, 9)
// 索引 10: 界
```

---

## 二、switch case 可以是表达式

**错误题目**：`switch` 的 case 只能是常量值。

**错误答案**：对

**正确答案**：错，Go 的 switch case **可以是表达式**，不限于常量

**知识点**：
- Go 的 `switch` 比 C/Java 更灵活
- case 可以是任意表达式，只要类型匹配即可
- 甚至可以省略 switch 后的表达式，在 case 中写条件

**示例**：
```go
n := 10
switch {
case n > 0 && n < 5:
    fmt.Println("小")
case n >= 5 && n < 10:
    fmt.Println("中")
default:
    fmt.Println("大")
}
```

---

## 三、continue 与循环条件

**错误题目**：以下代码会输出 0, 1, 2, 3？
```go
for i := 0; i < 3; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}
```

**错误答案**：错，应该是 0，1，3

**正确答案**：应该是 **0, 1**（i=2 时 continue 跳过，i=3 时循环条件 i<3 不满足，循环结束）

**分析**：
- i=0：输出 0
- i=1：输出 1
- i=2：continue 跳过，不输出
- i=3：i < 3 为 false，循环结束，不会输出 3

---

## 四、fallthrough 的行为

**错误题目**：
```go
switch n {
case 1, 3, 5:
    fmt.Print("奇数 ")
    fallthrough
case 2, 4, 6:
    fmt.Print("偶数 ")
default:
    fmt.Print("其他")
}
```

**错误答案**：奇数 其他

**正确答案**：**奇数 偶数**

**原因**：`fallthrough` 会**强制执行下一个 case**，不会跳到 default

**知识点**：
- `fallthrough` 会无条件执行下一个 case 的代码
- 它不会判断下一个 case 的条件是否满足
- 使用 `fallthrough` 后，程序会继续顺序执行，直到遇到 break 或 switch 结束

---

## 五、for-range 语法错误

**错误题目**：简述 `for-range` 遍历切片时，如果只想要值而不需要索引，应该如何写？

**错误答案**：`for _, value range xxx`，使用 `_` 丢弃值

**正确答案**：`for _, value := range xxx`

**错误原因**：遗漏了 `:=` 赋值符号

---

## 总结

| 错误类型 | 关键知识点 |
|---------|-----------|
| for-range 字符串 | 返回字节索引（不是字符索引）和 rune |
| switch case | 可以是表达式，不限于常量 |
| continue | 跳过当前迭代，但不会跳过循环条件判断 |
| fallthrough | 强制执行下一个 case，不是跳到 default |
| for-range 语法 | 必须使用 `:=` 或 `=` |
