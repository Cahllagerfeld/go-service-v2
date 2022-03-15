package server

import (
	"context"

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

}
