# Go 模块基础

## 核心知识点

### 1. 什么是 Go Module（模块）

Go Module 是 Go 1.11 引入的依赖管理系统，用于管理项目的依赖模块和版本。

### 2. 关键文件

- **go.mod**：定义模块名称、Go 版本和依赖项
- **go.sum**：记录依赖模块的加密校验和，确保依赖完整性

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

注意：
- 如果是 workspace 项目中的依赖，且还未发布模块到远程，则 go mod tidy 无法成功
- 因为 tidy 会确保这个模块在别处也能成功引用依赖，会试图从远程下载而不是使用 workspace 内的
- 原因：go mod xxx 系列命令都只针对当前模块，感知不到 workspace 配置

### 4. go.mod 文件结构

```
module example.com/myproject

go 1.21

require (
    github.com/some/package v1.2.3
)
```

- module 行是 go.mod 的第一行，后跟模块名，模块名为 git 仓库导入路径
- go.mod 中的 require 指定依赖模块的版本（锁定），而源码中的 import 仅指定依赖模块的名称（声明）
- go 1.26 新行为：使用 go mod init 时，go.mod 中的版本是当前版本-1
- go 行会影响项目中可用的特性（例如 go 1.25 的项目无法使用 1.26 的新特性，否则 gopls 等开发环境会警告）
- 在 workspace 中，无需 require（可能还未发布，url无效）也可以 import 工作区内其他模块

### 5. 模块路径规范

- 本地开发：可以使用简单名称如 `myproject`
- 开源项目：必须使用代码托管地址，如 `github.com/username/repo`

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

### 8. Package

- 每个 go 文件的第一行为 package + 包名
- 包名必须与 go 文件所属的目录同名，例如`package utils`，否则会导致各种工具识别异常

例如：在模块 `example.com/myadder` 中有目录 `utils`，其中的文件均为 `package utils`

导入时：
```go
import (
	"fmt"
	"example.com/myadder/utils" // 模块名+模块内路径（按约定，路径结尾应与包名一致）
)

utils.Add(1, 2) // 通过包名，可调用里面的公共函数
```

如果两个模块的包同名了，必须使用 alias：
```go
import (
	"fmt"
	autils "example.com/myadder/utils"
    butils "example.com/mysubber/utils"
)
```
然后通过 autils 和 butils 调用里面的公共函数