# ES概念
## ECMAScript（ECMA-262）与JS的关系
前者是后者的规范，后者是前者的实现  
JS由网景公司开发，提交给标准化组织ECMA成为国际标准  
## JS由三部分组成
- ECMAScript（核心语言功能）
- 文档对象模型DOM：访问和操作网页文档，把网页映射为多层节点结构，网页的每个部分都对应着某种类型的节点
- 浏览器对象模型BOM：访问和操作浏览器窗口
## ES6
也叫ES2015（某个大更新）
ES6泛指ES5.1之后的所有版本

# let
## 作用域
与会变量提升的var不同，let仅在它所在的代码块有效  
“块级作用域”是ES6新增的，ES5没有，ES5只有全局作用域和函数作用域  
好处：避免块（例如循环）里面的变量污染外面的变量  
典型例子：for循环+闭包引用var i变量（见`1.let.js`）造成不直观的输出  
注意：**for循环的小括号和大括号是两个不同的块级作用域**，可以let两个同名变量  
（但是引擎做了特殊处理使得大括号内可以访问小括号的let变量）  
let与块级作用域的概念紧密相关
## 提升
var和function会提升，**仅会被函数作用域限制，不会被块级作用域限制！**  
所以块内层的声明（不管是否执行到）都可能覆盖外层的声明（见`1.let.js`）  
（注意变量查找中的“截胡”）  
这是一种奇怪的现象（所以用改进版的let取代它）
- 在var a前使用a：a为undefined，不会报错
- 在let a前使用a：a未定义，报错
## 重复
var可以重复声明，let不可以（排他性）  
以下情况都会报错SyntaxError：
- 先let后let
- 先let后var
- 先var后let
- let**函数参数同名的变量**（但是不在函数顶级作用域内（套在块内）的话，不会冲突）  
不直观的行为：for循环变量用var i声明的话，每个循环内的i都是同一个（可以通过函数数组的闭包看出）
## 暂时性死区
只要块级作用域中存在let，它声明的变量就会绑定这个作用域，**取用该变量名时截胡，不再能访问到外面的同名变量**  
如果先取用再let就会造成：取用时截胡，但“没有定义”（位于暂时性死区*TDZ*内）抛出错误ReferenceError  
暂时性死区*TDZ*：块顶部到let的前一行  
### typeof
暂时性死区中，typeof 也不安全：也会抛出ReferenceError  
离谱：typeof 一个没有任何let的变量反而不会报错，而是得到undefined，但**typeof之外的取用还是会报错的！**
### 隐蔽死区情况
函数参数默认值中`x=y`的情况，如果`x=y`在`y=2`前面  
那就会抛出ReferenceError，因为**参数初始化的行为类似let**  
也不会提升，也有其前面的暂时性死区  
但如果把`y=2`放到`x=y`前面，问题就会消失

# const
const必须立即初始化，不能再次赋值，其他的行为和let完全一致，**也有受块级作用域限制，也不会提升，也有暂时性死区**
## mutate
废话：对于引用类型只能保证地址不会变，不能保证里面不Mutate  
如果需要禁止mutate，应该使用Object.freeze()：严格模式下报错，非严格模式下“不起作用”

# 解构赋值（Destructuring）
## 数组解构
按一定的模式从数组和对象中提取值，赋给变量，称为解构赋值  
这种写法属于“模式匹配”，两边结构一样哪怕嵌套很深也能解构
- `let [a, b, c] = [1, 2, 3]`
- `let [d, [[e], f], g] = [4, [[5], 6], 7]`
- `let [first, ...rest] = [1,2,3,4,5,6,7]`（此时rest为数组）
### 左侧比右侧多
如果解构不成功，变量的值为undefined或空数组（...的情况）
### 左侧比右侧少
不会有问题
### 数组解构目标必须iterable（实现iterator接口）  
但是如果试图解构一个不可遍历的东西例如数字，会抛出错误  
`[a, [b]] = [1, 2]`会抛出错误
有这些内置类型iterable：String、Array、TypedArray（如 Int8Array, Uint32Array 等）、Map、Set等
注：**遍历Map时得到的是键值对[key, value]**
### 解构可以有默认值（注意生效条件）
`[a = 2, b = 3] = [1]`
- 如果解构的值**严格等于undefined**，变量的默认值生效（null时不会生效）
- 如果解构的值不严格等于undefined，变量的默认值不生效
#### 刁钻情况
- 默认值是懒求值的（不用到默认值的时候，不会计算默认值表达式）
- 默认值可以用解构里的其他变量（但必须讲究顺序，`[x=y, y=1] = []`会抛出“初始化前使用”ReferenceError

