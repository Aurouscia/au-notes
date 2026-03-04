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