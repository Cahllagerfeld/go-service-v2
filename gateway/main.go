package main

import (
	"flag"
	"net/http"

	albumProtos "github.com/cahllagerfeld/go-service-v2/album/proto"
	"github.com/cahllagerfeld/go-service-v2/gateway/albums/handlers"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:9091", "the address to connect to")
)

func main() {
	flag.Parse()

	albumConn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer albumConn.Close()

	ac := albumProtos.NewAlbumClient(albumConn)

	handler := handlers.NewAlbums(ac)

	r := mux.NewRouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()
	postRouter := r.Methods(http.MethodPost).Subrouter()
	deleteRouter := r.Methods(http.MethodDelete).Subrouter()
	putRouter := r.Methods(http.MethodPut).Subrouter()

	postRouter.HandleFunc("/albums", handler.AddAlbum)

	getRouter.HandleFunc("/", handler.ListAlbums)

	deleteRouter.HandleFunc("/albums/{id}", handler.RemoveAlbum)

	putRouter.HandleFunc("/albums/{id}", handler.UpdateAlbum)

	http.ListenAndServe(":9090", r)
}
