# GORM 概念题

请阅读题目后直接写出答案。

---

## 第1题：模型定义

以下结构体定义了 GORM 模型，请说明每个标签的作用：

```go
type User struct {
    ID        uint      `gorm:"primaryKey"`
    Name      string    `gorm:"size:100;not null"`
    Email     string    `gorm:"uniqueIndex"`
    Age       int       `gorm:"default:18"`
    CreatedAt time.Time
}
```

- `primaryKey`：
    - 主键 ✅
- `size:100`：
    - 字符串长度 100 ✅
- `not null`：
    - 不能为 null ✅
- `uniqueIndex`：
    - 唯一索引 ✅
- `default:18`：
    - 默认值为 18 ✅

---

## 第2题：CRUD 操作

请写出完成以下操作的 GORM 代码：

1. 创建一条 User 记录（Name="张三", Age=25）：
    ```go
    db.Create(&User{Name:"张三", Age:25})
    ```
    **✅ 正确**

2. 查询 ID=5 的用户：
    ```go
    db.Where(&User{ID:5}).First(&user)
    ```
    **✅ 正确**，也可简写为 `db.First(&user, 5)`

3. 查询所有年龄大于 18 的用户：
    ```go
    db.Where("Age > ?", 18).Find(&users)
    ```
    **✅ 正确**，注意字段名用数据库列名（小写 `age`）或结构体字段名都可以

4. 将 ID=5 的用户年龄改为 30：
    ```go
    db.Where("ID = ?", 5).Update("Age", 30)
    ```
    **❌ 有误，正确写法：**
    ```go
    db.Model(&User{}).Where("id = ?", 5).Update("age", 30)
    // 或先查询再更新
    var user User
    db.First(&user, 5)
    db.Model(&user).Update("age", 30)
    ```
    **注意**：由于这个表达式看不出要操作的表，所以需要配合 Model 使用
    - Model 可以仅指定类型（新 struct 指针）
    - Model 也可以指定特定对象（已有 struct 指针）

5. 删除 ID=5 的用户（软删除）：
    ```go
    db.Where("ID = ?", 5).Delete()
    ```
    **❌ 有误，正确写法：**
    ```go
    db.Delete(&User{}, 5)
    // 或
    var user User
    db.First(&user, 5)
    db.Delete(&user)
    // 或
    db.Where("id = ?", 5).Delete(&User{})
    ```
    **注意**：Delete 需要传入模型指针或主键，不能空参数

---

## 第3题：查询方法辨析

`First`、`Take`、`Last` 三个方法有什么区别？

- First：
    - 按主键升序查询第一条记录 ✅
- Take：
    - 查询一条记录，不指定排序 **❌ 有误，不是"前几个"**
- Last：
    - 按主键降序查询最后一条记录 **❌ 有误，不是"后几个"**

    **补充说明**：
    - `First`：按主键升序排，取第一条（有默认排序）
    - `Take`：取一条，不保证顺序（性能最好）
    - `Last`：按主键降序排，取最后一条
    - 三者都是查**单条**记录，不是多条

---

## 第4题：Where 条件

以下两种查询方式有什么区别？什么情况下结果会不同？

```go
// 方式A：结构体条件
db.Where(&User{Name: "张三", Age: 0}).Find(&users)

// 方式B：Map 条件
db.Where(map[string]interface{}{"name": "张三", "age": 0}).Find(&users)
```

- 区别在于零值：结构体条件会忽略零值，Map 条件不会 ✅

---

## 第5题：Dialector

`sqlite.Open("test.db")` 返回的是什么类型？它的作用是什么？
- 返回的是`*sqlite.Dialector`，实现`gorm.Dialector`接口 ✅
- 作用是把数据库操作翻译为特定的 SQL 方言，这里是 sqlite ✅

---

## 第6题：关联关系

请说明以下三种关联关系的区别，并各举一个生活中的例子：

1. Has One（一对一）：
    - 一个实体绑定另一个实体 ✅
    - 例如：一个人只能拥有一张驾照，且一张驾照只能被一个人拥有 ✅
2. Has Many（一对多）：
    - 一个实体绑定多个实体 ✅
    - 例如：一个人可以拥有多辆车，但一辆车只能被一个人拥有 ✅
3. Many To Many（多对多）：
    - 多个实体绑定多个实体（需要中间表） ✅
    - 例如：一个人可以投资多只股票，一只股票也可以被多人投资 ✅

---

## 第7题：Preload

什么是 N+1 查询问题？Preload 如何解决它？

- 例如：“先查询用户，再查询用户所有的订单”这种问题，会需要进行多次查询
- Preload 使用 IN 表达式在一次查询中查出用户和订单

---

## 第8题：事务

以下两种事务写法有什么区别？各有什么优缺点？

```go
// 写法A：Transaction 块
db.Transaction(func(tx *gorm.DB) error {
    tx.Create(&user)
    tx.Create(&order)
    return nil
})

// 写法B：手动控制
tx := db.Begin()
tx.Create(&user)
tx.Create(&order)
tx.Commit()
```
- 第一种会捕获其中发生的异常自动回滚，但必须写在一个函数中完成开始和结束 ✅
- 第二种需要手动处理异常，但比较灵活，开始和结束可以分在不同函数中 ✅