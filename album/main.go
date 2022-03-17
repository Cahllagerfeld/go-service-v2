package main

import (
	"fmt"
	"log"
	"net"
	"os"

	albumProtos "github.com/cahllagerfeld/go-service-v2/album/proto"
	"github.com/cahllagerfeld/go-service-v2/album/server"
	numberProtos "github.com/cahllagerfeld/go-service-v2/number/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	nc := numberProtos.NewNumberClient(conn)

	gs := grpc.NewServer()

	alb := server.NewAlbum(nc)

	albumProtos.RegisterAlbumServer(gs, alb)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Fatal("Unable to create Server", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)
}
