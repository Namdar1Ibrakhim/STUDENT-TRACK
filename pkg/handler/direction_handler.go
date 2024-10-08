package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


func (h *Handler) getAllDirection(c *gin.Context) {
	directions, err := h.services.GetAllDirection() // Вызов метода GetAll из сервиса
	if err != nil {
		// Если произошла ошибка, отправляем статус 500 и сообщение
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Возвращаем статус 200 и список курсов в формате JSON
	c.JSON(http.StatusOK, directions)
}


func (h *Handler) getDirectionById(c *gin.Context) {
    directionId := c.Param("id") // Получаем ID из параметров маршрута

    // Преобразуем ID в int
    id, err := strconv.Atoi(directionId)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid direction ID"})
        return
    }

    direction, err := h.services.GetDirectionById(id) // Вызов метода GetById из сервиса
    if err != nil {
        // Если ошибка, например, курс не найден
        c.JSON(http.StatusNotFound, gin.H{"error": "Direction not found"})
        return
    }

    // Возвращаем статус 200 и курс в формате JSON
    c.JSON(http.StatusOK, direction)
}

func (h *Handler) getDirectionByName(c *gin.Context) {
    directionName := c.Param("name") // Получаем имя курса из параметров маршрута

    direction, err := h.services.GetDirectionByName(directionName) // Вызов метода GetByName из сервиса
    if err != nil {
        // Если ошибка, например, курс не найден
        c.JSON(http.StatusNotFound, gin.H{"error": "Direction not found"})
        return
    }

    // Возвращаем статус 200 и курс в формате JSON
    c.JSON(http.StatusOK, direction)
}

