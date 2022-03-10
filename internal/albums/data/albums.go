package data

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
)

type Album struct {
	Id     int
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

	newAlbum.Id = rand.Intn(999999)
	newAlbum.Artist = a.Artist
	newAlbum.Title = a.Title

	albums = append(albums, &newAlbum)

	return newAlbum
}

func RemoveAlbum(id int) {
	for i, a := range albums {
		if a.Id == id {
			albums = append(albums[:i], albums[i+1:]...)
		}
	}
}

func UpdateAlbum(id int, a *Album) error {
	_, pos, err := findAlbum(id)
	if err != nil {
		return err
	}
	a.Id = id
	albums[pos] = a

	return nil
}

func findAlbum(id int) (*Album, int, error) {
	for i, p := range albums {
		if p.Id == id {
			return p, i, nil
		}
	}
	return nil, -1, fmt.Errorf("Album not found")
}
