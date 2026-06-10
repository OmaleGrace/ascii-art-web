package main

import (
	"ascii-art-web/ascii"
	"html/template"
	"net/http"
)

func HandleArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	switch banner{
	case "standard", "shadow", "thinkertoy":
	default:
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	hold := "banner/" + banner + ".txt"

	ban, err := ascii.LoadBanner(hold)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	GenArt := ascii.GenerateArt(text, ban)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	data := PageData{
		Result: GenArt,
	}

	empl, err := template.ParseFiles("templates/input.html")
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	err = empl.Execute(w, data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}