package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Author 作者模型
type Author struct {
	ID   uint
	Name string
	Bio  string
}

// Book 图书模型
type Book struct {
	ID       uint
	Title    string
	ISBN     string
	Price    float64
	AuthorID uint
	Author   Author
}

// InitDB 初始化数据库连接
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"))
	// ❌ 应立即检查 err，不要拖到最后再返回
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Author{}, &Book{})
	return db, nil
}

// AddAuthor 添加作者
func AddAuthor(db *gorm.DB, name, bio string) (*Author, error) {
	author := &Author{
		Name: name,
		Bio:  bio,
	}
	// ❌ 任何操作返回的都是 db 本身（可以链式调用），如果出错，Error 也在里面，应在最后通过 .Error 检查
	if err := db.Create(author).Error; err != nil {
		return nil, err
	}
	return author, nil
}

// AddBook 添加图书
func AddBook(db *gorm.DB, title, isbn string, price float64, authorID uint) (*Book, error) {
	book := &Book{
		Title:    title,
		ISBN:     isbn,
		Price:    price,
		AuthorID: authorID,
	}
	if err := db.Create(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

// GetAllBooks 查询所有图书（带作者信息）
func GetAllBooks(db *gorm.DB) ([]Book, error) {
	var books []Book
	if err := db.Preload("Author").Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetBooksByAuthor 根据作者查询图书
func GetBooksByAuthor(db *gorm.DB, authorID uint) ([]Book, error) {
	var results []Book
	if err := db.Where("author_id = ?", authorID).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

// UpdateBookPrice 更新图书价格
func UpdateBookPrice(db *gorm.DB, bookID uint, newPrice float64) error {
	return db.Model(&Book{}).Where("id = ?", bookID).Update("price", newPrice).Error
}

// DeleteBook 删除图书（软删除）
func DeleteBook(db *gorm.DB, bookID uint) error {
	return db.Delete(&Book{}, bookID).Error
}

// AddAuthorWithBooks 事务：添加作者和图书
func AddAuthorWithBooks(db *gorm.DB, authorName, authorBio string, books []struct {
	Title string
	ISBN  string
	Price float64
}) (*Author, error) {
	var resAuthor *Author
	err := db.Transaction(func(tx *gorm.DB) error {
		a := &Author{Name: authorName, Bio: authorBio}
		// ❌ 事务内必须用 tx 对象操作（tx 是 transaction 的缩写）
		if err := tx.Create(a).Error; err != nil {
			return err
		}
		resAuthor = a

		newBooks := make([]Book, len(books))
		for idx, b := range books {
			newBooks[idx] = Book{
				Title:    b.Title,
				ISBN:     b.ISBN,
				Price:    b.Price,
				AuthorID: a.ID,
			}
		}
		if err := tx.CreateInBatches(&newBooks, 100).Error; err != nil {
			return err
		}
		return nil
	})
	return resAuthor, err
}

func main() {
	fmt.Println("图书管理系统")

	db, err := InitDB()
	if err != nil {
		fmt.Println("初始化数据库出错")
		fmt.Println(err)
	}

	a1, _ := AddAuthor(db, "author 1", "bio 1")
	a2, _ := AddAuthor(db, "author 2", "bio 2")
	b1, _ := AddBook(db, "book1", "1234567", 12.34, a1.ID)
	b2, _ := AddBook(db, "book2", "7654321", 23.45, a1.ID)
	AddBook(db, "book3", "abcdefg", 34.56, a1.ID)
	AddBook(db, "book4", "qwerty", 45.67, a2.ID)
	AddBook(db, "book5", "zxcvbn", 56.78, a2.ID)
	_ = b1
	_ = b2

	books, _ := GetAllBooks(db)
	fmt.Println("===所有图书===")
	for _, b := range books {
		fmt.Println(b.Title, b.Author.Name)
	}

	fmt.Println("===a1的所有图书===")
	a1sBooks, _ := GetBooksByAuthor(db, a1.ID)
	for _, b := range a1sBooks {
		fmt.Println(b.Title)
	}

	UpdateBookPrice(db, b1.ID, 0.01)
	DeleteBook(db, b2.ID)

	booksAfterOp, _ := GetAllBooks(db)
	fmt.Println("===操作后的所有图书===")
	for _, b := range booksAfterOp {
		fmt.Println(b.Title, b.Author.Name)
	}

	_, _ = AddAuthorWithBooks(db, "author 3", "bio 3", []struct {
		Title string
		ISBN  string
		Price float64
	}{
		{Title: "book6", ISBN: "poiuyt", Price: 114.514},
		{Title: "book7", ISBN: "lkjhgf", Price: 1919.810},
	})

	booksAfterTrans, _ := GetAllBooks(db)
	fmt.Println("===事务后的所有图书===")
	for _, b := range booksAfterTrans {
		fmt.Println(b.Title, b.Author.Name)
	}
}
