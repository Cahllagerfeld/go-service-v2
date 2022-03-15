package handlers

import (
	"net/http"
	"strconv"

	"github.com/cahllagerfeld/go-service-v2/gateway/albums/data"
	protos "github.com/cahllagerfeld/go-service-v2/number/proto"
	"github.com/gorilla/mux"
)

type Albums struct {
	nc protos.NumberClient
}

func NewAlbums(nc protos.NumberClient) *Albums {
	return &Albums{nc: nc}
}

func (a *Albums) ListAlbums(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	albums := data.GetAlbums()

	err := albums.ToJson(rw)

	if err != nil {
		http.Error(rw, "Failed to parse Albums", http.StatusInternalServerError)
	}

}

func (a *Albums) AddAlbum(rw http.ResponseWriter, r *http.Request) {

	rw.WriteHeader(http.StatusCreated)
	var album data.AlbumRequest
	err := album.FromJson(r.Body)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	data, err := data.AddAlbum(&album, a.nc)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	err = data.ToJson(rw)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}
}

func (a *Albums) RemoveAlbum(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusNoContent)
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}

	data.RemoveAlbum(id)
}

func (a *Albums) UpdateAlbum(rw http.ResponseWriter, r *http.Request) {
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
