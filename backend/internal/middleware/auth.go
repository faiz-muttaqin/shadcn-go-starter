package middleware

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// List of paths to ignore
		ignoredPaths := []string{
			"/auth/login", // Example: Skip login route
			// "/auth/register", // Example: Skip register route
			// "/public",        // Example: Skip public route
		}

		// Check if the current request path is in the ignored paths
		for _, path := range ignoredPaths {
			if c.Request.URL.Path == os.Getenv("VITE_API_BASE_URL")+path {
				// Skip middleware for this request and proceed
				c.Next()
				return
			}
		}

		// Cek cookie terlebih dahulu
		token := ""
		cookie, err := c.Cookie("accessToken") // Cek apakah cookie accessToken ada
		if err == nil {
			token = cookie
			fmt.Println("Cookie 'accessToken' found:", token)
		} else {
			// Jika cookie tidak ada, cek Bearer token di header Authorization
			authHeader := c.GetHeader("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token = authHeader[len("Bearer "):] // Ambil token setelah "Bearer "
				fmt.Println("Bearer Token found:", token)
			}
		}

		// Jika token ditemukan, lakukan validasi
		if token != "" {
			// Logika validasi token (misalnya, cek validitas JWT)
			fmt.Println("Token validated:", token)
			c.Next() // Lanjutkan ke handler berikutnya
		} else {
			// Tidak ada token atau token tidak valid
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort() // Hentikan pemrosesan lebih lanjut
		}
	}
}
