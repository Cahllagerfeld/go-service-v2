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

func GetAlbums() *albumProtos.GetAllAlbumsResponse {
	return &albumProtos.GetAllAlbumsResponse{Albums: albums}
}

func AddAlbum(a *albumProtos.CreateAlbumRequest, nc numerProtos.NumberClient) (*albumProtos.AlbumResponse, error) {
	resp, err := nc.GetRandomNumber(context.Background(), &numerProtos.GetRandomNumberRequest{})
	if err != nil {
		return &albumProtos.AlbumResponse{}, nil
	}
	var newAlbum = &albumProtos.AlbumResponse{Id: resp.Rand, Title: a.Title, Artist: a.Artist}
	albums = append(albums, newAlbum)
	return newAlbum, nil
}

func RemoveAlbum(a *albumProtos.DeleteAlbumRequest) *albumProtos.DeleteAlbumResponse {
	for i, alb := range albums {
		if alb.Id == a.Id {
			albums = append(albums[:i], albums[i+1:]...)
		}
	}

	return &albumProtos.DeleteAlbumResponse{}
}

func UpdateAlbum(a *albumProtos.ReplaceAlbumRequest) (*albumProtos.AlbumResponse, error) {
	for i, alb := range albums {
		if alb.Id == a.Id {
			albums[i].Title = a.Title
			albums[i].Artist = a.Artist

			return &albumProtos.AlbumResponse{Id: albums[i].Id, Title: albums[i].Title, Artist: albums[i].Artist}, nil
		}
	}
	return &albumProtos.AlbumResponse{}, fmt.Errorf("Album not found")
}

func FindAlbum(a *albumProtos.GetSingleAlbumRequest) (*albumProtos.AlbumResponse, error) {
	for _, a := range albums {
		if a.Id == a.Id {
			return a, nil
		}
	}
	return nil, fmt.Errorf("Album not found")
}
