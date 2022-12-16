package controllers

import (
	"net/http"

	"github.com/Anton6896/go_11_proj/books_server_with_db/pkg/models"
	log "github.com/sirupsen/logrus"
)

var book models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Info("handle [CreateBook]")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	log.Info("handle [GetBook]")
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	log.Info("handle [GetBookById]")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Info("handle [UpdateBook]")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	log.Info("handle [DeleteBook]")
}
