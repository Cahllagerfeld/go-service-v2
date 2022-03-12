package server
import (
	protos "github.com/cahllagerfeld/go-service-v2/number-service/proto/number_service"
)

type Random struct{}

func NewNumber() *Random {
	return &Random{}
}

func (r *Random) GetRandomNumber() *protos.GetRandomNumberResponse {

}
