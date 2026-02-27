# JavaScript 基础知识和变量

服务器发给客户端的文件由三部分组成：HTML、CSS 和 JS。
JS 属于客户端运行的东西，由浏览器直接解释执行。

JS 是一种小型的、轻量级的、面向对象的、跨平台的客户端脚本语言。
和浏览器捆绑在一起，只要有浏览器，就有 JavaScript。
Windows、Unix、Linux、Mac、Android、iOS 都支持这个东西。

创建一个 `xx.html` 文件，用 VS Code 打开：

```html
<head>
<script type="text/javascript">
// 此处写 JS 代码
</script>
</head>
```

JS 分大小写。

## JS 能干什么

1. 表单填写验证，看填的内容是否合法
2. 动态 HTML，自动动画效果
3. 用户交互动画效果（需要鼠标键盘介入）

## 客户端输出方法

`document.write("str")` 在网页中输出一个字符串，相当于在 `<body>` 写入一段 HTML 文本。
此处 `document` 是对象，`write` 是这个对象的方法。
`str` 是要输出的（HTML）内容，要用双引号引起来。
要加样式，可以直接往里写 HTML 的内容：

```javascript
document.write("<font size='7' color='red'>xxxxxxx</font>");
```

几乎所有 HTML 标签（除了 `head`、`html` 和 `body`）都可以写在里面，可以用来调试错误。

`window.alert("str")` 会弹出一个带有 `str` 的窗口。
`window` 代表浏览器窗口对象，`alert` 是其方法。
在用户点确定之前，程序会卡住，后面的东西加载不出来。

## HTML 中引入 JS 程序的方法

1. **内嵌式**
```html
<script type="text/javascript">
    // xxx
</script>
```

2. **外链式**
```html
<script type="text/javascript" src="外部js文件url"></script>
```
外部 JS 文件扩展名为 `.js`。

3. **行内式**
每个 HTML 标记都有事件属性（如：鼠标单击、双击、放上、移出）。
行内 JS 需要配合事件属性使用：
```html
<img src="xxx" onclick="javascript:window.alert('xxxx')">
```

## 变量

变量是什么：一个存储信息的盒子，里面可以装一个字符串、一个数字、一个 div 的引用等。

变量定义：
```javascript
var 变量名 = 变量值;
var name, age, sex, edu, school;
```

**变量名规则：**
- 只能是 a-z、A-Z、0-9 和下划线
- 以字母或下划线开头（不能以数字开头）
- 变量名区分大小写

多个单词组成的变量，可以驼峰命名也可以全小写下划线连接：
- 驼峰命名法：`magicNum`
- 下划线命名法：`magic_num`

## 数据类型

**基本数据类型：**
- `string` - 字符串
- `number` - 数字
- `boolean` - 布尔（是否）
- `undefined` - 未定义
- `null` - 空

**复合数据类型：**
- `array` - 数组
- `object` - 对象
- `function` - 函数

### 数值型

```javascript
var a = 0;
var a = 1.5;
```

值可能为 `NaN`，代表数据转换失败（不是个数字）。

### 字符串型

用单引号或者双引号引起来的字符串：
```javascript
var a = 'abc';
```

双引号里可以套单引号，单引号里可以套双引号。
如果要套相同的，需要使用转义字符。

### 布尔型

代表两种状态：
```javascript
var isMarried = true;  // 或 false
```

### 未定义型

代表变量没有赋值过，类型为 `undefined`，值也为 `undefined`：
```javascript
document.write(typeof(isRead));
```

### 空型

代表一个对象不存在，用 `typeof()` 判断类型为 `object`，值为 `null`：
```javascript
var str = window.prompt("xx");
document.write(typeof(str) + "," + str);
```

- 如果点取消，会返回类型 `object`，值 `null`
- 如果点确定，会返回类型 `string`，值 `""`

> 注：`window.prompt` 会弹出让用户输入内容的对话框，并返回用户填写的字符串。

### typeof 运算符

```javascript
string typeof(var)
```

返回一个字符串，指示 `var` 是什么类型的。
返回值有六种：`"string"`、`"number"`、`"boolean"`、`"undefined"`、`"object"`、`"function"`。

> 当数据类型为 `array`、`object` 和 `null` 时返回值为 `"object"`。

## 数据类型转换

### 1. 转为数值

- 非纯数字的字符串和 `undefined` 变成 `NaN`
- `true` 和 `false` 变成 `1` 和 `0`
- `null` 变成 `0`
- 纯数字的字符串变成相应数字

### 2. 转为布尔

- 非空字符串变为 `true`
- 空字符串变为 `false`
- `0` 变为 `false`
- 非 0 的任何数都变为 `true`
- `null` 变成 `false`

### 3. 转为字符串

略。

## 强制类型转换

当自动不好使的时候用：
- `Number()`
- `Boolean()`
- `String()`
