package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllStudentCourse(c *gin.Context) {
	studentCourses, err := h.services.GetAllStudentCourse() // Вызов метода GetAll из сервиса
	if err != nil {
		// Если произошла ошибка, отправляем статус 500 и сообщение
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем статус 200 и список курсов в формате JSON
	c.JSON(http.StatusOK, studentCourses)
}


// Хэндлер для получения всех курсов с пагинацией и сортировкой
func (h *Handler) getAllStudentCourseByFilter(c *gin.Context) {
	// Получаем параметры пагинации и сортировки из запроса
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")
	sortBy := c.Query("sortBy")

	// Преобразуем параметры пагинации в int
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	// Получаем курсы через сервис
	courses, err := h.services.StudentCourse.GetAllStudentCourseByFilter(page, pageSize, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем статус 200 и список курсов в формате JSON
	c.JSON(http.StatusOK, courses)
}

// Хэндлер для получения курса по его идентификатору
func (h *Handler) getStudentCourseById(c *gin.Context) {
	courseId := c.Param("id") // Получаем ID из параметров маршрута

	// Преобразуем ID в int
	id, err := strconv.Atoi(courseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := h.services.StudentCourse.GetStudentCourseById(id) // Вызов метода GetById из сервиса
	if err != nil {
		// Если ошибка, например, курс не найден
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// Возвращаем статус 200 и курс в формате JSON
	c.JSON(http.StatusOK, course)
}

// Хэндлер для получения курса по идентификатору студента
func (h *Handler) getStudentCourseByStudentId(c *gin.Context) {
	studentId := c.Param("studentId") // Получаем ID студента из параметров маршрута

	// Преобразуем ID в int
	id, err := strconv.Atoi(studentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	course, err := h.services.StudentCourse.GetStudentCourseByStudentId(id) // Вызов метода GetStudentCourseByStudentId из сервиса
	if err != nil {
		// Если ошибка, например, курс не найден
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// Возвращаем статус 200 и курс в формате JSON
	c.JSON(http.StatusOK, course)
}

// Хэндлер для получения курса по идентификатору курса
func (h *Handler) getStudentCourseByCourseId(c *gin.Context) {
	courseId := c.Param("courseId") // Получаем ID курса из параметров маршрута

	// Преобразуем ID в int
	id, err := strconv.Atoi(courseId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
		return
	}

	course, err := h.services.StudentCourse.GetStudentCourseByCourseId(id) // Вызов метода GetStudentCourseByCourseId из сервиса
	if err != nil {
		// Если ошибка, например, курс не найден
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// Возвращаем статус 200 и курс в формате JSON
	c.JSON(http.StatusOK, course)
}
