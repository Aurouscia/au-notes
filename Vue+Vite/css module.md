# css module

css module 是社区标准，不被 w3c 承认，但受到 vue、 vite、webpack 等生态的支持

## 基本用法

`<style module>` 是 vue 特定语法，会将 CSS 编译为对象，通过 $style 访问映射后的类名（会得到代码补全提示）

```vue
<template>
  <!-- 通过 $style 访问编译后的类名 -->
  <div :class="$style.container">
    <h1 :class="$style.title">Hello</h1>
  </div>
</template>

<style module>
.container {
  padding: 20px;
}
.title {
  font-size: 24px;
  color: #333;
}
</style>
```

编译后 .container 会变成类似 .container_a3f4b2 的唯一类名，完全隔离样式。

## 进阶用法

- 命名 module

    ```vue
    <style module="classes">
    /* 通过 $classes 而非 $style 访问 */
    </style>
    ```

- 在 script 中引用
    
    引用后可以在 script 中写复杂的动态类名逻辑
    ```js
    const $style = useCssModule()
    ```

## 独立文件

- css: 以`.module.css`命名

    ```css
    /* styles/button.module.css */
    .base {
        padding: 8px 16px;
        border-radius: 4px;
    }

    .primary {
        background: #007bff;
        color: white;
    }
    ```
- vite: vite.config.ts 中可以配置哪些文件作为 css module
    
    略

- vue: 组件中 import 导入

    ```vue
    <!-- ComponentA.vue -->
    <script setup>
    import buttonStyles from './styles/button.module.css'
    // buttonStyles = { base: 'base_abc123', primary: 'primary_def456' }
    </script>

    <template>
    <button :class="[buttonStyles.base, buttonStyles.primary]">
        Click me
    </button>
    </template>
    ```

## module 间组合式复用

通过`compose`和`from`导入其他文件的类名到当前类

```css
/* base.module.css */
.btn {
  padding: 8px;
  border: none;
}

/* primary.module.css */
.primaryBtn {
  composes: btn from './base.module.css';
  background: blue;
}
```