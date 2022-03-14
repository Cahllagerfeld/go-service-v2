package main

import (
	"flag"
	"net/http"

	"github.com/cahllagerfeld/go-service-v2/album-service/albums/handlers"
	protos "github.com/cahllagerfeld/go-service-v2/number-service/proto"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	addr = flag.String("addr", "localhost:9091", "the address to connect to")
)

func main() {
	flag.Parse()

	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	nc := protos.NewNumberClient(conn)

	handler := handlers.NewAlbums(nc)

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
