# Go Workspace（工作区）

## 核心知识点

### 1. 什么是 Go Workspace

Go Workspace（工作区）是 Go 1.18 引入的功能，用于同时管理多个模块。当你需要同时修改多个相互依赖的模块时，工作区非常有用。

### 2. 关键文件

- **go.work**：定义工作区，包含多个模块的路径
- **go.work.sum**：工作区的校验和文件，记录依赖的加密哈希值（应添加到 .gitignore）

### 3. 常用命令

```bash
# 初始化工作区
go work init [模块路径...]

# 添加模块到工作区
go work use [模块路径]

# 递归添加模块（自动发现并添加子目录中的模块）
go work use -r [模块路径]

# 从工作区移除模块
go work edit -dropuse=[模块路径]

# 编辑 go.work 文件（类似 go mod edit）
go work edit [选项]

# 同步所有模块的依赖版本（将共同的依赖项升级到兼容的最高版本）
go work sync

# 查看当前工作区模式（显示正在使用的 go.work 文件路径）
go env GOWORK
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

### 5. 使用场景：同时开发多个相关模块

```
myproject/
├── go.work
├── frontend/      # 前端应用模块
├── backend/       # 后端服务模块
└── shared/        # 共享库模块
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

3. **工作区内模块无需发布到远程的原因**
   - 工作区将本地模块作为**主模块（main modules）**处理
   - Go 命令会优先解析工作区中的本地模块，而不是从远程下载
   - 模块间可以直接相互导入，无需先 `go get` 或发布到代码仓库
   - 本地修改可以立即被其他模块感知，无需提交或打 tag
   - 这使得多模块联合开发和调试变得高效

4. **replace 指令**
   - 用法（本地临时替换）
      ```go
      // 在 go.work 中使用 replace 指令，使用本地包代替远程包
      replace github.com/mycompany/library => ../library
      ```
   - `go.work` 中的 `replace` 指令**优先于** `go.mod` 中的 `replace`
   - 主要用于覆盖不同工作空间模块中冲突的替换

5. **文件位置建议**
   - `go.work` 应该放在项目根目录或工作区根目录
   - 确保 `.gitignore` 同时包含 `go.work` 和 `go.work.sum`

6. **使用 workspace 替代 replace 指令**
   - 多个模块在同一个仓库时，使用 workspace 比在每个 `go.mod` 中写 `replace` 更方便
   - 避免了多个模块间 `replace` 指令的冲突和维护负担