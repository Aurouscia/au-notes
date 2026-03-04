# Go Workspace（工作区）

## 核心知识点

### 1. 什么是 Go Workspace

Go Workspace（工作区）是 Go 1.18 引入的功能，用于同时管理多个模块。当你需要同时修改多个相互依赖的模块时，工作区非常有用。

### 2. 关键文件

- **go.work**：定义工作区，包含多个模块的路径

### 3. 常用命令

```bash
# 初始化工作区
go work init [模块路径...]

# 添加模块到工作区
go work use [模块路径]

# 从工作区移除模块
go work edit -dropuse=[模块路径]

# 查看工作区状态
go work sync
```

### 4. go.work 文件结构

```go
go 1.21

use (
    ./myapp
    ./mylibrary
    ./utils
)

replace example.com/old => ./local/old
```

### 5. 使用场景

**场景 1：同时开发多个相关模块**

```
myproject/
├── go.work
├── frontend/      # 前端应用模块
├── backend/       # 后端服务模块
└── shared/        # 共享库模块
```

**场景 2：本地替换依赖进行调试**

```go
// 在 go.work 中使用 replace 指令
replace github.com/mycompany/library => ../library
```

### 6. 工作区 vs 模块

| 特性 | Module（模块） | Workspace（工作区） |
|------|---------------|-------------------|
| 管理范围 | 单个项目 | 多个相关模块 |
| 配置文件 | go.mod | go.work |
| 依赖解析 | 从远程下载 | 优先使用本地模块 |
| 使用场景 | 独立项目 | 多模块联合开发 |

### 7. 最佳实践

1. **不要将 go.work 提交到版本控制**
   - 在 `.gitignore` 中添加 `go.work` 和 `go.work.sum`
   - 工作区配置是个人开发环境相关的

2. **CI/CD 环境不要使用工作区**
   - 工作区只用于本地开发
   - 生产构建应该基于独立的 go.mod

3. **模块路径保持一致**
   - 工作区中的模块路径应与 go.mod 中声明的一致

---

**请阅读并理解上述内容，准备好后告诉我，我会为你创建相应的练习题。**
