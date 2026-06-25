package main

import (
	"fmt"
	"log"
	"net/http"
)

type Pagedata struct {
	Result string
	Text   string
	Banner string
}

func main() {
	http.HandleFunc("/", HandlePage)
	http.HandleFunc("/ascii-art", HandleArt)
	http.HandleFunc("/ascii-art-switch", HandleSwitch)
	fmt.Println("Server Running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
