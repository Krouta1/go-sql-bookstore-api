package models

import (
	"github.com/krouta1/go-sql-bookstore-api/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"column:name" json:"name"`
	Author      string `gorm:"column:author" json:"author"`
	Publication string `gorm:"column:publication" json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.First(&getBook, Id)
	return &getBook, db
}

func DeleteBook(Id int64) error {
	var book Book
	if err := db.Where("ID = ?", Id).First(&book).Error; err != nil {
		return err
	}
	return db.Delete(&book).Error
}
