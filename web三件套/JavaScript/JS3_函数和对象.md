# JavaScript 函数和对象

## 定义函数

```javascript
function 函数名(参数, 参数, 参数) {
    // xxxx;
    return xx;
}
```

### 匿名函数

没有标识符的函数：

```javascript
var b = function(a) {
    return a * a;
};
```

可以赋给数组元素：

```javascript
arr[2] = function(a) {
    return a * a;
};
var someFunc = arr[2];
someFunc(10);
```

```javascript
var a = getInfo;
a();
```

直接将函数名赋给变量，是传递了函数的地址。

### 变量作用域

- 函数内：`var` 定义的变量为局部变量
- 省略 `var` 定义的变量为全局变量

## 数据类型

**基本数据类型**（即为值类型）：
- `string`、`number`、`boolean`、`undefined`、`null`

**复合数据类型**（即为引用类型）：
- `array`、`object`、`function`

例子：
```javascript
var a = {name: "xxx", age: 24};
var b = a;  // 传递了地址，实际上依然是同一个东西
alert(b.age);
a.age = 30;
alert(b.age);  // 会改变
```

## JS 内置类

- `String`
- `Math` - 里面有 `Math.sqrt()`、`Math.abs()` 等
- `Array` - 里面有 `Array.length`、`Array.sort()` 等
- `Number` - 里面有 `Number.toFixed()`
- `Date` - 里面有 `getDay()`、`getFullYear()`
- `Function`
- `RegExp` - 正则表达式

## 创建 Object 实例

> 【自定义的对象，可以加任意属性和方法】

### 1. 使用 new 和 Object 构造函数

```javascript
var obj = new Object();
obj.name = "xxx";  // 会自动增加属性
obj.school = "xxx";
obj.showInfo = function(xx) {
    // xxx;
};
```

增加一个方法，左边不需要括号（表示函数本身而不是返回值），右边是函数定义。

使用时：
```javascript
document.write(obj.name);
obj.showInfo(xx);
```

也可以写：`obj['name']` 访问。

### 2. 使用大括号

```javascript
var obj = {
    name: "xx",
    school: "xx",
    showInfo: getInfo  // 不要括号，表示函数本身而非其返回值
};
```

## 类

### 1. 用 function 写个构造函数

```javascript
function MyClass() {
    this.id = 5;
    this.name = 'myclass';
    this.getName = function() {
        return this.name;
    };
}
var my = new MyClass();
```

### 2. 在函数中声明对象并添加属性，然后返回

```javascript
function MyClass() {
    var obj = {'id': 2, 'name': 'myclass'};
    return obj;
}
var my = new MyClass();
```

### 创建多个类似的自定义对象

```javascript
function createCar(color) {
    var tempCar = new Object();
    tempCar.color = color;
    tempCar.showColor = function() {
        alert(this.color);
    };  // 也可以把同样的函数提到外面避免重复
    return tempCar;
}
```

## String 类

### 创建 String 类的实例

1. 使用 `new` 和 `String` 构造函数：
   ```javascript
   var str1 = new String("xxxx");
   ```

2. 直接写：
   ```javascript
   var str1 = "xxxx";
   ```

### String 对象的方法

1. `charAt(index)` - 查找特定索引号的字符
   ```javascript
   var char = xxstr.charAt(0);
   ```
2. `indexOf(char)` - 从左到右查找子字符串位置
3. `lastIndexOf(char)` - 从右到左查找子字符串位置
4. `substr(startIndex, length)` - 提取子字符串（length 可选）
5. `substring(startIndex, endIndex)` - 提取子字符串
6. `split(分割符)` - 返回一个数组
7. `toLowerCase()` / `toUpperCase()` - 大小写转换
8. `localeCompare()` - 用当地语言的排序方式比较大小
   ```javascript
   var arr = ["xxx", "xxx", "xxx", "xxx"];
   arr.sort(orderby2);  // 提供比较方法
   function orderby2(str1, str2) {
       return str1.localeCompare(str2);
   }
   ```

## Date 类

### 创建 Date 对象的实例

1. 以当前的日期时间创建：
   ```javascript
   var today = new Date();
   ```

2. 指定日期字符串：
   ```javascript
   var yesterday = new Date("1990/10/12 10:08:08");
   // 或者写成
   var someday = new Date("2030-10-10");
   ```

3. 指定年月日时分秒毫秒：
   ```javascript
   var someday = new Date(year, month, day, hour, minute, second, msecond);
   ```
   可以从右向左省略，但至少有年月日三个。

### Date 对象的方法

1. `getFullYear()` - 取得四位的年份
2. `getMonth()` - 取得当前月份，0-11（加 1 才对）
3. `getDate()` - 取得当前几号，取值 1-31
4. `getHours()`
5. `getMinutes()`
6. `getSeconds()`
7. `getMilliseconds()`
8. `toLocaleString()` - 用中文显示
9. `toLocaleDateString()`

## 定时器

回忆：meta 标签里的 refresh 属性不适合做时钟，因为刷新整个网页十分耗费资源。

应该用：`window.setInterval(JS代码, 毫秒数)`

> 【动画就是用这个做的】

```javascript
window.setTimeout(JS代码(一般是个函数), 毫秒数);
window.setInterval(JS代码, 毫秒数);
```

- `setTimeout`：延迟一段时间后执行一次
- `setInterval`：隔一段时间执行一次

## Math 静态类

- `Math.abs()` - 绝对值
- `Math.ceil()` - 向上取整（天花板）
- `Math.floor()` - 向下取整
- `Math.round()` - 四舍五入
- `Math.sqrt()` - 平方根
- `Math.random()` - 返回 0-1 之间的小数

返回特定最大值和最小值之间的随机数：
```javascript
function editBg() {
    var min = 100000;
    var max = 999999;
    var color = Math.floor(Math.random() * (max - min) + min);
    document.bgColor = "#" + color;
}
```
