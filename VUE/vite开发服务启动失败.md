# 错误信息
No matching export in "node_modules/unicorn-magic/default.js" for import "traversePathUp"

# 分析
- 包`unicorn-magic`提供了`default`和`node`两种导出，但vite（打包器）对于这个web项目只会使用default导出
- 包`execa`依赖于`unicorn-magic`的node导出的功能，所以应该避免`import { execa } from 'execa'`进入运行时代码
- 在项目中找到一个ts脚本，其中混合着execa的使用以及interface定义
- 其他运行时代码import了该脚本的interface定义
- vite在启动开发服务器时解析到了该脚本，试图导入`unicorn-magic`的default导出，引发异常

# 结果
将问题脚本的interface分离到单独文件，问题消失