package bootstrap

import (
	// "context"
	"log/slog"
	"os"

	"api-gateway/internal/app/config"
	"api-gateway/internal/app/container"
)

// BuildContainer собирает DI контейнер приложения
func BuildContainer() (*container.Container, error) {
	// ctx := context.Background()

	log := SetupLogger(os.Getenv("ENV"))
	log.Info("logger initialized")

	cfg := config.New(log)
	log.Info("config loaded", slog.String("env", cfg.Env))

	// Собираем DI контейнер
	c, err := container.New(log, cfg)
	if err != nil {
		return nil, err
	}

	return c, nil
}
