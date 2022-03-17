package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	albumProtos "github.com/cahllagerfeld/go-service-v2/album/proto"
	"github.com/gorilla/mux"
)

type AlbumRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

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
	var album AlbumRequest
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

	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	resp, err := a.ac.DeleteAlbumById(context.Background(), &albumProtos.DeleteAlbumRequest{Id: id})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	_ = resp
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

}

func (a *Albums) UpdateAlbum(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
	}
	var album AlbumRequest
	err = json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	resp, err := a.ac.ReplaceAlbumById(context.Background(), &albumProtos.ReplaceAlbumRequest{Id: id, Title: album.Title, Artist: album.Artist})
	if err != nil {
		http.Error(rw, err.Error(), http.StatusNotFound)
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusNoContent)
	json.NewEncoder(rw).Encode(resp)
}
