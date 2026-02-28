# pnpm

使用了中央存储+哈希寻址+硬链接思维的 pnpm 比 npm 更加稳定高效

## 配置

可通过环境变量 `PNPM_HOME` `PNPM_STORE_DIR` 分别设置 pnpm 的全局包目录和存储目录

注意：全局包目录必须在 PATH 环境变量中，否则 pnpm 会拒绝安装全局包

## windows 特殊行为

硬链接无法跨盘符，所以 windows 上如果在 D 盘的项目 pnpm install，那么 D 盘根目录就会出现一个独立的 .pnpm-store 目录（无视配置）

可以考虑：创建同名的符号链接，指向其他盘的 .pnpm-store 目录以节省空间