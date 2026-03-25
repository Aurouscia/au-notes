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
- `size:100`：
- `not null`：
- `uniqueIndex`：
- `default:18`：

---

## 第2题：CRUD 操作

请写出完成以下操作的 GORM 代码：

1. 创建一条 User 记录（Name="张三", Age=25）：

2. 查询 ID=5 的用户：

3. 查询所有年龄大于 18 的用户：

4. 将 ID=5 的用户年龄改为 30：

5. 删除 ID=5 的用户（软删除）：

---

## 第3题：查询方法辨析

`First`、`Take`、`Last` 三个方法有什么区别？

- First：
- Take：
- Last：

---

## 第4题：Where 条件

以下两种查询方式有什么区别？什么情况下结果会不同？

```go
// 方式A：结构体条件
db.Where(&User{Name: "张三", Age: 0}).Find(&users)

// 方式B：Map 条件
db.Where(map[string]interface{}{"name": "张三", "age": 0}).Find(&users)
```

---

## 第5题：Dialector

`sqlite.Open("test.db")` 返回的是什么类型？它的作用是什么？

---

## 第6题：关联关系

请说明以下三种关联关系的区别，并各举一个生活中的例子：

1. Has One（一对一）：
2. Has Many（一对多）：
3. Many To Many（多对多）：

---

## 第7题：Preload

什么是 N+1 查询问题？Preload 如何解决它？

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
