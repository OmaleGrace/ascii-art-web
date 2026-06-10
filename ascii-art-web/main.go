package main

import (
	"net/http"
	"log"
	"fmt"
)

type PageData struct {
	Result string
}

func main() {
	http.HandleFunc("/", HandlePage)
	http.HandleFunc("/ascii-art", HandleArt)
	fmt.Println("Server Listening at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal()
	}
}