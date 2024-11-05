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

func (h *Handler) validateCSV(c *gin.Context, src io.Reader, role int) error {
	var err error
	switch role {
	case constants.RoleInstructor:
		err = h.services.ValidateCSVForInstructor(src)
	case constants.RoleStudent:
		err = h.services.ValidateCSVForStudent(src)
	default:
		return constants.ErrAccessDenied
	}
	return err
}

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

	userID, err := h.GetUserIDFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	role := h.getCurrentRole(c)
	if err := h.validateCSV(c, src, role); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if _, err := src.Seek(0, 0); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to process file")
		return
	}

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "Failed to read file content")
		return
	}

	validatedFiles.Store(userID, fileBytes)

	c.JSON(http.StatusOK, gin.H{
		"message": "File successfully uploaded and validated",
	})
}

func (h *Handler) PredictCSV(c *gin.Context) {
	userID, err := h.GetUserIDFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	fileBytes, exists := validatedFiles.Load(userID)
	if !exists {
		newErrorResponse(c, http.StatusBadRequest, constants.ErrFileNotFound.Error())
		return
	}

	role := h.getCurrentRole(c)
	isInstructor := role == constants.RoleInstructor

	var predictions map[int]*dto.PredictionResponseDto

	predictions, err = h.services.PredictCSV(userID, bytes.NewReader(fileBytes.([]byte)), isInstructor)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Prediction generated successfully",
		"prediction": predictions,
	})
}

func (h *Handler) getCurrentRole(c *gin.Context) int {
	if h.checkRole(c, constants.RoleInstructor) {
		return constants.RoleInstructor
	}
	if h.checkRole(c, constants.RoleStudent) {
		return constants.RoleStudent
	}
	return 0
}
