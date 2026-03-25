# GORM 基础

## 什么是 GORM？

GORM 是 Go 语言最流行的 ORM 库，全称为 **Go Object Relational Mapping**，它将数据库表映射为 Go 结构体，让你可以用面向对象的方式操作数据库。

## 安装

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/sqlite   # SQLite 驱动
go get -u gorm.io/driver/mysql    # MySQL 驱动
go get -u gorm.io/driver/postgres # PostgreSQL 驱动
```

## 核心概念

### 1. 连接数据库

```go
import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
)

func main() {
    // 连接 SQLite 数据库（文件形式）
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
}
```

**`sqlite.Open()` 详解**

`sqlite.Open("test.db")` 返回的是 `gorm.Dialector` 接口类型：

```go
// Dialector 接口定义了数据库驱动的统一规范
type Dialector interface {
    Name() string                          // 返回驱动名称，如 "sqlite"
    Initialize(*DB) error                  // 初始化数据库连接
    Migrator(db *DB) Migrator              // 返回迁移器，用于创建/修改表
    DataTypeOf(field *schema.Field) string // 将 Go 类型映射为数据库类型
    DefaultValueOf(field *schema.Field) clause.Expression
    BindVarTo(writer clause.Writer, stmt *Statement, v interface{})
    QuoteTo(writer clause.Writer, str string)  // SQL 标识符引用（如 `name` 或 "name"）
    Explain(sql string, vars ...interface{}) string
}
```

**作用**：Dialector 是 GORM 与具体数据库之间的**适配层**，负责：
- **SQL 方言转换**：不同数据库语法有差异（如 MySQL 用 `` ` `` 引用标识符，PostgreSQL 用 `"`）
- **数据类型映射**：Go 的 `string` 在 SQLite 中是 `TEXT`，在 MySQL 中是 `VARCHAR(255)`
- **连接管理**：建立和维护底层数据库连接

**连接流程**：

```go
// 第1步：创建 Dialector（数据库适配器）
dialector := sqlite.Open("test.db")  // 返回 gorm.Dialector

// 第2步：传入 GORM 创建 DB 实例
db, err := gorm.Open(dialector, &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),  // 可选：打印 SQL 日志
})
```

**常见驱动对比**：

| 数据库 | 导入路径 | 连接示例 |
|--------|----------|----------|
| SQLite | `gorm.io/driver/sqlite` | `sqlite.Open("test.db")` |
| MySQL | `gorm.io/driver/mysql` | `mysql.Open("user:pass@tcp(127.0.0.1:3306)/dbname")` |
| PostgreSQL | `gorm.io/driver/postgres` | `postgres.Open("host=localhost user=postgres password=xxx dbname=test")` |
| SQL Server | `gorm.io/driver/sqlserver` | `sqlserver.Open("sqlserver://user:pass@localhost:1433?database=test")` |

**DSN（数据源名称）格式**：

- **SQLite**：文件路径（如 `"test.db"`）或 `:memory:`（内存数据库）
- **MySQL**：`user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True`
- **PostgreSQL**：`host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable`

**获取底层数据库连接**：

```go
// 获取 *sql.DB 用于执行原生 SQL 或设置连接池
sqlDB, err := db.DB()
if err != nil {
    log.Fatal(err)
}

// 设置连接池参数
sqlDB.SetMaxIdleConns(10)    // 最大空闲连接数
sqlDB.SetMaxOpenConns(100)   // 最大打开连接数
sqlDB.SetConnMaxLifetime(time.Hour)  // 连接最大存活时间
```

- `sqlite.Open`返回`*sqlite.Dialector`（实现`gorm.Dialector`接口）
- dialect 意为“方言”，这个东西可用于转换 SQL 方言

### 2. 定义模型（Model）

模型是普通的 Go 结构体，通过标签定义数据库字段：`gorm:"xxx"`

```go
type User struct {
    ID        uint           `gorm:"primaryKey"`           // 主键
    Name      string         `gorm:"size:255;not null"`    // 字符串，非空
    Email     string         `gorm:"uniqueIndex"`          // 唯一索引
    Age       int            `gorm:"default:18"`           // 默认值
    CreatedAt time.Time      // 自动记录创建时间
    UpdatedAt time.Time      // 自动记录更新时间
    DeletedAt gorm.DeletedAt `gorm:"index"`                // 软删除标记
}
```

**常用标签**：
- `primaryKey` - 主键
- `autoIncrement` - 自增
- `size:n` - 字段长度
- `not null` - 非空
- `default:x` - 默认值
- `unique` / `uniqueIndex` - 唯一约束/索引
- `index` - 普通索引
- `column:xxx` - 指定列名
- `type:xxx` - 指定数据类型

### 3. 自动迁移（Auto Migration）

自动根据模型创建/更新数据库表结构：

