package pathhandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tmpdb "github.com/Anton6896/go_11_proj/muvie_server/tmp_db"
	"github.com/gorilla/mux"
)

type ResponceMsg struct {
	Msg string `json:"msg"`
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	log.Println("handling /movies [GET]")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tmpdb.Movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	log.Println("handling /movies [delete]")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for idx, movie := range tmpdb.Movies {
		if movie.ID == params["id"] {
			tmpdb.Movies = append(tmpdb.Movies[:idx], tmpdb.Movies[idx+1:]...)
			json.NewEncoder(w).Encode(ResponceMsg{Msg: fmt.Sprintf("data removed %v", params["id"])})
			return
		}
	}
	json.NewEncoder(w).Encode(ResponceMsg{Msg: fmt.Sprintf("cant find this id : %v", params["id"])})
}
