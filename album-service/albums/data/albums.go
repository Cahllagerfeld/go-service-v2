package data

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	protos "github.com/cahllagerfeld/go-service-v2/number-service/proto"
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

func AddAlbum(a *AlbumRequest, nc protos.NumberClient) error {
	var newAlbum *Album
	resp, err := nc.GetRandomNumber(context.Background(), &protos.GetRandomNumberRequest{})
	if err != nil {
		return fmt.Errorf("Failed to get random number")
	}
	newAlbum.Id = int(resp.Rand)
	newAlbum.Artist = a.Artist
	newAlbum.Title = a.Title

	albums = append(albums, newAlbum)

	return nil
}

func RemoveAlbum(id int) {
	for i, a := range albums {
		if a.Id == id {
			albums = append(albums[:i], albums[i+1:]...)
		}
	}
}

func UpdateAlbum(id int, a *AlbumRequest) error {
	_, pos, err := findAlbum(id)
	if err != nil {
		return err
	}

	album := Album{
		Title:  a.Title,
		Artist: a.Artist,
	}

	album.Id = id
	albums[pos] = &album

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
