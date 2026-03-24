# nvm windows 版本

最好使用 choco: `choco install nvm`

nvm 会在 C 盘创建一个 nvm4w 目录，符号链接指向实际的 node 实例目录

## 镜像

`nvm install <node 版本号>` 可能较慢，可设置 node 和 npm 下载的源：

```sh
nvm node_mirror https://npmmirror.com/mirrors/node/
nvm npm_mirror https://npmmirror.com/mirrors/npm/
```