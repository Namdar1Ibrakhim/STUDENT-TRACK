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
	rout	auth := router.Group("/auth")
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

<<<<<<< HEAD
	instructorRoutes := router.Group("/instructor")
	{
		instructorRoutes.POST("/sign-up", h.signUpInstrucor)
		instructorRoutes.POST("/sign-in", h.signIn) // один метод для всех полей
		instructorRoutes.GET("/student/:id", h.getStudentById) // один метод для всех полей

	}

	adminRoutes := router.Group("/admin")
	{
		adminRoutes.POST("/sign-up", h.signUpAdmin)
		adminRoutes.POST("/sign-in", h.signIn) // один метод для всех полей
		adminRoutes.GET("/user/:id", h.getUserById)
		

	}

=======
>>>>>>> bd88000fb7bc59bca576ee76a0a0de4ebfcf3a03
GET("/user/:id", h.getUserById)
		

	}

	return router
}
