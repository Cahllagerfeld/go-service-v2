package registersrv

import (
	"github.com/cahllagerfeld/go-service-v2/discovery/internal/core/domain"
	"github.com/cahllagerfeld/go-service-v2/discovery/internal/ports"
)

type Service struct {
	repo ports.DiscoveryRepository
}

func New(repository ports.DiscoveryRepository) *Service {
	return &Service{
		repo: repository,
	}
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
