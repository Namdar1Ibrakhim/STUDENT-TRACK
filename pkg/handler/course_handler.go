package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (h *Handler) getAllCourse(c *gin.Context) {
	courses, err := h.services.GetAllCourse() // Вызов метода GetAll из сервиса
	if err != nil {
		// Если произошла ошибка, отправляем статус 500 и сообщение
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем статус 200 и список курсов в формате JSON
	c.JSON(http.StatusOK, courses)
}


func (h *Handler) getCourseById(c *gin.Context) {
    courseId := c.Param("id") // Получаем ID из параметров маршрута

    // Преобразуем ID в int
    id, err := strconv.Atoi(courseId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid course ID"})
        return
    }

    course, err := h.services.Course.GetCourseById(id) // Вызов метода GetById из сервиса
    if err != nil {
        // Если ошибка, например, курс не найден
        c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        return
    }

    // Возвращаем статус 200 и курс в формате JSON
    c.JSON(http.StatusOK, course)
}

func (h *Handler) getCourseByName(c *gin.Context) {
    courseName := c.Param("name") // Получаем имя курса из параметров маршрута

    course, err := h.services.Course.GetCourseByName(courseName) // Вызов метода GetByName из сервиса
    if err != nil {
        // Если ошибка, например, курс не найден
        c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
        return
    }

    // Возвращаем статус 200 и курс в формате JSON
    c.JSON(http.StatusOK, course)
}

