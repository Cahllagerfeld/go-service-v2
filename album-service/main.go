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

	r := mux.NewRouter()
	getRouter := r.Methods(http.MethodGet).Subrouter()
	postRouter := r.Methods(http.MethodPost).Subrouter()
	deleteRouter := r.Methods(http.MethodDelete).Subrouter()
	putRouter := r.Methods(http.MethodPut).Subrouter()

	nc := protos.NewNumberClient(conn)

	postRouter.HandleFunc("/albums", handlers.AddAlbum)

	getRouter.HandleFunc("/", handlers.ListAlbums)

	deleteRouter.HandleFunc("/albums/{id}", handlers.RemoveAlbum)

	putRouter.HandleFunc("/albums/{id}", handlers.UpdateAlbum)

	http.ListenAndServe(":9090", r)
}
