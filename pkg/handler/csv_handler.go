package handler

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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

	userIdFromToken, exists := c.Get(userCtx)
	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	studentIdParam := c.Query("student_id")
	if h.checkRole(c, constants.RoleInstructor) || h.checkRole(c, constants.RoleAdmin) {
		if studentIdParam == "" {
			newErrorResponse(c, http.StatusInternalServerError, "missing student_id")
			return
		}
		studentId, err := strconv.Atoi(studentIdParam)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, "invalid student_id")
			return
		}

		err = h.services.ValidateCSV(src)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		src.Seek(0, 0)
		prediction, err := h.services.PredictCSV(studentId, src)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "File successfully uploaded and validated by Instructor",
			"prediction": prediction,
		})
		return

	}
	if h.checkRole(c, constants.RoleStudent) {
		if studentIdParam != "" && strconv.Itoa(userIdFromToken.(int)) != studentIdParam {
			newErrorResponse(c, http.StatusForbidden, "you can upload only own dataset")
			return
		}

		err = h.services.ValidateCSV(src)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		src.Seek(0, 0)
		prediction, err := h.services.PredictCSV(userIdFromToken.(int), src)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":    "File successfully uploaded and validated",
			"prediction": prediction,
		})
		return
	}
	newErrorResponse(c, http.StatusForbidden, "you don't have access to this resource")
}
