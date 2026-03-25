# GORM 实践题：图书管理系统

## 需求

实现一个简单的图书管理系统，使用 SQLite 数据库和 GORM。

## 数据模型

### Author（作者）
```go
type Author struct {
    ID   uint   // 主键
    Name string // 姓名，非空
    Bio  string // 简介
}
```

### Book（图书）
```go
type Book struct {
    ID       uint   // 主键
    Title    string // 书名，非空
    ISBN     string // ISBN号，唯一
    Price    float64 // 价格，默认0
    AuthorID uint   // 外键，关联Author
    Author   Author // 关联对象
}
```

## 功能要求

实现以下函数：

### 1. 初始化数据库
```go
func InitDB() (*gorm.DB, error)
```
- 连接 SQLite 数据库（文件名为 `library.db`）
- 自动迁移 Author 和 Book 表
- 返回 *gorm.DB

### 2. 添加作者
```go
func AddAuthor(db *gorm.DB, name, bio string) (*Author, error)
```
- 创建作者记录
- 返回创建的 Author

### 3. 添加图书
```go
func AddBook(db *gorm.DB, title, isbn string, price float64, authorID uint) (*Book, error)
```
- 创建图书记录
- 返回创建的 Book

### 4. 查询所有图书（带作者信息）
```go
func GetAllBooks(db *gorm.DB) ([]Book, error)
```
- 查询所有图书
- **使用 Preload 加载关联的作者信息**
- 返回图书列表

### 5. 根据作者查询图书
```go
func GetBooksByAuthor(db *gorm.DB, authorID uint) ([]Book, error)
```
- 查询指定作者的所有图书
- 返回图书列表

### 6. 更新图书价格
```go
func UpdateBookPrice(db *gorm.DB, bookID uint, newPrice float64) error
```
- 更新指定图书的价格

### 7. 删除图书
```go
func DeleteBook(db *gorm.DB, bookID uint) error
```
- 删除指定图书（软删除）

### 8. 事务：作者和图书一起添加
```go
func AddAuthorWithBooks(db *gorm.DB, authorName, authorBio string, books []struct{
    Title string
    ISBN  string
    Price float64
}) (*Author, error)
```
- 使用事务添加作者和多本图书
- 如果任何一步失败，全部回滚
- 返回创建的 Author

## 测试用例

在 `main()` 函数中完成以下测试：

1. 初始化数据库
2. 添加2个作者
3. 为每个作者添加2-3本图书
4. 查询所有图书并打印（显示书名和作者名）
5. 查询某个作者的所有图书并打印
6. 更新一本书的价格
7. 删除一本书
8. 再次查询所有图书，验证删除成功
9. 使用事务添加一个新作者和他的2本图书

## 运行

```bash
go mod tidy
go run main.go
```

## 预期输出示例

```
===== 所有图书 =====
《Go语言编程》- 张三 - ￥59.00
《Python入门》 - 李四 - ￥45.00
...

===== 张三的图书 =====
《Go语言编程》
...

===== 更新价格后 =====
...

===== 删除图书后 =====
...

===== 事务添加作者和图书 =====
作者：王五
图书：《Java编程》、《C++ primer》
```
