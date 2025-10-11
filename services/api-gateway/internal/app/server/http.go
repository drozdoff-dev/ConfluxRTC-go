package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"api-gateway/internal/app/config"
	"api-gateway/internal/app/container"
	httpRoutes "api-gateway/internal/transport/http"
)

type HTTPServer struct {
	cfg    *config.Config
	logger *slog.Logger
	engine *gin.Engine
	server *http.Server
}

func SlogMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// выполняем обработку
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		logger.Info("HTTP request",
			"method", method,
			"path", path,
			"status", status,
			"latency", latency,
			"client_ip", c.ClientIP(),
		)
	}
}

func NewHTTPServer(c *container.Container) *HTTPServer {
	router := gin.New()
	router.Use(SlogMiddleware(c.Resources.Logger), gin.Recovery())
	router.RedirectTrailingSlash = false // Отключаем автоматическое перенаправление с "/path/" на "/path"
	gin.SetMode(gin.DebugMode)           // Нужно изменить на ReleaseMode для продакшена

	// Роуты
	httpRoutes.RegisterRoutes(router, c.Servers.HttpServer.Handlers)

	httpServer := &http.Server{
		Addr:    c.Resources.Cfg.HTTPServer.Address,
		Handler: router,
	}

	return &HTTPServer{
		cfg:    c.Resources.Cfg,
		logger: c.Resources.Logger,
		engine: router,
		server: httpServer,
	}
}

func (h *HTTPServer) Start() error {
	h.logger.Info("HTTP server listening", "addr", h.cfg.HTTPServer.Address)
	if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("http serve: %w", err)
	}
	return nil
}

func (h *HTTPServer) Stop(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return h.server.Shutdown(ctx)
}
