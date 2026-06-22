package main

import (
	"fmt"
	"log"
	"net/http"
)

type PageData struct {
	Result string
	Text string
}

func main() {
	http.HandleFunc("/", HandlePage)
	http.HandleFunc("/ascii-art", HandleArt)
	fmt.Println("Server Running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
