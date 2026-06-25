package main

import (
	"github.com/omalegrace2009-g/ascii-art-web/ascii"

	"html/template"
	"net/http"
)

func HandleSwitch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "405 Wrong Method", http.StatusMethodNotAllowed)
		return
	}

	text := r.URL.Query().Get("text")
	banner := r.URL.Query().Get("banner")

	if text == "" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	file := "banner/" + banner + ".txt"
	banners, err := ascii.LoadBanner(file)

	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	genArt := ascii.GenerateArt(text, banners)

	data := Pagedata{
		Text:   text,
		Result: genArt,
		Banner: banner,
	}

	templ, err := template.ParseFiles(
		"templates/base.html",
		"templates/output.html",
	)

	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, data)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	}
}
