# 考试 2：变量与常量实践题

## 需求

创建一个 Go 程序，实现一个常量配置包和一个使用配置的 main 程序。

### 项目结构

```
exam2-practice/
├── go.mod
├── config/
│   └── config.go    # 常量定义
└── main.go          # 主程序
```

### 具体要求

#### 1. 创建 `config/config.go`

定义以下常量：

- `AppName`：应用程序名称，值为 `"MyApp"`
- `Version`：版本号，值为 `"1.0.0"`
- `MaxConnections`：最大连接数，值为 `100`

使用 `iota` 定义以下枚举常量（表示日志级别）：

- `LevelDebug` = 0
- `LevelInfo` = 1
- `LevelWarning` = 2
- `LevelError` = 3

#### 2. 创建 `main.go`

实现功能：

1. 使用短变量声明 `:=` 创建一个变量 `currentLogLevel`，初始值为 `config.LevelInfo`

2. 使用 `var` 声明一个变量 `connectionCount`，不初始化（使用零值）

3. 打印以下信息：
   - 应用名称和版本
   - 当前日志级别（数字）
   - 当前连接数（应该显示零值）
   - 最大连接数限制

4. 使用多重赋值交换两个变量的值：
   ```go
   a, b := 10, 20
   // 交换后，a = 20, b = 10
   ```
   打印交换前后的值

### 预期输出

```
应用: MyApp v1.0.0
日志级别: 1
当前连接: 0 / 100
交换前: a=10, b=20
交换后: a=20, b=10
```

## 提示

1. 初始化模块：`go mod init exam2-practice`
2. 导入 config 包：`import "exam2-practice/config"`
3. 使用 `iota` 时注意它是从 0 开始递增的
4. 交换变量使用多重赋值：`a, b = b, a`
