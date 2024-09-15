package handler

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/service"
	"github.com/gin-gonic/gin"
)

// Конструктор
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	authUser := router.Group("/auth")
	{
		authUser.POST("/sign-up", h.signUp)
		authUser.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		api.GET("/getInfo", h.getUser)
	}

	return router
}
