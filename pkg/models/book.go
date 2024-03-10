package models

import (
    "github.com/jinzhu/gorm"
    "github.com/monster0freason/golang-bookstore-management-system/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
    Name        string `gorm:"" json:"name"`
    Author      string ` json:"author"`
    Publication string ` json:"publication"`
}

func init() {
    config.Connect()
    db = config.GetDB()
    db.AutoMigrate(&Book{})
}


// CreateBook creates a new book in the database
func (b *Book)CreateBook() *Book {
    db.NewRecord(b)
    db.Create(&b)
    return b
}

// GetAllBooks retrieves all books from the database
func GetAllBooks() []*Book {
    var Books []*Book
    db.Find(&Books)
    return Books
}

// GetBookByID retrieves a book from the database by its ID
func GetBookByID(Id int64) (*Book, *gorm.DB) {
    var getBook Book
    db := db.Where("ID = ?", Id).Find(&getBook)
    return &getBook, db
}

// DeleteBook deletes a book from the database by its ID
func DeleteBook(ID int64) Book {
    var book Book
    db.Where("ID = ?", ID).Delete(book)
    return book
}