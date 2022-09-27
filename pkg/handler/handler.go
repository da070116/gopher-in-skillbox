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

	users := router.Group("/city")
	{
		users.GET("/list", h.getCities)
		users.GET("/filter/:param", h.filterCities)
		users.POST("/add", h.addCity)
		users.DELETE("/:id", h.deleteCity)
		users.PATCH("/:id", h.patchCity)

	}
	return router
}
