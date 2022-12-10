package main

import (
	"net/http"

	pathHandlers "github.com/Anton6896/go_11_proj/muvie_server/pathHandlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	PORT = ":8000"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("new movie server")

	r := mux.NewRouter()
	r.HandleFunc("/movies", pathHandlers.GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", pathHandlers.GetMovie).Methods("GET")
	r.HandleFunc("/movies", pathHandlers.CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", pathHandlers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", pathHandlers.DeleteMovie).Methods("DELETE")

	log.Info("starting server at : %v", PORT)
	log.Error(http.ListenAndServe(PORT, r))
}
