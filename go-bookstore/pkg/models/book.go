package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lintopaul/go-projects/go-bookstore/pkg/config"
)

var db *gorm.DB

// Tags are optional to use when declaring models. Refer the docs GORM supported tags:-
// https://gorm.io/docs/models.html
type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"default:My Book"`
	Author      string `json:"author" gorm:"default:Linto Paul"`
	Publication string `json:"publication" gorm:"default:Penguin"` 
	Cost		int	   `json:"cost,omitempty"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	// db.Debug().Create(&b) can be used for debugging
	db.Create(&b) 
	
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)

	return Books
}

func GetBookByID(ID int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", ID).Find(&getBook)

	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)

	return book
}
