package main

import (
	"net/http"

	"github.com/cahllagerfeld/go-service-v2/album-service/albums/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()
	postRouter := r.Methods(http.MethodPost).Subrouter()
	deleteRouter := r.Methods(http.MethodDelete).Subrouter()
	putRouter := r.Methods(http.MethodPut).Subrouter()

	postRouter.HandleFunc("/albums", handlers.AddAlbum)

	getRouter.HandleFunc("/", handlers.ListAlbums)

	deleteRouter.HandleFunc("/albums/{id}", handlers.RemoveAlbum)

	putRouter.HandleFunc("/albums/{id}", handlers.UpdateAlbum)

	http.ListenAndServe(":9090", r)
}
