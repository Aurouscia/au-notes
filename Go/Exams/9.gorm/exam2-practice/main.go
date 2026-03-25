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
	db.AutoMigrate(&Author{}, &Book{})
	return db, err
}

// AddAuthor 添加作者
func AddAuthor(db *gorm.DB, name, bio string) (*Author, error) {
	author := &Author{
		Name: name,
		Bio:  bio,
	}
	db.Create(author)
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
	db.Create(book)
	return book, nil
}

// GetAllBooks 查询所有图书（带作者信息）
func GetAllBooks(db *gorm.DB) ([]Book, error) {
	authors := []Book{}
	db.Preload("Author").Find(&authors)
	return authors, nil
}

// GetBooksByAuthor 根据作者查询图书
func GetBooksByAuthor(db *gorm.DB, authorID uint) ([]Book, error) {
	results := []Book{}
	db.Where(&Book{AuthorID: authorID}).Find(&results)
	return results, nil
}

// UpdateBookPrice 更新图书价格
func UpdateBookPrice(db *gorm.DB, bookID uint, newPrice float64) error {
	db.Model(&Book{}).Where("ID = ?", bookID).Update("Price", newPrice)
	return nil
}

// DeleteBook 删除图书（软删除）
func DeleteBook(db *gorm.DB, bookID uint) error {
	db.Model(&Book{}).Delete(bookID)
	return nil
}

// AddAuthorWithBooks 事务：添加作者和图书
func AddAuthorWithBooks(db *gorm.DB, authorName, authorBio string, books []struct {
	Title string
	ISBN  string
	Price float64
}) (*Author, error) {
	var resAuthor *Author
	db.Transaction(func(tx *gorm.DB) error {
		a, err := AddAuthor(db, authorName, authorBio)
		resAuthor = a
		if err != nil {
			return err
		}
		newBooks := make([]Book, len(books))
		for idx, b := range books {
			newBooks[idx] = Book{
				Title:    b.Title,
				ISBN:     b.ISBN,
				Price:    b.Price,
				AuthorID: a.ID, // 使用刚刚插入的作者ID
			}
		}
		db.CreateInBatches(&newBooks, 100)
		return nil
	})
	return resAuthor, nil
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

	AddAuthorWithBooks(db, "author 3", "bio 3", []struct {
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
