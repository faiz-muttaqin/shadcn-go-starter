package handler

import (
	"net/http"
	"os"

	"github.com/faiz-muttaqin/shadcn-admin-go-starter/backend/internal/helper"
	"github.com/gin-gonic/gin"
)

// VerifyAuth verifies if the user is authenticated with valid token
func VerifyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userData, err := helper.GetFirebaseUser(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user_id": userData.ExternalID,
			"email":   userData.Email,
			"user":    userData,
			"success": true,
			"message": "Authentication verified successfully",
			"data": gin.H{
				"authenticated": true,
				"tokenLength":   1000,
				"tokenPreview":  "12345678901234567890" + "...",
				"environment":   os.Getenv("GIN_MODE"),
			},
		})
	}
}
