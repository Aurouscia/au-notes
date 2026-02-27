# JavaScript 数组和流程控制

## 从字符串中提取数值

- `parseInt(str)` - 从左到右提取整数，遇到非数字字符就停止
- `parseFloat(str)` - 从左到右提取小数，遇到非数字字符就停止

例如：
```javascript
var a = parseInt('123abc');  // a 中就是 123
```

## 三元运算符

`x ? x : x` 三元运算符实例：

```javascript
var isMarried = true;
var result = isMarried ? "已婚" : "未婚";

var a = 100;
var b = 200;
var max = a > b ? a : b;

var school;
var result = school ? school : "未填写";
```

如果没有写，数据类型为 `undefined` 或者 `null`，转为布尔值是 `false`，就会选择后面那个。

## 比较运算符

### == 和 ===

- `==` 为 `true` 代表值一样
- `===` 为 `true` 代表值和类型一样（严格等于）

> 【`=` 是赋值，`==` 是判断是否相等，不要搞混，要不然检查都查不出来】

## 逻辑运算符

- `&&` - 表示"且"，两边都真才为真
- `||` - 表示"或"，一边为真就为真
- `!` - 表示"非"，将是非倒置

## 其他运算符

- `new` 运算符：创建一个对象
  ```javascript
  var today = new Date();
  ```
- `delete` 运算符：删除对象的属性或者数组的元素
- `void` 运算符：作用于任何值都返回 `undefined`

## 条件语句

### if-else

```javascript
if (判断条件) {
    // xxx
} else {
    // xxx
}
```

JS 的 `if` 和 `else` 必须要大括号。

### switch-case

和 C 语言的一样：

```javascript
switch (变量) {
    case 值1:  // 这里是冒号
        // xxxxx
        break;
    case 值2:
        // xxxxx
        break;
    default:
        // xxxxx
}
```

变量等于值几就会进入哪个分支。

## 数组

> 【数组是一些数据的集合，JS 中可以把类型不同的数据放进一个数组里】

### 创建数组对象

```javascript
// 直接把元素写进去
var arr = new Array(10, 20, 30, 40);

// 或者使用中括号创建数组
var arr = [10, 20, 30];

// 空数组
var arr = new Array();
arr[0] = new Array();  // 数组元素可以是数组，实现多维数组
arr[0][2] = "abc";

// 二维数组
var arr = [[1, 5], [2, 4], [2, 3]];
```

### 数组操作

- **扩大数组**：直接往后面写入就会自动扩大，不用担心数组越界
- **访问数组**：`arr[i]`，访问数组的第 i 个元素
- **删除数组元素**：`delete arr[x]`
  - 内容被清除，所占空间还在

### 数组对象的属性和方法

```javascript
var arr = new Array(xxx, xxx, xxx);
arr.xxx  // 来使用属性和方法
```

1. `length` - 统计长度：`var len = arr.length;`
2. `join(连接符)` - 将数组中的元素以指定的连接符连成字符串：`arr.join('+')`
3. **删除**
   - `shift()` - 删除第一个，后面补上来，长度减一
   - `pop()` - 删除最后一个，长度减一
   - `delete arr[x]` - 内容被清除，所占空间还在
4. **增加**
   - `unshift()` - 开头添加一个元素
   - `push()` - 结尾添加一个元素
   
   > 3 和 4 都是在原数组进行操作，不要再创建变量
5. `reverse()` - 反转数组
6. `sort(指定排序规则的函数)` - 排序
   ```javascript
   arr.sort(orderby);
   function orderby(a, b) {
       return a - b;
   }
   ```

## 循环

### while 循环

```javascript
while (判断条件) {  // true 就会继续循环，false 就会停止循环
    // 循环体
}
```

### for 循环

可以做初始化和循环变量更新：

```javascript
for (初始化; 判断条件; 循环变量更新) {
    // xxx
}
```

例如：
```javascript
for (var i = 1; i < 10; i++) {
    // xxx
}
```

### do-while 循环

```javascript
do {
    // xxxx
} while (判断条件);
```

一样的，只不过判断在循环后面，循环体至少执行一次。

## 循环控制

- `break` - 跳出循环
- `continue` - 停止本次循环，进入下次循环
