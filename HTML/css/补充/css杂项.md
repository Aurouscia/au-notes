# 单位中的em表示元素的字体高度
例如`p{font-size:12px; line-height:2em}`，`line-height`的1em即为12px，2em即为24px  
- 如果font-size也用em为单位，就会找父级元素的字体大小
- 单位中的ex表示小写x的高度
- 单位中的百分比表示相对于其父级元素的相同属性

# 一个元素可以设置多个class
```html
<div class="font18px myitalic myunderline"></div>
```

# 设置元素的背景
```css
div {
    background-image: url(../images/xx.png);
    background-position/repeat/size: ...
}
```

# 设置首行缩进
`text-indent:2em`
- 负数表示反向缩进(草)，这时需要在左侧搞点padding不然就跑出去了：
```css
*{ text-indent: 2em; padding-left: 2em }
```

# 为什么有时不换行
因为默认情况下不会把单词拆开，如果是一长串数字就会不换行
```css
*{
    word-wrap: break-word;
    word-break: break-all;
}
```