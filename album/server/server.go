package server

import (
	"context"

	"github.com/cahllagerfeld/go-service-v2/album/data"
	albumProto "github.com/cahllagerfeld/go-service-v2/album/proto"
	numberProto "github.com/cahllagerfeld/go-service-v2/number/proto"
)

type Album struct {
	nc numberProto.NumberClient
}

func NewAlbum(nc numberProto.NumberClient) *Album {
	return &Album{nc: nc}
}

func (a *Album) CreateAlbum(ctx context.Context, req *albumProto.CreateAlbumRequest) (*albumProto.AlbumResponse, error) {
	return data.AddAlbum(req, a.nc)
}

func (a *Album) GetAllAlbums(ctx context.Context, req *albumProto.GetAllAlbumsRequest) (*albumProto.GetAllAlbumsResponse, error) {
	return data.GetAlbums(), nil
}

func (a *Album) GetAlbumById(ctx context.Context, req *albumProto.GetSingleAlbumRequest) (*albumProto.AlbumResponse, error) {
	return data.FindAlbum(req)
}

func (a *Album) DeleteAlbumById(ctx context.Context, req *albumProto.DeleteAlbumRequest) (*albumProto.DeleteAlbumResponse, error) {
	return data.RemoveAlbum(req), nil
}

func (a *Album) ReplaceAlbumById(ctx context.Context, req *albumProto.ReplaceAlbumRequest) (*albumProto.AlbumResponse, error) {
	return data.UpdateAlbum(req)
}
