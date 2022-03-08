package handlers

import (
	"net/http"

	"github.com/cahllagerfeld/go-service-v2/internal/albums/data"
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
