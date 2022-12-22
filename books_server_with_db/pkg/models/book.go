package models

import (
	"github.com/Anton6896/go_11_proj/books_server_with_db/pkg/config"
	"github.com/jinzhu/gorm"

	log "github.com/sirupsen/logrus"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDb()

	log.Info("first migration")
	db.AutoMigrate(&Book{}) // smart migration
}

func (b *Book) CreateBook() *Book { // add function createBook() to this model object
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id_ int64) (*Book, *gorm.DB) {
	var b Book
	db := db.Where("ID=?", id_).Find(&b)
	return &b, db
}

func DeleteBook(id_ int64) *Book {
	var b Book
	db.Where("ID=?", id_).Delete(b)
	return &b
}
