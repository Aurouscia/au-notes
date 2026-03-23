# 试题进度总览

## 1.module - Go 模块基础

- exam1-concepts.md：概念题（填空、判断、简答）✅ 优秀
- exam2-practice/：实践题（创建模块、添加依赖）✅ 优秀
- exam3-workspace/：实践题（Go Workspace 多模块管理）✅ 优秀
- 状态：已完成

## 2.variable - 变量与常量

- exam1-basic-concepts.md：概念题 ✅ 需要加强（无类型常量、短变量声明复用理解有误）
- exam2-basic-practice/：实践题 ✅ 基本正确（注释有误）
- exam3-array-slice-concepts.md：数组与切片概念题 ✅ 已完成
- exam4-array-slice-practice/：数组与切片实践题 ✅ 已完成
- exam5-struct-concepts.md：结构体概念题 ✅ 需要加强（值接收者方法无法修改原结构体）
- exam6-struct-practice/：结构体实践题 ✅ 已完成
- exam7-advanced-concepts/：进阶概念题 ✅ 已完成（多行 const 的格式错误）
- 状态：已完成

## 3.control-flow - 流程控制

- exam1-concepts.md：概念题 ✅ 需要加强（switch case 规则、fallthrough 行为、for-range 语法细节）
- exam2-practice/：实践题（猜数字游戏）✅ 优秀
- exam3-advanced-concepts.md：进阶概念题 ✅ 需要加强（for-range 字符串字节索引、for 内部的 defer 执行顺序）
- 状态：已完成

## 4.pointer - 指针

- exam1-concepts.md：概念题 ✅ 已完成
- exam2-practice/：实践题 ✅ 已完成（修正了 for-range 用法）
- exam3-fix-bug/：纠错题 ✅ 已完成
- exam4-advanced-concepts.md：进阶概念题 ✅ 需要加强（`new()`内存分配机制、逃逸分析、`&`运算符适用范围、`unsafe.Pointer`概念）
- 状态：已完成

## 5.function - 函数

- exam1-concepts.md：概念题（填空、判断、简答）✅ 需要加强（裸返回概念理解有误）
- exam2-practice/：实践题（计算器函数实现）✅ 优秀
- exam3-advanced-concepts.md：高级特性概念题 ✅ 优秀
- exam4-advanced-practice/：高级特性实践题 ✅ 优秀
- 状态：已完成

## 6.interface - 接口

- exam1-concepts.md：概念题（填空、判断、简答）✅ 需要加强（类型断言零值、接口组合、值类型不能自动取地址）
- exam2-practice/：实践题（图形计算程序）✅ 需要加强（fmt.Stringer 方法名）
- exam3-pointer-practice/：进阶实践题（指针接收者、薄弱环节强化）✅ 需要加强（函数类型不变性、自动取地址与接口赋值区别、类型断言接口优于具体类型）
- 状态：已完成

## 7.goroutine - 协程与通道

- exam1-concepts.md：概念题（Channel 基础）✅ 已完成
- exam2-practice/：实践题（并发素数筛）✅ 已完成
- exam3-select-concepts.md：概念题（Select、WaitGroup、Context）✅ 需要加强（case 语法格式、WithValue 不返回 cancel、WaitGroup 可重复使用）
- exam4-select-practice/：实践题（多路复用下载器）✅ 需要加强（同步调用导致的超时失效、channel 关闭位置错误，Done 位置错误，goroutine 操作外部切片导致线程不安全问题）
- exam4-context-practice/：实践题（可取消的任务队列）⏳ 待完成
- exam5-fix-bug/：纠错题（找出并修复代码问题）⏳ 待完成
- 状态：进行中
