package healthz

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type HealthHandler struct {
	logger  *slog.Logger
	service Service
}

func New(log *slog.Logger, svc Service) *HealthHandler {
	return &HealthHandler{logger: log, service: svc}
}

type HealthStatus struct {
	Status    string            `json:"status"`
	Timestamp time.Time         `json:"timestamp"`
	Services  map[string]string `json:"services"`
}

type Service interface {
	CheckHealth(ctx context.Context) *HealthStatus
}

func (h *HealthHandler) HealthzHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	status := h.service.CheckHealth(ctx)

	w.Header().Set("Content-Type", "application/json")

	if status.Status == "unhealthy" {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		w.WriteHeader(http.StatusOK)
	}

	json.NewEncoder(w).Encode(status)
}
