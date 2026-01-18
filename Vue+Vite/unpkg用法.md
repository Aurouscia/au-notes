# 概念
unpkg是一个免费的开源CDN，提供了各种npm包的最新版本，方便开发者在项目中直接引入使用。

# 用法
1. （如果需要用自己的包）把包传到npm上，unpkg会自动更新
2. 直接在html文件中引入`https://unpkg.com/包名@版本号`，可使用importmap指定包的版本和路径

# 坑
- 不会自动解析依赖项目，每个依赖项都要加入到importmap中，供引入的包直接使用名称（例如“vue”）引入
- 如果使用了scoped包，不会自动根据包名找到入口文件，必须使用具体路径精确到入口js文件

# 示例

```html
<!doctype html>
<html>
<head>
  <meta charset="utf-8" />
  <title>Demo</title>
  <script type="importmap">
  {
    "imports": {
      "vue": "https://unpkg.com/vue@3/dist/vue.esm-browser.js",
      "color-convert": "https://unpkg.com/color-convert",
      "color-name": "https://unpkg.com/color-name"
    }
  }
  </script>
  <link rel="stylesheet" href="https://unpkg.com/@aurouscia/au-color-picker@1.0.2-dev.4/dist/au-color-picker.css">
</head>
<body>
  <div id="app">
    <au-color-picker/>
  </div>

  <script type="module">
    import { createApp } from 'vue'
    import { AuColorPicker } from 'https://unpkg.com/@aurouscia/au-color-picker@1.0.2-dev.4/dist/au-color-picker.es.js'

    createApp({
      components: { AuColorPicker }
    }).mount('#app')
  </script>
</body>
</html>
```