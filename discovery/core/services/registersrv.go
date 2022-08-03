package registersrv

import (
	"github.com/cahllagerfeld/go-service-v2/discovery/core/domain"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (serv *Service) register(service *domain.Service) (*domain.Service, error) {
	return nil, nil
}

func (serv *Service) getAll() ([]*domain.Service, error) {
	return nil, nil
}

func (serv *Service) unregister(id string) (*domain.Service, error) {
	return nil, nil
}
