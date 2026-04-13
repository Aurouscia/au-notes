# python

- uv python install 3.12
- uv python list
- uv venv --python 3.12（创建某版本的虚拟环境）
    - uv pip install（当在虚拟环境目录下时，会自动检测到，并安装依赖项到该虚拟环境）
- uv run --python 3.12 python script.py（临时使用某版本的 python 运行东西）
- uv python pin 3.11（让当前目录固定使用某版本）
    - uv run python（会自动使用 pin 的版本的 python）

## 隔离性

- uv 管理的不同 Python 版本 | ✅ 完全独立 | 多版本开发测试
- 同一 Python 的不同虚拟环境 | ✅ 完全独立 | 项目隔离（推荐日常使用）

# 注意区分两种 pip

- uv pip install 使用 uv 实现的兼容 pip 的命令（✅ 如果必须兼容 pip 的话，始终使用这个，速度更快）
- pip install 使用 Python 自带的 pip 程序安装依赖项

# uv add 与 uv pip install

两者不能混用，新项目应该使用 `uv add`

- uv pip install 
    - 不生成/修改锁文件
    - 只是 uv 对 pip install 的兼容封装（是自己的实现，里面没有一个 pip）
    - 比 pip 更严格
    - 在项目工作流中混用 `uv pip install` 与 `uv add` 容易导致 `pyproject.toml`、`uv.lock` 与实际环境不同步
- uv add
    - 生成/修改锁文件
    - 不容易出问题

# 工具

- uv tool install xxx
- uv tool upgrade xxx

- 工具的可执行文件会注册到 ~/.local/bin 下
- 工具的本体会放到 ~/AppData/Roaming/uv/tools/my-tool/ 下
    - 虽然其 site-packages 有大量文件，但大部分是硬链接的缓存目录