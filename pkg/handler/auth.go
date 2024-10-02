package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"

	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/gin-gonic/gin"
)

// SIGN UP FOR STUDENT
func (h *Handler) signUpStudent(c *gin.Context) {
	var input track.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input, constants.RoleStudent)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// SIGN UP FOR INSTRUCTOR
func (h *Handler) signUpInstructor(c *gin.Context) {
	var input track.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input, constants.RoleInstructor)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// SIGN UP FOR ADMIN
func (h *Handler) signUpAdmin(c *gin.Context) {
	var input track.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input, constants.RoleAdmin)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}

func (h *Handler) getUser(c *gin.Context) {

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	user, err := h.services.GetUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (h *Handler) getStudentById(c *gin.Context) {

	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid student ID")
		return
	}
	user, err := h.services.GetUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	if user.Role != constants.RoleStudent {
		newErrorResponse(c, http.StatusUnauthorized, "Student not found with this id")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (h *Handler) getUserById(c *gin.Context) {

	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid student ID")
		return
	}
	user, err := h.services.GetUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	idParam := c.Param("id")
	userIdFromPath, err := strconv.Atoi(idParam)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	userIdFromToken, exists := c.Get(userCtx)
	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	userIdFromTokenInt, ok := userIdFromToken.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid user id type")
		return
	}

	var input dto.UpdateUser
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid input data")
	}

	if userIdFromTokenInt != userIdFromPath {
		h.checkRole(c, constants.RoleAdmin) //checking permisson
		if c.IsAborted() {
			newErrorResponse(c, http.StatusForbidden, "you don't have access to this resource")
			return
		}

		existingUser, err := h.services.GetUser(userIdFromPath)

		if err != nil {
			newErrorResponse(c, http.StatusNotFound, "User not found")
			return
		}
		if err := h.services.UpdateUser(existingUser.Id, input); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User updated by admin",
		})

		return
	}

	existingUser, err := h.services.GetUser(userIdFromTokenInt)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "User not found")
		return
	}

	if err := h.services.UpdateUser(existingUser.Id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Profile updated",
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	userIdFromPath, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	userIdFromToken, exists := c.Get(userCtx)
	if !exists {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return
	}

	userIdFromTokenInt, ok := userIdFromToken.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "Invalid user id type")
		return
	}

	if userIdFromTokenInt != userIdFromPath {
		h.checkRole(c, constants.RoleAdmin)
		if c.IsAborted() {
			newErrorResponse(c, http.StatusForbidden, "you don't have access to this resource")
			return
		}

		existingUser, err := h.services.GetUser(userIdFromPath)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, err.Error())
			return
		}

		if err := h.services.DeleteUser(existingUser.Id); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User deleted by admin",
		})
		return
	}

	existingUser, err := h.services.GetUser(userIdFromTokenInt)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, "user not found")
		return
	}

	if err := h.services.DeleteUser(existingUser.Id); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Account deleted",
	})
}

func (h *Handler) editPasswordByCurrentUserId(c *gin.Context) {
	passwordParam := c.Param("password")

	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if err := h.services.EditPassword(userId, passwordParam); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Password Updated",
	})
}

func (h *Handler) editPasswordByUserId(c *gin.Context) {
	idParam := c.Param("id")
	passwordParam := c.Param("password")

	userId, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
		return
	}

	if err := h.services.EditPassword(userId, passwordParam); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Password Updated",
	})
}
