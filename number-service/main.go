package main

import (
	"fmt"
	"log"
	"net"
	"os"

	protos "github.com/cahllagerfeld/go-service-v2/number-service/proto"
	"github.com/cahllagerfeld/go-service-v2/number-service/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()

	numb := server.NewNumber()

	protos.RegisterNumberServer(gs, numb)

	reflection.Register(gs)

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", 9091))
	if err != nil {
		log.Fatal("Unable to create listener", "error", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
