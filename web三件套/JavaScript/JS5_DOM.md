# JavaScript DOM（文档对象模型）

DOM：可以动态修改网页的内容、样式、结构。

- **核心 DOM**：提供了操作 HTML 和 XML 文档的各种属性和方法（公共）
- **XML DOM**：只针对 XML 的接口
- **HTML DOM**：针对 HTML 操作的接口
- **Event DOM**：事件对象，提供了常用的事件比如 `onclick`、`onload`、`onchange`
- **CSS DOM**：JS 访问 CSS 各种属性

## HTML 节点树

`<html>` 为根元素，一个网页只有一个根元素。
`<head>` 和 `<body>` 是其子元素。
`<h1>` 是 `<body>` 的子元素。
文本是 `<h1>` 子元素，是文本节点。

这里分为两类：标记和文本节点。

## DOM 节点类型

DOM 对于 HTML 文档定义了五个节点类型：

1. **document 文档节点**
   - 对应于当前网页文档，对应于 `document` 对象，最高级，包含所有元素

2. **element 元素节点**
   - 所有 HTML 标记

3. **text 文本节点**
   - 最底层的节点

4. **attribute 属性节点**
   - 指元素对应的属性构成的节点列表

5. **comment 注释节点**

## 节点的访问

> 【已经淘汰，应该用 `getElementById()` 和 `getElementsByTagName()`】

增加、删除、修改、读取：

- `firstChild`：第一个子节点（文档里的上下顺序）（是属性）
- `lastChild`：最后一个子节点
- `nodeName`：节点名称
- `nodeValue`：节点的值（仅文本有）
- `childNodes`：子节点列表，是个数组

访问 HTML 节点：
```javascript
var node_html = document.firstChild;
// 或者
document.documentElement;
```

再访问 body 节点：
```javascript
var node_body = node_html.lastChild;
// 或者
document.firstChild.lastChild;
// 或者
document.body;
```

输出 table 节点得到：`object HTMLTableElement`

从上到下，先执行：
```javascript
alert(document.body.nodeName);
```
再出现 body，所以需要写在 `onload` 里面。

## 节点属性的修改

- `setAttribute(attName, attValue)` - 给一个节点增加属性（必须已有的属性才行，不能自定义）
- `removeAttribute(attName)` - 删除节点某属性
- `getAttribute(attName)` - 取得节点某属性的值

```html
<img/>
```
```javascript
img.setAttribute("src", "xxx/xxx.png");
```

例子：
```javascript
function Init() {
    divNode = document.getElementById("d");
}
var s = "width:600px;height:300px;background-color:#f0f0f0";
function change() {
    divNode.setAttribute("style", s);
}
```

（只能加行内样式，加不了内嵌式）

## 节点创建和删除

- `createElement(tag)` - 创建一个指定标记
  - 如：`parent` 或 `document.createElement("img")`，不要尖括号
- `createTextNode(text)` - 创建文本节点
  - 如：`document.createTextNode("年龄");`

只创建没有用，追加之后才有效：

- `parentNode.appendChild(node)` - 将一个节点追加到某父节点下（最后一个）
- `parentNode.insertBefore(node, current)` - 将一个节点插到 `current` 节点之前
- `parentNode.removeChild(node)` - 移除一个子节点

## CSS DOM

给每个元素对象添加样式。

每个 HTML 标记都有 `style` 属性，对应对象元素的 `style` 属性。
每个标记都对应一个元素对象，元素对象的属性和标记的属性一模一样。

```html
<div id="box"></div>
```
```javascript
var obj = document.getElementById("box");
obj.width = "300";  // HTML 的不要单位
obj.style = "padding:20px;";  // CSS 的要单位
```

`style` 也是个对象，与 CSS 中的属性一一对应。
（如果之前没给 `style` 赋值字符串）可以写：
```javascript
obj.style.border = "1px solid #FF0000";
obj.style.position = "absolute";
```

### CSS 属性与 style 对象属性的转换

如果是一个单词的属性，CSS 与 `style` 一模一样，如：
```javascript
obj.style.width = "400px";
```

如果是多个单词的属性，转成 `style` 对象属性时：
连字符去掉，改成驼峰命名法。

- `background-color` 变成 `obj.style.backgroundColor = "#FF0000";`
- `font-size` 变成 `obj.style.fontSize = "18px";`

都相当于是行内式。

## HTML DOM

每一个 HTML 标记都对应一个元素对象。
HTML 标记属性与元素对象属性一模一样（继承核心 DOM）。

```javascript
imgObj.width = 400;
imgObj.src = "images/xx.png";
```

### 获取 HTML 标记对象

```javascript
document.getElementById();
```

或者批量获取：
```javascript
parentNode.getElementsByTagName();
```

返回一个数组：
```javascript
var arr = parentNode.getElementsByTagName("li");
```

取得某父节点下的标记为 `<li>` 的所有对象，可以减少 id 的数量，简化程序。

或者通过 `name`（和 id 一回事）属性访问（一般用于表单）：
```html
<form name="form1" action="login.php" method="post" onsubmit="return checkForm()">
    用户名:<input type="text" name="username"/>
</form>
```
```javascript
function checkForm() {
    if (document.form1.username.value == "") {
        window.alert("用户名不能为空");
        return false;  // 事件的返回值为假，阻止表单提交
    } else {
        return true;
    }
}
```

### 通用属性

- `tagName`：取得标记的名称，同 `nodeName` 属性
- `innerHTML`：指标记对中的（含 HTML 标记）文本（双边标记才有）
- `style`
- `id`
- `offsetWidth`：指元素的宽度，不含滚动条中内容（可视宽度）（只读）
- `offsetHeight`：指元素的高度，不含滚动条中内容（可视高度）（只读）
- `scrollWidth` / `scrollHeight`：指元素的总宽高，含滚动条中内容
- `scrollTop` / `scrollLeft`：指内容向上/左滚动进去了多少，如果没有滚动条，值为 0（可写）

> 【↑ 可以用来做动画】

- `onscroll`：当内容滚动时，发生的事件

自动隐藏多余的文字，出现滚动条：
```css
overflow: auto;
```

## Event DOM

主要：鼠标键盘和网页交互。

| 事件 | 事件句柄 |
|------|----------|
| click | onclick |
| dblclick | ondblclick |
| mouseover | onmouseover |
| mouseout | onmouseout |

每个事件都对应一个事件句柄属性，该属性是 HTML 标记的事件属性：
- `onload`
- `onchange`
- `onscroll`
- `onsubmit` - 当单击提交按钮时
- `onreset` - 当单击重置按钮时
- `onfocus` - 当（文本框）获得焦点时
- `onblur` - 当（文本框）失去焦点时（在文本框外单击）
- `onselect` - 当选择下拉菜单元素时

每个 HTML 标记都有事件句柄属性：
```html
<img onclick="xxx()">
```

HTML 标记对应的对象也具有事件句柄属性【必须全小写】（JS 分大小写）：
```javascript
obj.onclick = xxx;  // 函数的地址而非返回值（所以不要括号）
```

给图片标记绑定事件：
```html
<img onclick="show(this)">
```
```javascript
function show(obj) {
    // xxx
}
```

给图片对象绑定事件：
```javascript
img.onclick = jmp;  // 函数的地址而非返回值（所以不要括号）
// 或者匿名函数
img.onclick = function() {
    location.href = "http://www.xxx.com";
};
```

链接的图片，最好：
```javascript
obj.style.cursor = "pointer";
```

### 事件返回值

阻止默认动作执行：
```javascript
onclick = "return false"  // 或 "return xxx()"
onsubmit = "return checkForm()"
```
