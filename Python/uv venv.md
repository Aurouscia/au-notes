# 创建 .venv

- uv venv --python 3.12（创建指定解释器版本的虚拟环境）
- uv sync（自动创建 .venv（如果不存在）+ 根据 uv.lock 安装）

# 激活 .venv

在新时代，无需关心虚拟环境激活了没有

uv 会根据当前所在的目录或其父目录，自动检测并使用 `.venv` 环境内或全局的 python

# 添加依赖项

## 最佳做法

- uv add xxx（添加依赖项并记录到 uv.lock 文件）
- uv remove xxx（移除依赖项并更新 uv.lock 文件）
- uv lock（根据 pyproject.toml 重新生成 uv.lock 文件）
- uv run python main.py（自动使用 .venv，无需激活）

## 不推荐的做法

- uv pip install（当在虚拟环境目录下时，会自动检测到，并安装依赖项到该虚拟环境）
    - 新项目不建议使用，应该改为使用 uv add

# 类比理解

- uv.lock 相当于 package-lock.json 或 pnpm-lock.yaml，用于记录项目的依赖项版本
- .venv 目录相当于一个隔离的 Python 运行环境（类似包含 node 二进制 + `node_modules`），用来存储 Python 解释器和已安装的依赖项（uv 会把依赖项的文件通过硬链接或符号链接等方式连接到缓存目录，避免多个项目中同样的文件浪费存储空间）