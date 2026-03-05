# Go 模块基础

## 核心知识点

### 1. 什么是 Go Module（模块）

Go Module 是 Go 1.11 引入的依赖管理系统，用于管理项目的依赖包和版本。

### 2. 关键文件

- **go.mod**：定义模块名称、Go 版本和依赖项
- **go.sum**：记录依赖包的加密校验和，确保依赖完整性

### 3. 常用命令

```bash
# 初始化新模块
go mod init <模块名>

# 下载依赖（部署时用）
go mod download
# 不会更改 go.mod 或 go.sum

# 整理并下载缺失的依赖（开发时用）
go mod tidy
# 做三件事：
# 1. 扫描代码，添加缺失的依赖到 go.mod
# 2. 删除未使用的依赖
# 3. 下载依赖并更新 go.sum

# 查看依赖关系
go mod graph
```

### 4. go.mod 文件结构

```
module example.com/myproject

go 1.21

require (
    github.com/some/package v1.2.3
)
```

- go 1.26 新行为：使用 go mod init 时，go.mod 中的版本是当前版本-1
- go 行会影响项目中可用的特性（go 1.25 的项目无法使用 1.26 的新特性）

### 5. 模块路径规范

- 本地开发：可以使用简单名称如 `myproject`
- 开源项目：建议使用代码托管地址，如 `github.com/username/repo`

---

### 6. 镜像

go 的模块一般在github上，如果网络不好可以设置镜像：
```bash
# 临时使用
go env -w GOPROXY=https://goproxy.cn,direct

# 或设置环境变量
export GOPROXY=https://goproxy.cn,direct
```

### 7. 自动工具链切换

go 运行时会根据本地设置（默认为auto）以及项目要求（toolchain行和依赖项），自动下载并运行对应版本的工具链

本地设置为：
- `GOTOOLCHAIN`环境变量
- 用户级：`~/.config/go/env`
- 系统级：`$GOROOT/go.env`

可选值：
- local → 使用捆绑工具链（不切换）
- name → 使用指定工具链
- name+auto → 进入智能选择流程
    - 读取 go.work 或 go.mod 的 toolchain 行
    - 如果比name更高，则自动使用该版本

自动下载的工具链存储在`~/go/pkg/mod/golang.org/toolchain@*`（会越堆越多，不会自动清理，可手动安全删除）