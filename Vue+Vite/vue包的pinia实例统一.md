# Currency 模块 "找不到 Pinia" Bug 分析报告

## 问题描述

`@fickit/currency` 模块在启动后出现 "找不到 pinia" 的错误，尽管宿主应用 (`apps/web`) 已正确配置并注册了 Pinia。

## 根本原因

**`@fickit/currency` 模块的 Vite 构建配置未将 `pinia` 标记为外部依赖，导致 `pinia` 被内联打包进了库文件。**

### 详细分析

### 1. 配置不一致问题

**`package.json` 中的声明**（正确）：
```json
"peerDependencies": {
  "vue": "^3.4.0",
  "pinia": "^3.0.0"
}
```

这里正确地将 `pinia` 声明为 `peerDependency`，表示应由宿主应用提供。

**`vite.config.ts` 中的配置**（错误）：
```typescript
rollupOptions: {
  // 只排除了 'vue'，没有排除 'pinia'
  external: ['vue'],
  output: {
    globals: {
      vue: 'Vue',
    },
  },
},
```

### 2. 问题产生机制

```
┌─────────────────────────────────────────────────────────────────┐
│                        宿主应用 (apps/web)                        │
│  ┌─────────────────┐                                            │
│  │  createPinia()  │ ──创建 Pinia 实例 A                         │
│  │  app.use(pinia) │ ──将实例 A 关联到 Vue 应用                   │
│  └─────────────────┘                                            │
│                           │                                     │
│                           ▼                                     │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │           引入 @fickit/currency/dist/index.js            │   │
│  │  ┌─────────────────────────────────────────────────┐    │   │
│  │  │   打包进来的 pinia 副本（内联在库文件中）          │    │   │
│  │  │   createPinia() ──创建 Pinia 实例 B（未与Vue关联）  │    │   │
│  │  └─────────────────────────────────────────────────┘    │   │
│  │                           │                             │   │
│  │                           ▼                             │   │
│  │  useCurrencyConverterStore() ──从实例 B 获取 store      │   │
│  │  ❌ 报错：找不到 pinia（实例 B 未与 Vue 应用关联）        │   │
│  └─────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
```

### 3. 为什么报错

当 `CurrencyConverter.vue` 调用 `useCurrencyConverterStore()` 时：
1. 该组件从打包进来的 `pinia` 副本（实例 B）获取 store
2. 但这个 `pinia` 实例 B 从未通过 `app.use(pinia)` 与 Vue 应用关联
3. Pinia 内部检查到没有活跃的 Pinia 实例，抛出 "找不到 pinia" 错误

## 修改方案

### 修改文件

**`fic-kit-frontend/packages/@fickit/currency/vite.config.ts`**

### 修改内容

```typescript
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import dts from 'vite-plugin-dts'
import { resolve } from 'path'

export default defineConfig({
  plugins: [
    vue(),
    dts({
      insertTypesEntry: true,
    }),
  ],
  build: {
    lib: {
      entry: resolve(__dirname, 'src/index.ts'),
      name: 'FicKitCurrency',
      formats: ['es'],
      fileName: 'index',
    },
    rollupOptions: {
      // 将 pinia 标记为外部依赖，使用宿主应用提供的实例
      external: ['vue', 'pinia'],
      output: {
        globals: {
          vue: 'Vue',
          pinia: 'Pinia',
        },
      },
    },
  },
})
```

### 修改说明

| 修改项 | 原值 | 新值 | 说明 |
|--------|------|------|------|
| `external` | `['vue']` | `['vue', 'pinia']` | 将 pinia 标记为外部依赖，不参与打包 |
| `globals.pinia` | 无 | `'Pinia'` | 指定 pinia 的全局变量名（UMD/IIFE 格式需要） |

## 验证步骤

1. 修改 `vite.config.ts` 后重新构建 currency 模块：
   ```bash
   cd fic-kit-frontend/packages/@fickit/currency
   npm run build
   ```

2. 在宿主应用中测试：
   ```bash
   cd fic-kit-frontend/apps/web
   npm run dev
   ```

3. 确认 CurrencyConverter 组件正常工作，不再报错。

## 最佳实践建议

1. **peerDependencies 与 external 保持一致**：所有声明为 `peerDependencies` 的依赖，都应在 `rollupOptions.external` 中排除。

2. **库模式打包原则**：
   - 框架级依赖（Vue、Pinia、React 等）应作为 peer dependency
   - 工具库可以内联打包或作为普通 dependency
   - 同一生态内的共享状态管理库必须使用宿主实例

3. **避免重复打包**：重复打包同一库会导致：
   - 包体积增大
   - 运行时出现多个实例
   - 状态管理失效（如此处的 Pinia 问题）
