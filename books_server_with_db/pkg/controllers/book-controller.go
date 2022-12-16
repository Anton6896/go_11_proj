package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Anton6896/go_11_proj/books_server_with_db/pkg/models"
	"github.com/Anton6896/go_11_proj/books_server_with_db/pkg/utils"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var book models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Info("handle [CreateBook]")
	newBook := &models.Book{} // address for new book
	utils.ParseBody(r, newBook)
	createdBook := models.CreateBook(newBook)

	// res, err := json.Marshal(createdBook)
	// if err != nil {
	// 	log.Error("cant parse json while creating new book")
	// }
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdBook) // also way to send response

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	log.Info("handle [GetBook]")
	books := models.GetAllBooks()
	res, err := json.Marshal(books)
	if err != nil {
		log.Error(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Infof("handle [GetBookById] : %v\n", params["bookID"])
	parsedId, parseErr := strconv.ParseInt(params["bookID"], 0, 0)
	if parseErr != nil {
		log.Error("get book by id, cant parse book ID")
	}
	bookDetail, _ := models.GetBookById(parsedId)
	res, jsonErr := json.Marshal(bookDetail)
	if jsonErr != nil {
		log.Errorf("cant parse json data fro book : %v/n", params["bookID"])
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Infof("handle [DeleteBook] : %v\n", params["bookID"])
	parsedId, parseErr := strconv.ParseInt(params["bookID"], 0, 0)
	if parseErr != nil {
		log.Error("get book by id, cant parse book ID")
	}
	removedBook := models.DeleteBook(parsedId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(removedBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Info("handle [UpdateBook]")
}
