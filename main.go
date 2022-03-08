package main

import (
	"net/http"

	"github.com/cahllagerfeld/go-service-v2/internal/albums/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	getRouter := r.Methods("GET").Subrouter()
	postRouter := r.Methods("POST").Subrouter()

	postRouter.HandleFunc("/albums", handlers.AddAlbum)

	getRouter.HandleFunc("/", handlers.ListAlbums)
	http.ListenAndServe(":9090", r)
}
