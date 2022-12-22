package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anton6896/go_11_proj/go_server/handlers"
)

const (
	PORT = ":8080"
)

func main() {
	fmt.Println("new go server")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", handlers.FormHandler)
	http.HandleFunc("/hello", handlers.HelloHandler)

	log.Printf("running server at port : %v\n", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}
