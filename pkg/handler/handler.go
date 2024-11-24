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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}


func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger(), CORSMiddleware()) 

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/sign-in", h.signIn) // Единый эндпоинт входа

			auth.POST("/sign-up/student", h.signUpStudent)
			auth.POST("/sign-up/instructor", h.signUpInstructor)
			auth.POST("/sign-up/admin", h.signUpAdmin)
		}

		protected := v1.Group("")
		protected.Use(h.userIdentity)
		{
			profile := protected.Group("/profile")
			{
				profile.GET("", h.getUser)
				profile.PUT("", h.UpdateUser)
				profile.DELETE("", h.DeleteUser)
				profile.PATCH("/password", h.editPasswordByUser)
			}

			admin := protected.Group("/admin")
			{
				admin.GET("/users", h.getAllUsers)
				admin.GET("/users/:id", h.getUserById)
				admin.PUT("/users/:id", h.UpdateUser)
				admin.DELETE("/users/:id", h.DeleteUser)
				admin.PATCH("/users/:id/password", h.editPasswordByAdmin)
			}

			instructor := protected.Group("/instructor")
			{
				instructor.GET("/students", h.getStudents)
				instructor.GET("/students/:id", h.getStudentById)
			}
			predict := protected.Group("/predict")
			{
				predict.POST("/upload", h.UploadCSV)
				predict.POST("/analyze", h.PredictCSV)
				/*Example CSV file
					_____________________________________________________________________________
					subject1, subject2,.... subject7, Hackathons attended, Topmost Certification, -> continue                  ---|Headers
					70,       70,           90,       1,                   DBMS Certification,
					_________________________________________________________________________________________________________
				  ->Personality, Management or technical, Leadership, Team, Self Ability | IF role == Instructor, + Student_id ---|Headers
					Extravert,   Management,               NO,         YES,  NO,                                   220202222

				*/
			}

			courses := protected.Group("/courses")
			{
				courses.GET("", h.getAllCourse)
				courses.GET("/:id", h.getCourseById)
				courses.GET("/search", h.getCourseByName)
			}

			directions := protected.Group("/directions")
			{
				directions.GET("", h.getAllDirection)
				directions.GET("/:id", h.getDirectionById)
				directions.GET("/search", h.getDirectionByName)
			}

			predictions := protected.Group("/predictions")
			{
				predictions.GET("", h.getAllPrediction)
				predictions.GET("/:id", h.getPredictionById)
				predictions.GET("/student/:studentId", h.getPredictionByStudentId)
				predictions.GET("/direction/:directionId", h.getPredictionByDirectionId)
			}

			studentCourses := protected.Group("/student-courses")
			{
				studentCourses.GET("", h.getAllStudentCourse)
				studentCourses.GET("/:id", h.getStudentCourseById)
				studentCourses.GET("/student/:studentId", h.getStudentCourseByStudentId)
				studentCourses.GET("/course/:courseId", h.getStudentCourseByCourseId)
				studentCourses.GET("/filter", h.getAllStudentCourseByFilter)
			}
		}
	}

	return router
}
