package models

import (
	"github.com/i101dev/books-mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var bookToGet Book
	db := db.Where("ID=?", ID).Find(&bookToGet)
	return &bookToGet, db
}

func DeleteBook(ID int64) Book {
	var b Book
	db.Where("ID=?", ID).Delete(b)
	return b
}
