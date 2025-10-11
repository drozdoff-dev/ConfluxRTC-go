package server

// import (
// 	"api-gateway/internal/app/container"
// 	"log/slog"
// )

// func Start(c *container.Container) {
// 	log := c.Resources.Logger

// 	// Start HTTP server
// 	httpServer := NewHTTPServer(c)
// 	go func() {
// 		if err := httpServer.Start(); err != nil {
// 			log.Error("HTTP server stopped", slog.Any("error", err))
// 		}
// 	}()

// 	log.Info("all servers started")
// }
