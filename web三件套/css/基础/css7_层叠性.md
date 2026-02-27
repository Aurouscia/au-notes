# 层叠
指的是优先级高的覆盖优先级低的

## 最顶级的优先关系：
	网站作者的设置>用户对浏览器的设置>浏览器默认设置>html规范的默认样式

## 如果是平级的，那就看css的来源
	行内式>嵌入式>外链式

## 如果是平级的，那就看选择器的特殊性
	id选择器>类选择器>标签选择器>通配符(差不多那个意思)

## 如果是组合的选择器，可以通过总权重计算出优先级
- id选择器(100)	
- 类选择器(10)	属性选择器(10)
- 标签选择器(1)	伪选择器(1)
- 通配(0)

## 如果是平级的，那就看页面中出现的先后顺序
	后出现的会覆盖先出现的

## 继承得到的样式没有权重
	会被任何其他样式覆盖

## 手动指定最高优先级
```css
#header{color:red!important;}
```

## @layer
显式声明哪些样式优先
```css
@layer reset, base, theme, utilities;
/* utilities 最强，reset 最弱 */

/* 方式 A：在样式块前加层名 */
@layer base {
  body { margin: 0; font-family: system-ui; }
}
@layer theme {
  body { background: #111; color: #eee; }
}
/* 方式 B：用 @import 直接指定层 */
@import url(reset.css) layer(reset);
```
- 层内 正常规则输时，!important 仍能提高优先级。
- 层之间 的优先级 高于 !important：即 后层正常规则 > 先层 !important。