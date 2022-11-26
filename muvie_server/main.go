package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	pathhandlers "github.com/Anton6896/go_11_proj/muvie_server/pathHandlers"
	// "github.com/gorilla/mux"
	"github.com/gorilla/mux"
)

const (
	PORT = ":8000"
)

func main() {
	fmt.Println("new movie server")

	r := mux.NewRouter()
	r.HandleFunc("/movies", pathhandlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", pathhandlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies", pathhandlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", pathhandlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", pathhandlers.DeleteMovie).Methods("DELETE")

	log.Printf("starting server at : %v", PORT)
	log.Fatal(http.ListenAndServe(PORT, r))
}
