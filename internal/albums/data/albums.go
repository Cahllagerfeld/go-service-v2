package data

import (
	"encoding/json"
	"io"
)

type Album struct {
	Id     int64
	Title  string
	Artist string
}

type AlbumRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

type Albums []*Album

var albums = []*Album{
	{
		Id:     1,
		Title:  "Aghori",
		Artist: "Kool Savas",
	},
}

func GetAlbums() Albums {
	return albums
}

func (a *Albums) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}
