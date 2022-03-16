package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	albumProtos "github.com/cahllagerfeld/go-service-v2/album/proto"
	"github.com/cahllagerfeld/go-service-v2/gateway/albums/data"
	"github.com/gorilla/mux"
)

type Albums struct {
	ac albumProtos.AlbumClient
}

func NewAlbums(ac albumProtos.AlbumClient) *Albums {
	return &Albums{ac: ac}
}

func (a *Albums) ListAlbums(rw http.ResponseWriter, r *http.Request) {

	resp, err := a.ac.GetAllAlbums(context.Background(), &albumProtos.GetAllAlbumsRequest{})
	if err != nil {
		http.Error(rw, "Failed to get all Albums", http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(resp.Albums)

}

func (a *Albums) AddAlbum(rw http.ResponseWriter, r *http.Request) {
	var album data.AlbumRequest
	err := json.NewDecoder(r.Body).Decode(&album)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	resp, err := a.ac.CreateAlbum(context.Background(), &albumProtos.CreateAlbumRequest{Title: album.Title, Artist: album.Artist})

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(resp)

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
