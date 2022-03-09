package data

import (
	"encoding/json"
	"io"
	"math/rand"
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

func (a *Album) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(a)
}

func (a *AlbumRequest) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(a)
}

func AddAlbum(a *AlbumRequest) Album {
	var newAlbum Album

	newAlbum.Id = int64(rand.Intn(999999))
	newAlbum.Artist = a.Artist
	newAlbum.Title = a.Title

	albums = append(albums, &newAlbum)

	return newAlbum
}
