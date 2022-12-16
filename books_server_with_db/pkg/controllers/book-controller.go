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

func CreateBook(w http.ResponseWriter, r *http.Request) {
	log.Info("handle [CreateBook]")
	NewBook := &models.Book{}
	utils.ParseBody(r, NewBook)
	createdBook := NewBook.CreateBook()

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
	log.Infof("handle [DeleteBook] : %v", params["bookID"])
	parsedId, parseErr := strconv.ParseInt(params["bookID"], 0, 0)
	if parseErr != nil {
		log.Error("get book by id, cant parse book ID")
	}
	removedBook := models.DeleteBook(parsedId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(removedBook)
}

// this is only for learning purpose -> very poor implementation
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	log.Info("handle [UpdateBook]")
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookID := vars["bookID"]
	id_, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		log.Error("error while parsing")
	}
	bookDetails, db := models.GetBookById(id_)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)

	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
