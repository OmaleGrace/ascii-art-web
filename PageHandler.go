package main

import (
	"html/template"
	"net/http"
)

func HandlePage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 Invalid Method", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	tem, err := template.ParseFiles(
		"templates/base.html",
		"templates/input.html",
	)
	
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tem.Execute(w, nil)
	if err != nil {
		http.Error(w, "500 Internal server Error", http.StatusInternalServerError)
		return
	}
}
