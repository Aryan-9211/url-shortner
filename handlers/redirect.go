package handlers

import (
	"github.com/Aryan-9211/url-shortner/db"
	"net/http"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {

	shortURL := r.URL.Path[1:]

	originalURL, err := db.GetOriginalURL(shortURL)
	if err != nil || originalURL == "" {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
