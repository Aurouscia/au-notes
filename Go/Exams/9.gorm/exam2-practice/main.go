package main

import (
	"fmt"

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
	// TODO: 实现数据库连接和自动迁移
	return nil, nil
}

// AddAuthor 添加作者
func AddAuthor(db *gorm.DB, name, bio string) (*Author, error) {
	// TODO: 实现添加作者
	return nil, nil
}

// AddBook 添加图书
func AddBook(db *gorm.DB, title, isbn string, price float64, authorID uint) (*Book, error) {
	// TODO: 实现添加图书
	return nil, nil
}

// GetAllBooks 查询所有图书（带作者信息）
func GetAllBooks(db *gorm.DB) ([]Book, error) {
	// TODO: 使用 Preload 查询所有图书
	return nil, nil
}

// GetBooksByAuthor 根据作者查询图书
func GetBooksByAuthor(db *gorm.DB, authorID uint) ([]Book, error) {
	// TODO: 查询指定作者的所有图书
	return nil, nil
}

// UpdateBookPrice 更新图书价格
func UpdateBookPrice(db *gorm.DB, bookID uint, newPrice float64) error {
	// TODO: 更新图书价格
	return nil
}

// DeleteBook 删除图书（软删除）
func DeleteBook(db *gorm.DB, bookID uint) error {
	// TODO: 删除图书
	return nil
}

// AddAuthorWithBooks 事务：添加作者和图书
func AddAuthorWithBooks(db *gorm.DB, authorName, authorBio string, books []struct {
	Title string
	ISBN  string
	Price float64
}) (*Author, error) {
	// TODO: 使用事务添加作者和图书
	return nil, nil
}

func main() {
	// TODO: 完成测试用例
	fmt.Println("图书管理系统")
}
