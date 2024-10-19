package handler

import (
	"bytes"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"sync"
)

var validatedFiles sync.Map

func (h *Handler) UploadCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to upload file")
		return
	}

	src, err := file.Open()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to open file")
		return
	}
	defer src.Close()

	userIdFromToken, exists := c.Get("userId")
	if !exists {
		newErrorResponse(c, http.StatusUnauthorized, "User Id not found")
		return
	}

	if h.checkRole(c, constants.RoleInstructor) {
		err = h.services.ValidateCSVForInstructor(src)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	} else if h.checkRole(c, constants.RoleStudent) {
		err = h.services.ValidateCSVForStudent(src)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	} else {
		newErrorResponse(c, http.StatusUnauthorized, "User Id not found")
		return
	}

	src.Seek(0, 0)

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to read file content")
		return
	}

	validatedFiles.Store(userIdFromToken, fileBytes)

	c.JSON(http.StatusOK, gin.H{
		"message": "File successfully uploaded and validated",
	})
}

func (h *Handler) PredictCSV(c *gin.Context) {
	userIdFromToken, exists := c.Get(userCtx)
	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	fileBytes, exists := validatedFiles.Load(userIdFromToken)
	if !exists {
		newErrorResponse(c, http.StatusBadRequest, "CSV file not found")
		return
	}

	src := bytes.NewReader(fileBytes.([]byte))

	var predictions map[int]*dto.PredictionResponseDto
	var err error

	if h.checkRole(c, constants.RoleInstructor) {
		predictions, err = h.services.PredictCSV(userIdFromToken.(int), src, true)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else if h.checkRole(c, constants.RoleStudent) {
		predictions, err = h.services.PredictCSV(userIdFromToken.(int), src, false)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
	} else {
		newErrorResponse(c, http.StatusForbidden, "you don't have access to this resource")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Prediction generated successfully",
		"prediction": predictions,
	})
}
