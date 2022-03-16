package data

import (
	"context"
	"fmt"

	albumProtos "github.com/cahllagerfeld/go-service-v2/album/proto"
	numerProtos "github.com/cahllagerfeld/go-service-v2/number/proto"
)

type Albums []*albumProtos.AlbumResponse

var albums = []*albumProtos.AlbumResponse{
	{
		Id:     1,
		Title:  "Aghori",
		Artist: "Kool Savas",
	},
}

func GetAlbum() *albumProtos.GetAllAlbumsResponse {
	return &albumProtos.GetAllAlbumsResponse{Albums: albums}
}

func AddAlbum(a *albumProtos.CreateAlbumRequest, nc numerProtos.NumberClient) (*albumProtos.AlbumResponse, error) {
	resp, err := nc.GetRandomNumber(context.Background(), &numerProtos.GetRandomNumberRequest{})
	if err != nil {
		return &albumProtos.AlbumResponse{}, nil
	}
	var newAlbum = &albumProtos.AlbumResponse{Id: resp.Rand, Title: a.Title, Artist: a.Artist}
	return newAlbum, nil
}

func RemoveAlbum(id int64) *albumProtos.DeleteAlbumResponse {
	for i, a := range albums {
		if a.Id == id {
			albums = append(albums[:i], albums[i+1:]...)
		}
	}

	return &albumProtos.DeleteAlbumResponse{}
}

func findAlbum(id int64) (*albumProtos.AlbumResponse, error) {
	for _, a := range albums {
		if a.Id == id {
			return a, nil
		}
	}
	return nil, fmt.Errorf("Album not found")
}
