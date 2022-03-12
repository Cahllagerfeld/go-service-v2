package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ListAlbums(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	albums := data.GetAlbums()

	err := albums.ToJson(rw)

	if err != nil {
		http.Error(rw, "Failed to parse Albums", http.StatusInternalServerError)
	}

}

func AddAlbum(rw http.ResponseWriter, r *http.Request) {

	rw.WriteHeader(http.StatusCreated)
	var album data.AlbumRequest
	err := album.FromJson(r.Body)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	data := data.AddAlbum(&album)

	err = data.ToJson(rw)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func RemoveAlbum(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNoContent)
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	data.RemoveAlbum(id)
}

func UpdateAlbum(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	var album data.AlbumRequest
	err = album.FromJson(r.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	err = data.UpdateAlbum(id, &album)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
	}
}
