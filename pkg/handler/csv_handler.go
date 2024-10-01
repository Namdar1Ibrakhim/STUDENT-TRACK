package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
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

	err = h.services.ValidateCSV(src)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File successfully uploaded and validated",
	})
}
