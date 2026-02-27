# JavaScript BOM（浏览器对象模型）

BOM 是浏览器对象模型。
一组对象的集合，提供了访问和操作浏览器各组件的方法。
有 `window` 对象、`history` 对象和 `location` 对象。

DOM 是文档对象模型。
一组对象的集合，提供了访问和操作网页各元素的方法。
所有 HTML 标记都是网页元素，都可以用 JS 操作。
操作指增加、删除、修改、读取四种操作。

BOM 和 DOM 并不属于 JS，是 W3C 开发的。

## BOM 对象层次

- `window` 对象 - 代表一个浏览器窗口【其他对象都是其子级】
- `document` 对象 - 代表网页文档
- `history` 对象 - 代表历史记录，可以前进后退
- `location` 对象 - 代表顶上的地址栏，可以修改和读取，禁止外链访问
- `screen` 对象 - 代表显示器，可以取得分辨率、色深等
- `navigator` 对象 - 代表浏览器软件

### document 对象

对 DOM 中所有对象的一个引用。
比如：
```
<html>
  <body>
    <table>
      <tr>
        <td>内容（文本节点）
```

## window 对象

`window` 对象是最顶层的对象，是全局对象，在任何地方都可以直接调用其属性和方法。
因为每次都写 `window` 太麻烦，所以在写子对象时可以省略。

```javascript
window.document.body.bgColor = "#FF0000";
// 可以直接写
document.body.bgColor = "#FF0000";
```

### 属性

（可以直接写属性名，不用 `window`）

- `closed` - 判断窗口关闭状态，一个布尔值
- `name` - 浏览器窗口的名称，一般用于超链接的 target 使用
- `innerWidth` / `innerHeight` - 窗口宽高，不含滚动条、菜单栏
- `outerWidth` / `outerHeight` - 窗口宽高，含滚动条、菜单栏等

> 起名不要起这些：
> - `_top` 代表顶层窗口
> - `_parent` 代表父窗口
> - `_self` 代表当前窗口

### 方法

- `alert()` - 弹出一条信息
- `prompt()` - 输入对话框
- `confirm()` - 确认对话框

例子：
```javascript
function confirmDel() {
    if (window.confirm("确认删除？")) {
        window.alert("已经删除");
    }
}
```
```html
<a href="javascript:void(0)" onclick="confirmDel()">删除</a>
```

- `close()` - 关闭窗口
- `open(url, name, options)` - 打开新窗口

```javascript
var winObj = window.open(url, name, options);
```

返回一个窗口对象，这个 `winObj` 可以用所有窗口的属性和方法。

参数：
- `url`：是要打开的文件的地址，如果为空，打开一个空网页
- `name`：给打开的窗口起个名字，一般用于链接的 target 属性
  - 【name 一般填 `_self`、`_blank`，代表弹出到本窗口和新窗口】
- `options`：
  - `width`、`height`、`left`（距离屏幕左端多远）、`top`
  - `menubar`（是否显示菜单栏 yes/no）、`toolbar`、`status`
  - `scrollbars`、`location`

例子：
```javascript
var winObj = window.open("xxx.html", "wd1", "width=400,height=300,left=xx,menubar=no,toolbar=yes");
```
```html
<a href="xxx" target="wd1">看看在哪打开</a>
```

target 写创建窗口时起的名字，或者 `_self`、`_top`、`_parent`、`_blank`。

> 一般不用，因为基本上都会被浏览器阻止。

### 定时器

```javascript
var timer = window.setTimeout(code, milliSec);
```

返回一个延时器。
- `code` 是任何 JS 代码，一般是执行一个函数
- `milliSec` 是正整数，停留几个毫秒之后执行这个（仅一次）

```javascript
var timer = window.setTimeout("win.close()", 2000);  // 两秒后关闭窗口
clearTimeout(timer);  // 将其停止（并不会变成 null）
```

```javascript
var timer = window.setInterval(code, milliSec);
```

返回一个周期定时器。

```javascript
clearInterval(timer);  // 将其停止（并不会变成 null）
```

例子：
```html
<script type="text/javascript">
    var i = 1;
    var obj;
    var timer;
    function Init() {
        obj = document.getElementById("result");
    }
    function tick(obj) {
        var str = "该程序已经运行" + i + "秒";
        obj.value = str;  // 区分于 innerHTML
        i++;
    }
    function start() {
        if (timer == null) {
            timer = window.setInterval("tick(obj)", 1000);
        }
    }
    function stop() {
        window.clearInterval(timer);
        timer = null;
    }
</script>
```

`onload` 事件：指网页中所有内容（文字图片视频）加载完毕时，去调用一个 JS 函数。

### innerHTML

在一对标签之间插入 HTML 代码（字符串）。

## screen 对象

代表显示器屏幕。

- `width`、`height`
- `availWidth`：不含任务栏
- `availHeight`

## navigator 对象

代表浏览器软件。

- `appName`
- `appVersion`
- `platform` - 什么操作系统
- `systemLanguage`
- `userLanguage`

## location 对象

代表地址栏。

```
http://www.sina.com.cn/about/index.html?username=yao&password=123456#top
|------|---------------------|------------------|----------------------------------|----|
协议        域名/主机名              目录和文件名            查询字符串(传值用)           锚点
```

- `href`：完整的 URL 地址
- `protocol`：协议
- `host`：主机名 `www.sina.com.cn`
- `pathname`：路径和文件名 `about/index.html`
- `search`：查询字符串 `?username=yao&password=123456`
- `hash`：锚点名 `#top`

> 【以上各个属性可以直接赋值，会自动刷新网页】

手动刷新：
```javascript
onclick="javascript:location.reload()"
```

## history 对象

代表历史记录。

- `length`：条数
- `back()`：相当于返回按钮
- `forward()`：相当于前进按钮
- `go(x)`：跳到某一条，-1 代表上一页，1 代表下一页