```go
// 自动创建表（如果不存在）传入结构体的指针，提供类型信息
db.AutoMigrate(&User{})

// 同时迁移多个模型
db.AutoMigrate(&User{}, &Product{}, &Order{})
```

### 4. 增删改查（CRUD）

#### 创建（Create）

```go
// 创建单条记录
user := User{Name: "张三", Email: "zhangsan@example.com", Age: 25}
result := db.Create(&user) // 传入指针（因为需要修改自增后的 ID）

fmt.Println(user.ID)           // 返回插入数据的主键
fmt.Println(result.Error)      // 返回 error
fmt.Println(result.RowsAffected) // 返回插入记录的条数

// 批量创建
users := []User{
    {Name: "张三", Email: "zs@example.com"},
    {Name: "李四", Email: "ls@example.com"},
}
db.Create(&users)
```

#### 查询（Read）

- 取单条：`First` `Take` `Last` 传入结构体指针，让它使用查到的数据给结构体赋值
- 取多条：`Find` 传入切片指针，让它使用查到的数据给切片赋值（注意：无筛选条件的话会查出整个库）
- 筛选：`Where(占位符条件/结构体条件/map条件)`
    - 然后使用`First`或`Find`取数据
    - 注意结构体条件会忽略零值，map条件不会
- 投影：`Select("name", "age")`
    - 可以写 as：`Select("age, count(*) as count")`
- 排序：`Order("age desc, name asc")`
- 分页：
    - `Offset`（相当于 ef 的 Skip）
    - `Limit`（相当于 ef 的 Take）
- 更新：`Model` 选择特定实体（表）
    - `Update(key, value)` 更新单个字段
    - `Updates(结构体/map)` 更新多个字段
    - 如果是 Count 之类看不出表的查询，需要使用 Model 手动指定表
    - 如果以 Find(&users) 结尾（能看出是 users 表）则不需要
- 删除：`Delete`逻辑删除（需要 DeletedAt 字段）
    - `Unscoped().Delete()`物理删除

```go
// 根据主键查询
var user User
db.First(&user, 1)           // 查询 ID = 1 的记录
db.First(&user, "10")        // 查询 ID = 10 的记录
db.Take(&user)               // 获取第一条记录（不指定排序）
db.Last(&user)               // 获取最后一条记录

// 条件查询
var users []User
db.Where("name = ?", "张三").Find(&users)
db.Where("age > ?", 18).Find(&users)
db.Where("name IN ?", []string{"张三", "李四"}).Find(&users)

// 结构体条件（零值会被忽略）
db.Where(&User{Name: "张三", Age: 0}).Find(&users)

// Map 条件（零值有效）
db.Where(map[string]interface{}{"name": "张三", "age": 0}).Find(&users)

// 获取单条记录
var u User
db.Where("email = ?", "test@example.com").First(&u)

// 获取所有记录
db.Find(&users)

// 指定字段查询
db.Select("name", "age").Find(&users)
```

#### 更新（Update）

```go
// 更新单个字段
db.Model(&user).Update("name", "李四")

// 更新多个字段
db.Model(&user).Updates(User{Name: "李四", Age: 30})
db.Model(&user).Updates(map[string]interface{}{"name": "李四", "age": 30})

// 批量更新
db.Model(&User{}).Where("age > ?", 18).Update("status", "adult")
```

#### 删除（Delete）

```go
// 软删除（默认，需要 DeletedAt 字段）
db.Delete(&user)

// 永久删除
db.Unscoped().Delete(&user)

// 根据条件删除
db.Where("age < ?", 18).Delete(&User{})
```

### 5. 高级查询

```go
// 排序
db.Order("age desc").Find(&users)
db.Order("age desc, name asc").Find(&users)

// 分页
db.Limit(10).Offset(20).Find(&users)  // 第3页，每页10条

// 计数
var count int64
db.Model(&User{}).Where("age > ?", 18).Count(&count)

// 分组
type Result struct {
    Age   int
    Count int64
}
var results []Result
db.Model(&User{}).Select("age, count(*) as count").Group("age").Scan(&results)

// 原生 SQL
db.Raw("SELECT * FROM users WHERE age > ?", 18).Scan(&users)
db.Exec("UPDATE users SET age = age + 1 WHERE name = ?", "张三")
```

### 6. 关联关系

#### 一对一（Has One）

```go
type User struct {
    gorm.Model
    Name      string
    CreditCard CreditCard  // 一个用户有一张信用卡
}

type CreditCard struct {
    gorm.Model
    Number string
    UserID uint  // 外键
}
```

#### 一对多（Has Many）

```go
type User struct {
    gorm.Model
    Name  string
    Orders []Order  // 一个用户有多个订单
}

type Order struct {
    gorm.Model
    Product string
    UserID  uint  // 外键
}
```

#### 多对多（Many To Many）

```go
type User struct {
    gorm.Model
    Name    string
    Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
    gorm.Model
    Name string
}
```

#### 预加载（Preload）

