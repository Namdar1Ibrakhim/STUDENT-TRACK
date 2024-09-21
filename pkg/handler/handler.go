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
		authUser.POST("/sign-up", h.signUpStudent) 
		authUser.POST("/sign-in", h.signIn) // один метод для всех полей
		authUser.GET("/get", h.getUser)
		//authUser.PUT("/update", h.updateUser)
		//authUser.DELETE("/delete", h.updateUser)
	}

	instructorRoutes := router.Group("/instrucor")
	{
		instructorRoutes.POST("/sign-up", h.signUpInstrucor)
		instructorRoutes.POST("/sign-in", h.signIn) // один метод для всех полей
	}

	adminRoutes := router.Group("/admin")
	{
		adminRoutes.POST("/sign-up", h.signUpAdmin)
		adminRoutes.POST("/sign-in", h.signIn) // один метод для всех полей
	}

	return router
}
