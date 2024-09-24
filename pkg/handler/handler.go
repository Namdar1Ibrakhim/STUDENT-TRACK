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

	auth := router.Group("/auth")
	{
		authStudent := auth.Group("/student")
		{
			authStudent.POST("/sign-up", h.signUpStudent)
			authStudent.POST("/sign-in", h.signIn) // один метод для всех полей
		}
		authInstructor := auth.Group("/instructor")
		{
			authInstructor.POST("/sign-up", h.signUpInstructor)
			authInstructor.POST("/sign-in", h.signIn) // один метод для всех полей
		}
		authAdmin := auth.Group("/admin")
		{
			authAdmin.POST("/sign-up", h.signUpAdmin)
			authAdmin.POST("/sign-in", h.signIn) // один метод для всех полей
		}
		auth.GET("/get", h.getUser)
		//authUser.PUT("/update", h.updateUser)
		//authUser.DELETE("/delete", h.updateUser)
	}

	return router
}
