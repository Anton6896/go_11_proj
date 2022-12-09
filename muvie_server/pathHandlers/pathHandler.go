package pathhandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	tmpDB "github.com/Anton6896/go_11_proj/muvie_server/tmp_db"
	"github.com/gorilla/mux"
)

type ResponseMsg struct {
	Msg string `json:"msg"`
}

func GetMovies(w http.ResponseWriter, r *http.Request) {
	log.Println("handling /movies [GET]")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tmpDB.Movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf("handling /movies/%v [GET]\n", params["id"])
	w.Header().Set("Content-Type", "application/json")

	for _, movie := range tmpDB.Movies {
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

	for idx, movie := range tmpDB.Movies {
		if movie.ID == params["id"] {
			tmpDB.Movies = append(tmpDB.Movies[:idx], tmpDB.Movies[idx+1:]...)
			json.NewEncoder(w).Encode(ResponseMsg{Msg: fmt.Sprintf("data removed %v", params["id"])})
			return
		}
	}
	json.NewEncoder(w).Encode(ResponseMsg{Msg: fmt.Sprintf("cant find this id : %v", params["id"])})
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	log.Printf("handling /movies [POST]\n")
	w.Header().Set("Content-Type", "application/json")
	var movie tmpDB.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		json.NewEncoder(w).Encode(ResponseMsg{Msg: "please use valid Json data"})
		return
	}

	movie.ID = strconv.Itoa(rand.Intn(100000))
	tmpDB.Movies = append(tmpDB.Movies, movie)
	json.NewEncoder(w).Encode(ResponseMsg{Msg: fmt.Sprintf("movie created id : %v", movie.ID)})
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	log.Printf("updating /movies/%v [PUT]\n", params["id"])

	// remove existence movie
	for idx, mv := range tmpDB.Movies {
		if mv.ID == params["id"] {
			tmpDB.Movies = append(tmpDB.Movies[:idx], tmpDB.Movies[idx+1:]...)
		}
	}

	// add new movie
	var movie tmpDB.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		json.NewEncoder(w).Encode(ResponseMsg{Msg: "movie update failure"})
	}

	movie.ID = params["id"]
	tmpDB.Movies = append(tmpDB.Movies, movie)
	json.NewEncoder(w).Encode(ResponseMsg{Msg: fmt.Sprintf("movie updated id : %v", movie.ID)})
}
