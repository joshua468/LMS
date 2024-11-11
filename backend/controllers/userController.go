package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joshua468/lms-backend/services"
)

func GetUserProfile(c *gin.Context) {
	userID := c.Param("userID")
	user, err := services.GetUserProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}
