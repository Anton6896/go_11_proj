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
	db.AutoMigrate(&Book{}) // create empty book
}
