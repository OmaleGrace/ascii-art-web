package main

import (
	"ascii-art-web/ascii"
	"net/http"
)

var generatedArt string

func HandleArt(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" || banner == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	switch banner {
	case "standard", "shadow", "thinkertoy":
	default:
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	hold := "banner/" + banner + ".txt"

	ban, err := ascii.LoadBanner(hold)
	if err != nil {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	GenArt := ascii.GenerateArt(text, ban)

	generatedArt = GenArt

	http.Redirect(w, r, "/result", http.StatusSeeOther)
}