## 对象解构
不同：数组的解构是按顺序对应的，对象的键值对没有顺序概念，是靠键对应的（顺序可随便调整）  
注意：解构赋值的大括号：左边是解构出的，右边是赋值目标`const obj = {m: 2}; let {n:n1} = obj`（如果新旧名字一样即可无需冒号简写）  
试图解构原对象中不存在的键：得到undefined  
注：函数也可以解构出来用  
也可以嵌套解构，也是模式匹配（函数解构和对象解构可以混合嵌套）
### 模式不可取用
`const {p:{arr:arr1, name:name1}} = obj2`中，p只是模式名字，不是实际存在的对象，取用会报错（未定义）
如果同时需要p和p内的属性，应该写成：`const {p:p1, p:{arr:arr2, name:name2}} = obj2`
### 默认值
生效条件与数组一样（想要的值===undefined）
`const {x:msg='Hello'} = {}`
### 坑：无法区分解构赋值与代码块
对于一个已经let的变量，往里解构不能写在行首
```js
let x = 1
{x} = {x: 1} //{x}会被理解为一个代码块，造成错误 SyntaxError: Unexpected token '='
({x} = {x: 1}); //用小括号包起来就不会有问题
({} = {x: 1}) //注意两个小括号即使隔了行也会被识别为函数调用，必须加分号
```
### 没用但没错的写法
解构等号左侧可以是[]或{}（不会有问题但也没用）  
### 数组也是对象
数组也是对象，可以使用数字索引当键（索引是数字或字符串都行），用对象解构语法解构数组  
用这种方法可以精确提取某些位置的值
`const { 1:a } = [1, 2, 3]` 得到a为2（索引为1）
### 字符串也是iterable
可以用数组解构得到每个字符，也可以通过对象解构得到其length属性
### 转对象
如果解构赋值右边不是对象，则会转为对象（例如可以提取出number的原型链上的方法）
`let {toString: s,valueOf: v} = 123`
### 函数参数
函数参数的解构也可以使用默认值。
```js
function move({x = 0, y = 0}) { //注意默认值
	return [x, y];
}
console.log(move({x: 3, y: 8})); // [3, 8]

// 解构赋值可以方便地将一组参数与变量名对应起来
// 参数是一组有序的值
function f([x, y, z]) { ... }
f([1, 2, 3]);
// 参数是一组无序的值
function f({x, y, z}) { ... }
f({z: 3, y: 2, x: 1});
```


# 模板字符串
使用反引号：内部可随意使用反引号之外的任何字符，会保留换行和缩进（可使用trim去掉前后的换行）  
可解决可读性差的字符串拼接和转义字符的问题  
仅有反引号本身需要反斜杠转义
```js
`abc${expression}def`
```
expression可以是任意js表达式（包括函数调用），但不能是多个语句（不能有分号）！

# 函数参数默认值
传统做法：p = p || 'someVal'，容易出问题，因为||仅判断falsy，不是严格undefined  
函数参数默认值也是es6新加的东西  
注意：触发条件与解构赋值的默认值相同（严格等于undefined）  
参数默认值也是惰性求值的
## 坑
若参数位置是对象解构，但没有对应的参数  
这种情况会报错TypeError: Cannot destructure  
应该传入一个空对象`{}`或使用默认值`{a, b=2} = {}`兜住
## 位置
有默认值的参数应该在尾部，不在尾部也不会有问题，但无法省略那个位置（造成错位）只能显式输入undefined
```js
function func3(x = 1, y){
    console.log(x, y)
}
func3(2) //其实传到x了
func3(, 2) //只有数组能用这种光逗号的省略写法，参数列表不行！
func3(undefined, 2)
```

# rest参数
在ES5中，函数内可以使用arguments（疑似保留字）数组访问所有参数  
ES6中，可以在参数列表末尾写`...rest`或`...others`来获取前面参数列表之外剩下的参数（是一个数组）
## 冷知识
函数也有length属性，表示其参数个数（不包括rest参数，`a, ...b`会算作1）

# 扩展运算符
同样是三个点，出现在形参时叫`rest参数`，出现在实参时叫`扩展运算符`  
可以把一个数组的元素当几个值分别填入函数参数`const res = func(first, ...arr)`  
任何Iterator接口的对象，都可以用扩展运算符转为真正数组`[...arr]`  
例如：`const arr = [...'hello']`
## rest参数必须在末尾
```js
const [...butLast, last] = [1, 2, 3, 4, 5]; //错误语法
const [first, ...middle, last] = [1, 2, 3, 4, 5] //错误语法
```

# 箭头函数
- 如果箭头函数的代码块不仅只有return语句，就要使用语句块将它们括起来。
- 如果箭头函数返回一个对象，必须在对象外面加上括号和return。  
`(a, b)=>a+b`  
`(a, b)=>{return{a+b}}`  
`(a, b)=>{a, b}` 不能直接返回对象！会被认作是代码块

- 函数体内的`this`对象，就是定义时所在的对象，而不是使用时所在的对象。
- 不可以当作构造函数，也就是说，不可以使用new命令，否则会抛出一个错误。
- 不可以使用arguments对象，该对象在函数体内不存在。如果要用，可以用rest参数代替

