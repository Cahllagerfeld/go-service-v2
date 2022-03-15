package server

import (
	"context"

	albumProto "github.com/cahllagerfeld/go-service-v2/album/proto"
)

type Album struct {
	// nc numberProto.NumberClient
}

func NewAlbum() *Album {
	return &Album{}
}

func (a *Album) CreateAlbum(ctx context.Context, req *albumProto.CreateAlbumRequest) (*albumProto.AlbumResponse, error) {
	return &albumProto.AlbumResponse{Id: 43, Title: "Sonny Black", Artist: "Bushido"}, nil
}

func (a *Album) GetAllAlbums(ctx context.Context, req *albumProto.GetAllAlbumsRequest) (*albumProto.GetAllAlbumsResponse, error) {
	var response = []*albumProto.AlbumType{
		{Id: 4, Title: "Aghori", Artist: "Savas"},
	}
	return &albumProto.GetAllAlbumsResponse{Albums: response}, nil
}

func (a *Album) GetAlbumById(ctx context.Context, req *albumProto.GetSingleAlbumRequest) (*albumProto.AlbumResponse, error) {
	return &albumProto.AlbumResponse{Id: 43, Title: "Sonny Black", Artist: "Bushido"}, nil
}

func (a *Album) DeleteAlbumById(ctx context.Context, req *albumProto.DeleteAlbumRequest) (*albumProto.DeleteAlbumResponse, error) {
	return &albumProto.DeleteAlbumResponse{}, nil
}

func (a *Album) ReplaceAlbumById(ctx context.Context, req *albumProto.ReplaceAlbumRequest) (*albumProto.AlbumResponse, error) {
	return &albumProto.AlbumResponse{Id: 43, Title: "Sonny Black", Artist: "Bushido"}, nil
}
