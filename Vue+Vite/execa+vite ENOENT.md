# ENOENT

代表找不到可执行文件

注意区分 `execa("vite", ["build"])` 与 `execa("vite build")`

后者是错误写法，会把 vite build 当作一个可执行文件（然后找不到，显示 ENOENT）

前者才是使用 vite，然后传入参数 build