package main

import (
	"fmt"
	"log"
	"net/http"

	pathHandlers "github.com/Anton6896/go_11_proj/muvie_server/pathHandlers"
	"github.com/gorilla/mux"
)

const (
	PORT = ":8000"
)

func main() {
	fmt.Println("new movie server")

	r := mux.NewRouter()
	r.HandleFunc("/movies", pathHandlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", pathHandlers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", pathHandlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", pathHandlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", pathHandlers.DeleteMovie).Methods("DELETE")

	log.Printf("starting server at : %v", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
