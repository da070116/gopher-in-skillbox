package handler

import (
	"github.com/gin-gonic/gin"
	"skillbox-test/pkg/service"
)

type Handler struct {
	services *service.Service
}

// NewHandler - constructor
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes - configure all routes
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	users := router.Group("/user")
	{
		users.GET("/list", h.getUsers)
		users.POST("/add", h.addUser)
		users.DELETE("/:id", h.deleteUser)
		users.PATCH("/:id", h.patchUser)
		users.PATCH(":id/befriend", h.addFriend)
	}
	return router
}
