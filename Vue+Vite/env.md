# env文件

vite 支持读取 env 文件

- .env
- .env.local
- .env.development
- .env.production
- .env.development.local
- .env.production.local

其中带有修饰的优先级更高

## 指定位置

默认情况下 vite 在项目根目录寻找这些文件

也可以在 vite.config.ts 中通过 envDir 属性指定

## 使用

为了安全起见，只有 `VITE_` 开头的行能在代码中导入

```env
VITE_API_BASE="xxx"
```

```ts
const base = import.meta.env.VITE_API_BASE
```