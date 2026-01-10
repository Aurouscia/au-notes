# 颜色控制
如果需要使用css控制svg图标的颜色，而不使用其本身的颜色

## 解决方案1
不使用img标签，而使用svg标签
- img标签和background-image都会导致图片加载（样式无法被本页面控制）
- 而svg标签可以通过css控制样式
- 以下的这种写法，可以将同一段svg代码，放在不同的地方使用
```html
<svg class="icon" aria-hidden="true">
  <use xlink:href="#icon-xxx"></use>
</svg>
<svg display="none">
  <symbol id="icon-xxx" viewBox="0 0 1024 1024">
    ...
  </symbol>
</svg>
```

## 解决方案2
使用img标签，然后使用filter控制颜色
- 缺点：不一定能达到想要的效果

## 解决方案3
使用div标签，然后使用`mask`标签，将svg图案变为遮罩
- 缺点：只能实现纯色的svg图案
```css
.menu-icon{
    width: 24px;
    height: 24px;
    mask-image: url('@/assets/menu.svg');
    mask-repeat: no-repeat;
    mask-size: contain;
    background-color: black; /* 在这里设置svg图标的颜色 */
}
```
如果看不到效果，可能是：
- svg有一个占满自己的非透明背景
- svg非常大，而没有被`mask-size`控制