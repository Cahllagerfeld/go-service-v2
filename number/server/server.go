package server

import (
	"context"
	"fmt"
	"math/rand"

	protos "github.com/cahllagerfeld/go-service-v2/number/proto"
)

type Random struct{}

func NewNumber() *Random {
	return &Random{}
}

func (r *Random) GetRandomNumber(ctx context.Context, req *protos.GetRandomNumberRequest) (*protos.GetRandomNumberResponse, error) {
	rand := rand.Intn(999999)
	fmt.Println("Random Number:", rand)
	return &protos.GetRandomNumberResponse{Rand: int64(rand)}, nil

}
