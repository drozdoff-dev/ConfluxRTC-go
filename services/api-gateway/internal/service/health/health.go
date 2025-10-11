package health

import (
	"context"
	"time"

	handler "api-gateway/internal/transport/http/healthz"
)

type healthService struct {
}

func New() *healthService {
	return &healthService{}
}

func (h *healthService) CheckHealth(ctx context.Context) *handler.HealthStatus {
	services := make(map[string]string)
	overallStatus := "healthy"

	return &handler.HealthStatus{
		Status:    overallStatus,
		Timestamp: time.Now(),
		Services:  services,
	}
}
