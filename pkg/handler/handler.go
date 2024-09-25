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
		instructorRoutes := router.Group("/instructor")
		{
			instructorRoutes.POST("/sign-up", h.signUpInstructor)
			instructorRoutes.POST("/sign-in", h.signIn)            // один метод для всех полей
			instructorRoutes.GET("/student/:id", h.getStudentById) // один метод для всех полей

		}

		adminRoutes := router.Group("/admin")
		{
			adminRoutes.POST("/sign-up", h.signUpAdmin)
			adminRoutes.POST("/sign-in", h.signIn) // один метод для всех полей
			adminRoutes.GET("/user/:id", h.getUserById)

		}
		//authUser.PUT("/update", h.updateUser)
		//authUser.DELETE("/delete", h.updateUser)
	}

	return router
}
