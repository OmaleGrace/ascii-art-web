package main

import (
	"net/http"
	"html/template"
)

func HandlePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w ,"400 Bad Request", http.StatusBadRequest)
		return
	}
	
	if r.URL.Path != "/" {
		http.Error(w,  "404 Not Found", http.StatusNotFound)
		return
	}
	
	templ, err := template.ParseFiles("templates/input.html")
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	err = templ.Execute(w, nil)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}