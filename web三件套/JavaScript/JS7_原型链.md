# JavaScript 原型链

构造对象时，使用 `__proto__`（前后各两个下划线）来指定其原型。

```javascript
objA = {
    id: 1,
    name: "David",
    __proto__: {  // 这种构造写法依然支持
        a: 7,
        b: 8
    }
};
console.log(objA.__proto__);  // 这种访问写法已经废弃
console.log(objA.a);  // 访问某对象的属性时，自己的找不到，就在原型里找
```

如果对象本身就有 `a`，那么原型里的 `a` 就会被忽略，称为 **property shadowing**（属性遮蔽）。

```javascript
child = {
    value: 3,
    __proto__: {
        value: 5,
        fun() {
            return this.value;
        }
    }
};
console.log(child.fun());  // 3
```

在子级中调用父级的方法，其中的 `this` 是指向子级的。
