package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
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

	c.Set(userCtx, userId)
}

func (h *Handler) checkRole(c *gin.Context, requiredRole int) bool {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		c.Abort()
		return false
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		c.Abort()
		return false
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		c.Abort()
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
