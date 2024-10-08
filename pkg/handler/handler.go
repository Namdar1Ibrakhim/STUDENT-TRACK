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
			instructorRoutes.GET("/student/:id", h.getStudentById) // получение студента по айди

		}

		adminRoutes := router.Group("/admin")
		{
			adminRoutes.POST("/sign-up", h.signUpAdmin)
			adminRoutes.POST("/sign-in", h.signIn)                                 // один метод для всех полей
			adminRoutes.GET("/user/:id", h.getUserById)                            // получение пользователя по айди
			adminRoutes.GET("/editPassword/:id/:password", h.editPasswordByUserId) //изменить пароль может только админ для всех пользователей
		}

	}

	// Пример изменения профиля юзеров по ID и проверяет доступ на админа
	profile := router.Group("/profile", h.userIdentity)
	{
		profile.PUT("/users/:id", h.UpdateUser)
		profile.DELETE("/users/:id", h.DeleteUser) //удаление акк через админ и юзер
		profile.GET("/editPassword/:password", h.editPasswordByCurrentUserId)

	}

	main := router.Group("/main", h.userIdentity)
	{
		main.POST("/upload-csv", h.UploadCSV)
		/////////////
		//Example CSV file

		//subjectName,Grade   --- Headers
		//Math,90             |   data
		//Physic,70           |
		//Discrete math,83    |
	}
	
	course := router.Group("/course", h.userIdentity)
	{
		course.GET("/getAll", h.getAllCourse)
		course.GET("/getById/:id", h.getCourseById)
		course.GET("/getByName/:name", h.getCourseByName)
	}

	return router
}
