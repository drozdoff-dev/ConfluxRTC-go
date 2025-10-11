package http

import (
	"github.com/gin-gonic/gin"

	"api-gateway/internal/app/container"

	v1 "api-gateway/internal/transport/http/apiV1"
)

func RegisterRoutes(r *gin.Engine, handlers *container.HttpHandlers) {
	api := r.Group("/api")
	{
		v1.RegisterRoutes(api, handlers)
	}
	r.GET("/healthz", gin.WrapF(handlers.Healthz.HealthzHandler))
}
