package handlers

import (
	"encoding/json"
	"github.com/Aryan-9211/url-shortner/db"
	"github.com/Aryan-9211/url-shortner/utils"
	"net/http"
)

type ShortenRequest struct {
	OriginalURL string `json:"original_url"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	shortURL := utils.GenerateShortURL(req.OriginalURL)

	if err := db.SaveURLMapping(shortURL, req.OriginalURL); err != nil {
		http.Error(w, "Failed to save URL mapping", http.StatusInternalServerError)
		return
	}

	resp := ShortenResponse{ShortURL: "http://localhost:8080/" + shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
