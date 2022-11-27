package pathhandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	tmpdb "github.com/Anton6896/go_11_proj/muvie_server/tmp_db"
	"github.com/gorilla/mux"
)

type ResponseMsg struct {
	Msg string `json:"msg"`
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	log.Println("handling /movies [GET]")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tmpdb.Movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf("handling /movies/%v [GET]\n", params["id"])
	w.Header().Set("Content-Type", "application/json")

	for _, movie := range tmpdb.Movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	json.NewEncoder(w).Encode(ResponseMsg{Msg: fmt.Sprintf("cant find this movie : %v", params["id"])})
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf("handling /movies/%v [DELETE]\n", params["id"])
	w.Header().Set("Content-Type", "application/json")

	for idx, movie := range tmpdb.Movies {
		if movie.ID == params["id"] {
			tmpdb.Movies = append(tmpdb.Movies[:idx], tmpdb.Movies[idx+1:]...)
			json.NewEncoder(w).Encode(ResponseMsg{Msg: fmt.Sprintf("data removed %v", params["id"])})
			return
		}
	}
	json.NewEncoder(w).Encode(ResponseMsg{Msg: fmt.Sprintf("cant find this id : %v", params["id"])})
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling /movies [POST]\n")
	w.Header().Set("Content-Type", "application/json")
	var movie tmpdb.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		json.NewEncoder(w).Encode(ResponseMsg{Msg: "please use valid Json data"})	
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(100000))
	tmpdb.Movies = append(tmpdb.Movies, movie)
	json.NewEncoder(w).Encode(ResponseMsg{Msg: fmt.Sprintf("movie created id : %v", movie.ID)})
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {

}
