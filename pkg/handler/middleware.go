package handler

import (
	"github.com/Namdar1Ibrakhim/student-track-system/pkg/constants"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	bearerSchema        = "Bearer"
)

func extractTokenFromHeader(c *gin.Context) (string, error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", constants.ErrEmptyAuthHeader
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || !strings.EqualFold(headerParts[0], bearerSchema) {
		return "", constants.ErrInvalidAuthHeader
	}

	return headerParts[1], nil
}

func (h *Handler) userIdentity(c *gin.Context) {
	token, err := extractTokenFromHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	userId, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId)
	c.Next()
}

func (h *Handler) getUserIdentity(c *gin.Context) int {
	token, err := extractTokenFromHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	userId, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return 0
	}
	return userId
}

func (h *Handler) checkRole(c *gin.Context, requiredRole int) bool {
	token, err := extractTokenFromHeader(c)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return false
	}

	userId, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return false
	}

	user, err := h.services.GetUser(userId)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		c.Abort()
		return false
	}

	if user.Role != requiredRole {
		//logrus.WithFields(logrus.Fields{
		//	"userRole":     user.Role,
		//	"requiredRole": requiredRole,
		//}).Info("Access denied due to role mismatch")
		//newErrorResponse(c, http.StatusForbidden, "you don't have access to this resource")
		c.Abort()
		return false
	}

	c.Set(userCtx, userId)
	c.Next()
	return true
}

func (h *Handler) GetUserIDFromContext(c *gin.Context) (int, error) {
	userID, exists := c.Get(userCtx)
	if !exists {
		return 0, constants.ErrUserNotFound
	}
	return userID.(int), nil
}
