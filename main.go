package main

import (
	"github.com/Aryan-9211/url-shortner/db"
	"github.com/Aryan-9211/url-shortner/handlers"
	"log"
	"net/http"
)

func main() {
	db.InitRedis()
	db.InitPostgres()

	http.HandleFunc("/shorten", handlers.ShortenUrlHandler)
	http.HandleFunc("/", handlers.RedirectHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
