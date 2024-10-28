package handler

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/service"
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

	auth := router.Group("/auth")
	{
		authStudent := auth.Group("/student")
		{
			authStudent.POST("/sign-up", h.signUpStudent)
			authStudent.POST("/sign-in", h.signIn) // один метод для всех полей
		}
		instructorRoutes := auth.Group("/instructor")
		{
			instructorRoutes.POST("/sign-up", h.signUpInstructor)
			instructorRoutes.POST("/sign-in", h.signIn)            // один метод для всех полей
			instructorRoutes.GET("/student/:id", h.getStudentById) // получение студента по айди
			instructorRoutes.GET("/get", h.getUser)

		}

		adminRoutes := auth.Group("/admin")
		{
			adminRoutes.POST("/sign-up", h.signUpAdmin)
			adminRoutes.POST("/sign-in", h.signIn)
			adminRoutes.GET("/get", h.getUser)
			adminRoutes.GET("/user/:id", h.getUserById)                            // получение пользователя по айди
			adminRoutes.GET("/editPassword/:id/:password", h.editPasswordByUserId) //изменить пароль может только админ для всех пользователей
			adminRoutes.DELETE("/users/:id", h.DeleteUser)                         //удаление акк через админ и юзер

		}

	}

	profile := router.Group("/profile", h.userIdentity)
	{
		profile.GET("/get", h.getUser)
		profile.PUT("/users/:id", h.UpdateUser)
		profile.DELETE("/users/:id", h.DeleteUser) //удаление акк через админ и юзер
		profile.GET("/editPassword/:password", h.editPasswordByCurrentUserId)

	}

	main := router.Group("/main", h.userIdentity)
	{
		main.POST("/upload-csv", h.UploadCSV)
		main.POST("/upload-csv/predict", h.PredictCSV)
		/*Example CSV file
			_____________________________________________________________________________
			subject1, subject2,.... subject7, Hackathons attended, Topmost Certification, -> continue                  ---|Headers
			70,       70,           90,       1,                   DBMS Certification,
			_________________________________________________________________________________________________________
		  ->Personality, Management or technical, Leadership, Team, Self Ability | IF role == Instructor, + Student_id ---|Headers
			Extravert,   Management,               NO,         YES,  NO,                                   220202222

		*/
	}

	course := router.Group("/course", h.userIdentity)
	{
		course.GET("/getAll", h.getAllCourse)
		course.GET("/getById/:id", h.getCourseById)
		course.GET("/getByName/:name", h.getCourseByName)
	}

	direction := router.Group("/direction", h.userIdentity)
	{
		direction.GET("/getAll", h.getAllDirection)
		direction.GET("/getById/:id", h.getDirectionById)
		direction.GET("/getByName/:name", h.getDirectionByName)
	}

	prediction := router.Group("/prediction", h.userIdentity)
	{
		prediction.GET("/getAll", h.getAllPrediction)
		prediction.GET("/getById/:id", h.getPredictionById)
		prediction.GET("/getByStudentId/:studentId", h.getPredictionByStudentId)
		prediction.GET("/getByDirectionId/:directionId", h.getPredictionByDirectionId)
	}

	student_course := router.Group("/studentCourse", h.userIdentity)
	{
		student_course.GET("/getAll", h.getAllStudentCourse)
		student_course.GET("/getById/:id", h.getStudentCourseById)
		student_course.GET("/getByStudentId/:studentId", h.getStudentCourseByStudentId)
		student_course.GET("/getByCourseId/:courseId", h.getStudentCourseByCourseId)
		student_course.GET("/getByFilter", h.getAllStudentCourseByFilter)

	}

	return router
}
