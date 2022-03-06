package main

import (
	"net/http"

	"go-service-v2/internal/albums"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	getRouter := r.Methods("GET").Subrouter()

	getRouter.HandleFunc("/", albums.ListAlbums)
	http.ListenAndServe(":9090", r)
}
