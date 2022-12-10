package controllers

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)

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
