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

func main() {
	fmt.Println("new movie server")

	router := mux.NewRouter()
	router.HandleFunc("/movies", pathhandlers.GetMovies).Methods("GET")
	router.HandleFunc("/movie/{id}", pathhandlers.GetMovies).Methods("GET")
}
