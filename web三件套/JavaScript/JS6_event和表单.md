# JavaScript 事件和表单

## Checkbox 技巧

```php
echo "<input type=\"hidden\" name=\"".$name."\" value=false>";
echo "<input type=\"checkbox\" name=\"".$name."\" value=true>";
```

一般情况下，不打勾会不提交任何东西，这样写可以让打勾提交 `true`，不打勾提交 `false`。

## Event 对象

事件发生时，会向调用函数传递一个 `event` 对象，记录当前环境信息。
一个事件对应一个 `event` 对象，`event` 对象短暂存在。

```html
<img id="img01" src="images/01.jpg" onclick="get_xy(event)"/>
```
```javascript
function get_xy(e) {  // e 是形参，怎么写随便
    alert(e);
}
```

或者：
```javascript
var obj = document.getElementById("img01");
obj.onclick = get_xy;
function get_xy(e) {
    alert(e);
}
```

### Event 对象属性

- `e.clientX`、`e.clientY` - 点击的坐标（相对于窗口（不包括顶上滚动掉没显示的）左上角）
- `e.pageX`、`e.pageY` - 点击的坐标（相对于网页（包括顶上滚动掉没显示的）左上角）
- `e.screenX`、`e.screenY` - 点击的坐标（相对于屏幕左上角）

## 表单对象

### form 的属性

- `name`：表单名称
- `method`：提交方式，`get` 和 `post`（这个要和后端人员商量）
- `action`：处理程序
- `enctype`：加密方式

### form 的方法

- `submit()` - 提交表单
- `reset()` - 重置表单（同 reset 按钮）

### form 的事件

- `onsubmit()` - 单击提交按钮后，发送数据前发生
- `onreset()` - 重置时发生

### 表单提交三种方法

1. submit 按钮结合 `onsubmit` 事件
2. submit 按钮结合 `onclick` 事件
3. button 按钮结合 `onclick` 事件和 `submit()` 方法

`onsubmit` 可以接受返回值。

form 里写 `onsubmit="return check()"`，可以在条件不满足时阻止表单提交。
或者：`input type="submit"` 里写 `onclick="return check()"`，也行。

## 文本框对象

### 属性

`name`、`value`、`size`、`disabled`、`readonly`、`maxlength`

### 事件

- `onfocus` - 单击之（获得焦点）
- `onblur` - 单击别处（失去焦点）

### 方法

- `select()`
- `focus()`
- `blur()`

> 最好：`<form>` 套 `<table>`

## 复选框

- `name`
- `value`
- `checked`（`true`/`false`）

`<fieldset></fieldset>` 形成一个框。

如果几个对象 `name` 都一样，获取将产生一个数组：
```html
<input type="checkbox" name="hobby"/>
<input type="checkbox" name="hobby"/>
```
```javascript
var arr = document.form1.hobby;
```

## 二级联动菜单

可以使用以字符串为索引的数组（比较复杂）：
```javascript
var cityList = new Array();
cityList['北京市'] = ['xx区', 'xx区', 'xx区'];
```

### 下拉列表 select 对象的常用属性

- `options` 数组，返回所有 option 组成的数组
- `name`
- `selectedIndex` - 当前选中的索引号
- `value` - 当前选中的 option 的值
- `length` - 长度，option 的个数

### 选项 option 的常用属性

- `text` - 指一对 option 标签中间的文本
- `value`
- `index` - 指其索引号
- `selected`

给下拉列表写入 option 时，应该先给其指定高度，否则会显示异常：
```javascript
province.length = arr_province.length;
```

然后创建 option 对象：
```javascript
province[i].text = arr_province[i];
province[i].value = arr_province[i];
```

数组的 `length` 是可读可写的。
