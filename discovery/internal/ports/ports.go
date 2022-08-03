package ports

import "github.com/cahllagerfeld/go-service-v2/discovery/internal/core/domain"

type DiscoveryRepository interface {
	Save(domain.Service) (*domain.Service, error)
	Delete(id string) error
	GetAll() ([]*domain.Service, error)
}

type DiscoveryService interface {
	register(service domain.Service) (*domain.Service, error)
	getAll() ([]*domain.Service, error)
	unregister(id string) error
}
