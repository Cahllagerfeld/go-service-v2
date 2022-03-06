package albums

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Album struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

var albums = []Album{
	{
		Id:     1,
		Title:  "Aghori",
		Artist: "Kool Savas",
	},
}

func ListAlbums(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	jsonAlbums, err := json.Marshal(albums)

	fmt.Println(string(jsonAlbums))

	if err != nil {
		http.Error(rw, "Failed to parse Albums", http.StatusInternalServerError)
	}

	rw.Write((jsonAlbums))
}
