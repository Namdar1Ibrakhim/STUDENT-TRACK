package handler

import (
	"bytes"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
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

	err = h.services.ValidateCSV(src)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	src.Seek(0, 0)

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to read file content")
		return
	}

	userIdFromToken, exists := c.Get(userCtx)
	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
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

	studentIdParam := c.Query("student_id")
	var studentId int
	if h.checkRole(c, constants.RoleInstructor) || h.checkRole(c, constants.RoleAdmin) {
		if studentIdParam == "" {
			newErrorResponse(c, http.StatusInternalServerError, "missing student_id")
			return
		}
		studentId, _ = strconv.Atoi(studentIdParam)
	} else if h.checkRole(c, constants.RoleStudent) {
		studentId = userIdFromToken.(int)
		if studentIdParam != "" && strconv.Itoa(studentId) != studentIdParam {
			newErrorResponse(c, http.StatusForbidden, "you can upload only own dataset")
			return
		}
	} else {
		newErrorResponse(c, http.StatusForbidden, "you don't have access to this resource")
		return
	}

	prediction, err := h.services.PredictCSV(studentId, src)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Prediction generated successfully",
		"prediction": prediction,
	})
}
