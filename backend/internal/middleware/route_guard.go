package middleware

import (
	"github.com/gin-gonic/gin"
)

func RouteGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Implement route guarding logic here
		// userData, err := helper.ClerkGetUser(c.Request.Context())
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{
		// 		"success": false,
		// 		"error":   err.Error(),
		// 	})
		// 	return
		// }

		c.Next()
	}
}
