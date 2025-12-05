package middleware

import (
	"github.com/gin-gonic/gin"
)

func Security() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Security-Policy", "frame-ancestors 'none'")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("Strict-Transport-Security", "max-age=16070400; includeSubDomains")
		c.Header("Referrer-Policy", "no-referrer")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Next()
	}
}
