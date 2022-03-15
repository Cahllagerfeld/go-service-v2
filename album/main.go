package main

import (
	"fmt"
	"log"
	"net"
	"os"

	albumProtos "github.com/cahllagerfeld/go-service-v2/album/proto"
	"github.com/cahllagerfeld/go-service-v2/album/server"

	"google.golang.org/grpc"
)

func main() {

	gs := grpc.NewServer()

	alb := server.NewAlbum()

	albumProtos.RegisterAlbumServer(gs, alb)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Fatal("Unable to create Server", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
