package handler

import (
	"github.com/Munovv/go-url-crop/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.POST("/url-crop", h.cropUrl)
		api.GET("/get-url", h.getUrl)
	}

	return router
}