# 迭代器Iterator
内部：不断用next()获取下一个元素，直到最后一个元素  
一般被`for of`使用  
以下代码相当于直接`for v of arr`
```js
const arr = ['red', 'green', 'blue'];
let iterator = arr[Symbol.iterator](); //得到“返回其iterator的一个函数”
for(let v of iterator) {
    console.log(v);
}
```
对象的键值对没有先后顺序，所以没有`iterator`，**不能for of一个非Iterable的对象**
## for of 与 for in 遍历数组的区别
`for in`遍历的是键（但可以应用于所有对象）  
`for of`只会遍历到数字索引的属性（不包括length），且得到的是值

# Promise
## 特点
- 状态不受外界影响，只由其执行结果决定
- 有pending（初始）、resolved、rejected三种状态
- 一旦离开pending状态变为resolved或rejected，**状态就固定不能再变**
- 与“错过就没了”的事件不同，即使promise已经结束，再添加回调函数也会得到其结果
## 等价关系
- 在promise中throw error相当于reject(error)，要么被try-catch/链后面的catch()捕获，要么向上传递到运行环境
- p.catch(e)相当于then(xxx, e)，catch是then的第二参数语法糖
- 回调地狱（多个有先后顺序依赖的异步操作）等价于多个then链式调用
## 多个promise的聚合
```js
// Promise.all：把promise数组聚合为一个“全部完成后才完成”的promise（返回值为promise数组对应的结果数组）
const p3 = Promise.all([p1, p2]).then(console.log)
// Promise.race：把promise数组聚合为一个“任意一个完成后就完成”的promise（返回值为首个完成的promise的结果值）
const p4 = Promise.race([p1, p2]).then(console.log)
```

# 模块
node的commonJS模块：`const fs = require('fs')`  
- 把指定名称的文件内的导出作为一个对象返回
- 运行时加载，无法做编译时静态优化
## export
export规定的是模块外部**访问它的接口**，不是值  
export时应该让接口能与变量**建立一一对应的关系**  
所以：**export后必须是声明或大括号**
```js
export 1 //错误：不能export一个值

let a = 1
export a //错误：不能export一个值

let obj = {}
export obj //错误：不能export一个值

function f1(){}
export f1 //错误：不能export一个值

export let a = 1 //OK
export {a, obj} //OK
export {obj as obj1} //OK
export {f1} //OK
export function f1(){ } //OK
```
## import
import的名必须与export的一致（export default除外）  
但是可以起别名（import { a as b } from 'xxx'）
### 单向动态绑定关系
import导入的let变量是只读的（试图替换会报错） 
因为import变量本质上是个输入接口（单向） 
注：只是不允许替换而已，还是可以mutate的
目标模块中如果替换了值，import到的变量值也会立即改变（动态绑定机制，没有缓存）
### 静态
import和export的大括号里不能有表达式，只能是标识符和as  
必须作为顶级语句，不能塞到循环、条件和函数里
### 多次import同一模块
import是singleton的
```js
import { a } from 'xyz'
import { b } from 'xyz'
//与
import { a, b } from 'xyz'
//完全一致（在静态分析中合并了）
```
### 默认导入
对于`export default function(){ ... }`（默认导出）注意：默认导出可以导出匿名函数，不一定要给这个函数起名字
则import时无需知道变量名，可自己起名字`import someFunc from 'xxx'`  
也可以把某个变量as default指定成默认导出`export { a as default }`
默认导入和具名导入可以混合使用：`import _, { forEach } from 'lodash'`
### 导入后立即导出
`export { xxx } from 'xxx'`  
`export * from 'xxx'`  
> 注：公司不允许这种写法，必须先写一行import，再写一行export
### 通配符导入
`import * as sth from 'xxx'`一次性导入所有export并将其整理为一个命名空间对象  
不能as后直接解构，错误写法：`import * as { foo } from 'mod.js'`，应该as起名后再另起一行解构  
可以通过`sth.default`访问到默认导出  
> 注：公司不允许通配符导入
> 注：必须有as

# 新特性
## Array.prototype.includes()  
检查某个东西在不在数组里，注：使用的是三等号判断（值必须类型相同，引用必须地址相同）  
第二参数：从什么index开始找（如果超出数组范围不会报错，会返回false）（**和at一样可以负数表示倒数**）  
> 注：`NaN == NaN`和`NaN === NaN`都返回false  
> 如果需要指定比较逻辑，应该使用（可以接收比较函数）的`some`方法
## Array.prototype.at()
作用和`[]`索引一样，超出范围会返回undefined  
支持负数：`arr.at(-1)`指的是arr最后一个对象
## Object.entries()
静态方法，不能通过object实例调用    
可以获取一个对象的所有键值对（以`[[key, value], [key, value], ...]`双层数组的形式）  
可以结合数组解构：`for(let [k, v] of Object.entries(obj))`
注意：
- 不包括原型链上的属性
- 值为undefined的属性也会包括在里面
- 不会处理Symbol类型的属性
## Object.keys()
可以获取一个对象的所有键`[key, key, ...]`  
行为与entries相同  
与`for-in`的区别：for-in会包括原型链上的键，Object.keys()不会
## Object.values()
同上