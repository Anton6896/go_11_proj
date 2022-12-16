package main

import (
	"fmt"
	"net/http"

	"github.com/Anton6896/go_11_proj/books_server_with_db/pkg/routes"
	"github.com/gorilla/mux"
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

const (
	PORT = ":9010"
)

func main() {
	log.Info(fmt.Sprintf("starting server at : 0.0.0.0%v", PORT))
	r := mux.NewRouter()
	routes.BookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, r))

}

/*
	https://www.youtube.com/watch?v=jFfo23yIWac&list=WL&index=3&ab_channel=freeCodeCamp.org 1:40

	! db for this exercise

	docker run --rm --name ant-mysql \
	-e MYSQL_USER=anton \
	-e MYSQL_PASSWORD=132123 \
	-e MYSQL_DATABASE=simplerest \
	-e MYSQL_ROOT_PASSWORD=123123 \
	-p 3306:3306 \
	-d mysql:latest
*/
