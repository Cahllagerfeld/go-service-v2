package discoveryhdl

import (
	"github.com/cahllagerfeld/go-service-v2/discovery/internal/ports"
)

type HttpHandler struct {
	discoveryService ports.DiscoveryService
}

func NewHttpHandler(discoveryService ports.DiscoveryService) *HttpHandler {
	return &HttpHandler{
		discoveryService: discoveryService,
	}
}