预加载用于解决 **N+1 查询问题**，在查询主记录时**一次性**把关联记录也查出来。

**什么是 N+1 问题？**

```go
// ❌ 错误做法：产生 N+1 次查询
db.Find(&users)  // 第1次查询：SELECT * FROM users
for _, user := range users {
    db.Model(&user).Association("Orders").Find(&orders)  // N次查询
}
// 总共 1 + N 次查询，性能极差

// ✅ 正确做法：使用 Preload，只需 2 次查询
db.Preload("Orders").Find(&users)
// 第1次：SELECT * FROM users
// 第2次：SELECT * FROM orders WHERE user_id IN (1,2,3,...)
```

**基本用法**：

```go
// 查询所有用户，同时加载他们的订单
db.Preload("Orders").Find(&users)

// 条件预加载：只加载金额大于100的订单
db.Preload("Orders", "amount > ?", 100).Find(&users)

// 使用函数条件
db.Preload("Orders", func(db *gorm.DB) *gorm.DB {
    return db.Order("orders.amount DESC").Where("amount > ?", 100)
}).Find(&users)
```

**多级预加载（嵌套关联）**：

```go
// 用户 → 订单 → 商品
type User struct {
    gorm.Model
    Name   string
    Orders []Order
}

type Order struct {
    gorm.Model
    UserID    uint
    Products  []Product
}

type Product struct {
    gorm.Model
    OrderID uint
    Name    string
}

// 加载用户 → 他们的订单 → 订单里的商品
db.Preload("Orders.Products").Find(&users)
// 生成3条SQL：
// 1. SELECT * FROM users
// 2. SELECT * FROM orders WHERE user_id IN (...)
// 3. SELECT * FROM products WHERE order_id IN (...)
```

**预加载多个关联**：

```go
// 同时加载 Orders 和 CreditCard
db.Preload("Orders").Preload("CreditCard").Find(&users)

// 或使用 Preloads（效果相同）
db.Preloads("Orders").Preloads("CreditCard").Find(&users)
```

**Joins 预加载（一条 SQL 完成）**：

```go
// Preload 使用多条 SQL，Joins 使用 JOIN 一条 SQL
db.Joins("CreditCard").Find(&users)
// 生成：SELECT users.*, credit_cards.* FROM users 
//       LEFT JOIN credit_cards ON credit_cards.user_id = users.id

// 注意：Joins 适合一对一关系，一对多会产生重复数据
```

**Preload vs Joins 对比**：

| 特性 | Preload | Joins |
|------|---------|-------|
| SQL 条数 | 多条（1+N） | 1条 |
| 适用关系 | 一对一、一对多、多对多 | 主要一对一 |
| 条件筛选 | 可对关联表筛选 | 主表和关联表都可筛选 |
| 性能 | 数据量大时更好 | 数据量小时简单 |
| 重复数据 | 无 | 一对多时会产生 |

**预加载条件与主查询条件结合**：

```go
// 查询年龄>18的用户，并预加载他们的订单（只显示金额>50的）
db.Where("age > ?", 18).
   Preload("Orders", "amount > ?", 50).
   Find(&users)

// 注意：Preload 的条件只影响预加载的数据，不影响主查询
// 上面的例子：即使某用户的订单都<50，该用户仍会被查出（只是Orders为空）
```

**自定义预加载查询**：

```go
// 使用完整 Query 链
db.Preload("Orders", func(db *gorm.DB) *gorm.DB {
    return db.Select("*").
              Where("status = ?", "paid").
              Order("created_at DESC").
              Limit(5)  // 只预加载最近5个已支付订单
}).Find(&users)
```

### 7. 事务（Transaction）

```go
// 方式1：使用事务块
db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user).Error; err != nil {
        return err  // 返回错误会自动回滚
    }
    if err := tx.Create(&order).Error; err != nil {
        return err
    }
    return nil  // 返回 nil 提交事务
})

// 方式2：手动控制
tx := db.Begin()
if err := tx.Create(&user).Error; err != nil {
    tx.Rollback()
    return err
}
if err := tx.Create(&order).Error; err != nil {
    tx.Rollback()
    return err
}
tx.Commit()
```

## 常用配置选项

```go
db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),  // 打印 SQL 日志
})
```

## 总结

| 操作 | 方法 |
|------|------|
| 连接 | `gorm.Open(driver, config)` |
| 迁移 | `db.AutoMigrate(&Model{})` |
| 创建 | `db.Create(&model)` |
| 查询 | `db.First/Last/Take/Find(&model, conditions)` |
| 更新 | `db.Model(&model).Update/Updates(values)` |
| 删除 | `db.Delete(&model)` |
| 条件 | `db.Where(conditions).Action()` |
| 排序 | `db.Order(field).Find()` |
| 分页 | `db.Limit(n).Offset(m).Find()` |
| 事务 | `db.Transaction(func(tx) error)` |
