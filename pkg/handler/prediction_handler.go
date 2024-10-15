package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (h *Handler) getAllPrediction(c *gin.Context) {
	prediction, err := h.services.GetAllPrediction() // Вызов метода GetAll из сервиса
	if err != nil {
		// Если произошла ошибка, отправляем статус 500 и сообщение
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем статус 200 и список курсов в формате JSON
	c.JSON(http.StatusOK, prediction)
}


func (h *Handler) getPredictionById(c *gin.Context) {
    predictionId := c.Param("id") // Получаем ID из параметров маршрута

    // Преобразуем ID в int
    id, err := strconv.Atoi(predictionId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prediction ID"})
        return
    }

    prediction, err := h.services.GetPredictionById(id) // Вызов метода GetById из сервиса
    if err != nil {
        // Если ошибка, например, курс не найден
        c.JSON(http.StatusNotFound, gin.H{"error": "Prediction not found"})
        return
    }

    // Возвращаем статус 200 и курс в формате JSON
    c.JSON(http.StatusOK, prediction)
}

// Хэндлер для получения предикта по идентификатору студента
func (h *Handler) getPredictionByStudentId(c *gin.Context) {
	studentId := c.Param("studentId") // Получаем ID студента из параметров маршрута

	// Преобразуем ID в int
	id, err := strconv.Atoi(studentId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
		return
	}

	prediction, err := h.services.Prediction.GetPredictionByStudentId(id) // Вызов метода GetPredictionByStudentId из сервиса
	if err != nil {
		// Если ошибка, например, курс не найден
		c.JSON(http.StatusNotFound, gin.H{"error": "Prediction not found"})
		return
	}

	// Возвращаем статус 200 и курс в формате JSON
	c.JSON(http.StatusOK, prediction)
}

// Хэндлер для получения предикта по идентификатору направления
func (h *Handler) getPredictionByDirectionId(c *gin.Context) {
	directionId := c.Param("directionId") // Получаем ID студента из параметров маршрута

	// Преобразуем ID в int
	id, err := strconv.Atoi(directionId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid direction ID"})
		return
	}

	prediction, err := h.services.Prediction.GetPredictionByDirectionId(id) // Вызов метода GetPredictionByDirectionId из сервиса
	if err != nil {
		// Если ошибка, например, курс не найден
		c.JSON(http.StatusNotFound, gin.H{"error": "Prediction not found"})
		return
	}

	// Возвращаем статус 200 и курс в формате JSON
	c.JSON(http.StatusOK, prediction)
}