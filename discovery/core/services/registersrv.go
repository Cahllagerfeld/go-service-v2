package registersrv

import (
	"github.com/cahllagerfeld/go-service-v2/discovery/domain"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (serv *Service) register(service *domain.Service) (*domain.Service, error) {}
