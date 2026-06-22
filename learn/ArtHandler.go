package main

import (
	"html/template"
	"learn/ascii"
	"net/http"
)

func HandleArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	if text == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	lban := "banners/" + banner + ".txt"
	ban, err := ascii.LoadBanner(lban)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	genArt := ascii.GenerateArt(text, ban)
	data := PageData{
		Result: genArt,
		Text: text,
	}

	templ, err := template.ParseFiles(
		"templates/base.html",
		"templates/input.html",
	)

	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, data)
}
