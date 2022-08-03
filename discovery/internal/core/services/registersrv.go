package registersrv

import (
	"fmt"

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

func (serv *Service) register(service domain.Service) (*domain.Service, error) {
	result, err := serv.repo.Save(service)
	if err != nil {
		err = fmt.Errorf("Failed to register service: %s", err)
		return nil, err
	}
	return result, nil
}

func (serv *Service) getAll() ([]*domain.Service, error) {
	result, err := serv.repo.GetAll()
	if err != nil {
		err = fmt.Errorf("Failed to get all: %s", err)
		return nil, err
	}
	return result, nil
}

func (serv *Service) unregister(id string) error {
	err := serv.repo.Delete(id)
	if err != nil {
		err = fmt.Errorf("Failed to delete: %s", err)
		return err
	}
	return nil
}
