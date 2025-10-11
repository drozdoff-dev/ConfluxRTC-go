package container

import (
	"log/slog"
	// "time"

	// usersv2pb "gitlab.dev.keep-calm.ru/keep-calm/loyalty-libs/pkg/proto/example-service/v2/users"
	"api-gateway/internal/app/config"
	// userservice "gitlab.dev.keep-calm.ru/keep-calm/loyalty-service-example/internal/service/api/users"
	// "google.golang.org/grpc"
	// "google.golang.org/grpc/credentials/insecure"
	healthzSvc "api-gateway/internal/service/health"
	"api-gateway/internal/transport/http/healthz"
)

// --- Resources ---
type Resources struct {
	Logger *slog.Logger
	Cfg    *config.Config
}

// --- Servers ---
type Servers struct {
	HttpServer *HttpServer
}

type HttpServer struct {
	Handlers *HttpHandlers
	Services *HttpServices
}

// --- HttpHandlers ---
type HttpHandlers struct {
	Healthz *healthz.HealthHandler
}

// --- HttpServices ---
type HttpServices struct {
	Healthz healthz.Service
}

func newHandlers(r *Resources, s *HttpServices) *HttpHandlers {
	return &HttpHandlers{
		Healthz: healthz.New(r.Logger, s.Healthz),
	}
}

func newServices(r *Resources) *HttpServices {
	return &HttpServices{
		Healthz: healthzSvc.New(),
	}
}

// --- Root Container ---
type Container struct {
	Resources *Resources
	Servers   *Servers
}

func New(
	logger *slog.Logger,
	cfg *config.Config,
) (*Container, error) {

	resources := &Resources{
		Logger: logger,
		Cfg:    cfg,
	}

	services := newServices(resources)

	handlers := newHandlers(resources, services)

	httpServer := &HttpServer{
		Handlers: handlers,
		Services: services,
	}

	servers := &Servers{
		HttpServer: httpServer,
	}

	return &Container{
		Resources: resources,
		Servers:   servers,
	}, nil
}
