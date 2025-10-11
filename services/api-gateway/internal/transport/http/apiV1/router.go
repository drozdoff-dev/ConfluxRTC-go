package apiv1

import (
	"github.com/gin-gonic/gin"

	"api-gateway/internal/app/container"
)

func RegisterRoutes(rg *gin.RouterGroup, handlers *container.HttpHandlers) {
	// API routes
	v1 := rg.Group("/v1")
	{
		_ = v1
		// v1.POST("/users/login", users.UsersHandler(svcContainer.UserService))
	}
}
