# fmt

## 常用格式化动词

| 动词    | 说明       | 示例                             |
| ----- | -------- | ------------------------------ |
| `%v`  | 默认格式     | `fmt.Printf("%v", user)`       |
| `%+v` | 结构体时带字段名 | `fmt.Printf("%+v", user)`      |
| `%#v` | Go 语法表示  | `fmt.Printf("%#v", user)`      |
| `%T`  | 类型       | `fmt.Printf("%T", 123)` // int |
| `%d`  | 十进制整数    | `fmt.Printf("%d", 42)`         |
| `%s`  | 字符串      | `fmt.Printf("%s", "hello")`    |
| `%f`  | 浮点数      | `fmt.Printf("%.2f", 3.14159)`  |
| `%t`  | 布尔值      | `fmt.Printf("%t", true)`       |
| `%p`  | 指针       | `fmt.Printf("%p", &x)`         |

## 标准输出函数 os.Stdout

| 函数                                        | 功能             | 示例                                |
| ----------------------------------------- | -------------- | --------------------------------- |
| `Print(a ...interface{})`                 | 无格式输出，无换行      | `fmt.Print("a", "b")` → `ab`      |
| `Println(a ...interface{})`               | 无格式输出，自动加空格和换行 | `fmt.Println("a", "b")` → `a b\n` |
| `Printf(format string, a ...interface{})` | 格式化输出，无换行      | `fmt.Printf("name: %s", "tom")`   |

## 标准输入函数 os.Stdin

| 函数                                       | 功能             | 示例                                |
| ---------------------------------------- | -------------- | --------------------------------- |
| `Scan(a ...interface{})`                 | 从标准输入扫描值，以空格分隔 | `fmt.Scan(&name, &age)`           |
| `Scanln(a ...interface{})`               | 扫描值，遇到换行停止     | `fmt.Scanln(&name)`               |
| `Scanf(format string, a ...interface{})` | 按格式扫描          | `fmt.Scanf("%s %d", &name, &age)` |

## 字符串格式化函数（返回字符串）

| 函数                                                | 功能            | 示例                                       |
| ------------------------------------------------- | ------------- | ---------------------------------------- |
| `Sprint(a ...interface{}) string`                 | 无格式拼接为字符串     | `s := fmt.Sprint("a", "b")` // "ab"      |
| `Sprintln(a ...interface{}) string`               | 拼接为字符串，加空格和换行 | `s := fmt.Sprintln("a", "b")` // "a b\n" |
| `Sprintf(format string, a ...interface{}) string` | 格式化返回字符串      | `s := fmt.Sprintf("id: %d", 100)`        |

## 字符串读取函数（输入字符串）

| 函数                                                    | 功能        | 示例                                   |
| ----------------------------------------------------- | --------- | ------------------------------------ |
| `Sscan(str string, a ...interface{})`                 | 从字符串扫描    | `fmt.Sscan("123 abc", &num, &str)`   |
| `Sscanln(str string, a ...interface{})`               | 从字符串扫描到换行 | `fmt.Sscanln("123", &num)`           |
| `Sscanf(str string, format string, a ...interface{})` | 按格式从字符串扫描 | `fmt.Sscanf("id:123", "id:%d", &id)` |


## 错误格式化函数（返回错误）

| 函数                                              | 功能      | 示例                                        |
| ----------------------------------------------- | ------- | ----------------------------------------- |
| `Errorf(format string, a ...interface{}) error` | 格式化创建错误 | `return fmt.Errorf("invalid id: %d", id)` |

## 从 Reader 和 Writer 读取和写入

| 函数                                                      | 功能             | 示例                                      |
| ------------------------------------------------------- | -------------- | --------------------------------------- |
| `Fprint(w io.Writer, a ...interface{})`                 | 无格式写入 Writer   | `fmt.Fprint(os.Stderr, "error")`        |
| `Fprintln(w io.Writer, a ...interface{})`               | 无格式写入，加空格换行    | `fmt.Fprintln(file, "line")`            |
| `Fprintf(w io.Writer, format string, a ...interface{})` | 格式化写入 Writer   | `fmt.Fprintf(w, "code: %d", 200)`       |
| `Fscan(r io.Reader, a ...interface{})`                  | 从 Reader 扫描    | `fmt.Fscan(file, &x)`                   |
| `Fscanln(r io.Reader, a ...interface{})`                | 从 Reader 扫描到换行 | `fmt.Fscanln(file, &x)`                 |
| `Fscanf(r io.Reader, format string, a ...interface{})`  | 按格式从 Reader 扫描 | `fmt.Fscanf(file, "%d,%s", &id, &name)` |

## 记忆方法：

- 基底：print 和 scan
- 前缀：S 表示字符串操作，F 表示流操作
- 后缀：无后缀表示空白字符分隔，ln 表示行分隔，f 表示提供格式字符串
- 特例：error 只有 Errorf 一个函数

## & 符号

scan 系列的函数需要修改值，所以必须传入值的指针，而不是值本身（使用 & 取地址）