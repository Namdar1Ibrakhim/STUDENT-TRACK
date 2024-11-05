package handler

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/dto"
	"net/http"
	"strconv"

	track "github.com/Namdar1Ibrakhim/student-track-system"
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"github.com/gin-gonic/gin"
)

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
	userId, err := h.GetUserIDFromContext(c)
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

func (h *Handler) getAllUsers(c *gin.Context) {
	_, err := h.GetUserIDFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if !h.checkRole(c, constants.RoleAdmin) {
		newErrorResponse(c, http.StatusForbidden, constants.ErrAccessDenied.Error())
		return
	}

	users, err := h.services.GetAllUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"users": users,
	})
}

func (h *Handler) getStudentById(c *gin.Context) {

	h.checkRole(c, constants.RoleInstructor)
	if c.IsAborted() {
		newErrorResponse(c, http.StatusForbidden, constants.ErrAccessDenied.Error())
		return
	}

	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.ErrInvalidStudentId.Error())
		return
	}
	user, err := h.services.GetUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	if user.Role != constants.RoleStudent {
		newErrorResponse(c, http.StatusUnauthorized, constants.ErrStudentNotFound.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"user": user,
	})
}

func (h *Handler) getStudents(c *gin.Context) {
	if h.checkRole(c, constants.RoleInstructor) {
		if c.IsAborted() {
			newErrorResponse(c, http.StatusForbidden, constants.ErrAccessDenied.Error())
			return
		}
	}

	students, err := h.services.GetStudents()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"students": students,
	})
}

func (h *Handler) getUserById(c *gin.Context) {

	userId := c.Param("id")
	id, err := strconv.Atoi(userId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.ErrInvalidStudentId.Error())
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
	var userIdFromPath int
	idParam := c.Param("id")
	if idParam != "" {
		var err error
		userIdFromPath, err = strconv.Atoi(idParam)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, constants.ErrInvalidUserId.Error())
			return
		}
	}

	userIdFromContext, err := h.GetUserIDFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input dto.UpdateUser
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.ErrInvalidInputData.Error())
		return
	}

	if idParam == "" || userIdFromContext == userIdFromPath {
		existingUser, err := h.services.GetUser(userIdFromContext)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, constants.ErrUserNotFound.Error())
			return
		}

		if err := h.services.UpdateUser(existingUser.Id, input); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Profile updated",
		})
	} else {
		h.checkRole(c, constants.RoleAdmin)
		if c.IsAborted() {
			newErrorResponse(c, http.StatusForbidden, constants.ErrAccessDenied.Error())
			return
		}

		existingUser, err := h.services.GetUser(userIdFromPath)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, constants.ErrUserNotFound.Error())
			return
		}

		if err := h.services.UpdateUser(existingUser.Id, input); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User updated by admin",
		})
	}
}

func (h *Handler) DeleteUser(c *gin.Context) {
	var userIdFromPath int
	idParam := c.Param("id")
	if idParam != "" {
		var err error
		userIdFromPath, err = strconv.Atoi(idParam)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, constants.ErrInvalidUserId.Error())
			return
		}
	}

	userIdFromContext, err := h.GetUserIDFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if idParam == "" || userIdFromContext == userIdFromPath {
		existingUser, err := h.services.GetUser(userIdFromContext)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, constants.ErrUserNotFound.Error())
			return
		}

		if err := h.services.DeleteUser(existingUser.Id); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Account deleted",
		})
	} else {
		h.checkRole(c, constants.RoleAdmin)
		if c.IsAborted() {
			newErrorResponse(c, http.StatusForbidden, constants.ErrAccessDenied.Error())
			return
		}

		existingUser, err := h.services.GetUser(userIdFromPath)
		if err != nil {
			newErrorResponse(c, http.StatusNotFound, constants.ErrUserNotFound.Error())
			return
		}

		if err := h.services.DeleteUser(existingUser.Id); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"message": "User deleted by admin",
		})
	}
}

func (h *Handler) editPasswordByAdmin(c *gin.Context) {
	if !h.checkRole(c, constants.RoleAdmin) {
		newErrorResponse(c, http.StatusForbidden, constants.ErrAccessDenied.Error())
		return
	}

	idParam := c.Param("id")
	userId, err := strconv.Atoi(idParam)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.ErrInvalidUserId.Error())
		return
	}

	var input dto.EditPasswordAdminRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.ErrInvalidInputData.Error())
		return
	}

	if err := h.services.EditPassword(userId, "", input.NewPassword, true); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Password updated by admin",
	})
}

func (h *Handler) editPasswordByUser(c *gin.Context) {
	userId, err := h.GetUserIDFromContext(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input dto.EditPasswordUserRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, constants.ErrInvalidInputData.Error())
		return
	}

	if err := h.services.EditPassword(userId, input.OldPassword, input.NewPassword, false); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Password updated",
	})
}
