# 考试 3：Go Workspace 实践

## 需求

你需要创建一个 Go Workspace，同时管理两个模块：一个工具库和一个使用它的应用程序。

### 项目结构要求

```
exam3-workspace/
├── go.work
├── mylib/              # 工具库模块
│   ├── go.mod
│   └── utils.go        # 提供 Greet 函数
└── myapp/              # 应用程序模块
    ├── go.mod
    └── main.go         # 调用 mylib 的 Greet 函数
```

### 具体要求

#### 1. 创建工具库模块 `mylib`

- 模块路径：`example.com/mylib`
- 创建 `utils.go` 文件，实现 `Greet(name string) string` 函数
- 函数返回格式：`"Hello, [name]!"`

#### 2. 创建应用程序模块 `myapp`

- 模块路径：`example.com/myapp`
- 创建 `main.go` 文件，导入并使用 `example.com/mylib`
- 在 `main()` 中调用 `mylib.Greet("Workspace")` 并打印结果

#### 3. 初始化并配置 Workspace

- 在 `exam3-workspace/` 目录下初始化 workspace
- 将 `mylib` 和 `myapp` 都添加到 workspace
- 确保 `myapp` 可以直接使用本地 `mylib` 模块（无需发布到远程）

#### 4. 运行验证

- 在 `myapp` 目录下运行 `go run .`，应输出：`Hello, Workspace!`

## 提示

1. 创建模块：`go mod init 模块路径`
2. 初始化 workspace：`go work init [模块路径...]` 或先 init 再 use
3. 添加模块到 workspace：`go work use [模块路径]`
4. Workspace 内的模块可以直接相互导入
