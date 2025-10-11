package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"api-gateway/internal/app/bootstrap"
	"api-gateway/internal/app/container"
	"api-gateway/internal/app/server"
)

type App struct {
	httpServer *server.HTTPServer

	container *container.Container
}

func New() *App {
	return &App{}
}

func (a *App) Run() {
	container, err := bootstrap.BuildContainer()
	if err != nil {
		fmt.Println("failed to bootstrap container:", err)
		os.Exit(1)
	}
	a.container = container

	// Start servers in goroutines
	a.startServers()

	// Wait for interrupt signal
	a.gracefulShutdown()
}

func (a *App) startServers() {
	log := a.container.Resources.Logger

	// Start HTTP server
	a.httpServer = server.NewHTTPServer(a.container)
	go func() {
		if err := a.httpServer.Start(); err != nil {
			log.Error("HTTP server stopped", slog.Any("error", err))
		}
	}()

	log.Info("all servers started")
}

func (a *App) gracefulShutdown() {
	log := a.container.Resources.Logger

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	sig := <-quit
	log.Info("received shutdown signal", slog.String("signal", sig.String()))

	// Create shutdown context with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Stop servers
	log.Info("stopping servers...")
	if err := a.httpServer.Stop(shutdownCtx); err != nil {
		log.Error("failed to stop HTTP server", slog.Any("error", err))
	} else {
		log.Info("HTTP server stopped gracefully")
	}

	log.Info("application shutdown completed")
}
