package middleware

import (
	"net/http"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
	"github.com/gin-gonic/gin"
)

// func AuthClerk() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// Use Clerk middleware for HTTP handler
// 		handler := clerkhttp.RequireHeaderAuthorization()(
// 			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 				// Kalau masuk sini => token valid
// 				// Set back r into gin.Context supaya bisa lanjut ke handler Gin
// 				c.Request = r
// 				c.Next()
// 			}),
// 		)

//			handler.ServeHTTP(c.Writer, c.Request)
//			// If Clerk middleware has already rejected, we stop the chain
//			if c.IsAborted() {
//				return
//			}
//		}
//	}
func AuthClerk() gin.HandlerFunc {
	// Clerk middleware (net/http)
	clerkHandler := clerkhttp.RequireHeaderAuthorization()(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// no-op, Clerk hanya memasukkan context
	}))

	// Wrap Clerk middleware into Gin middleware
	return gin.WrapH(clerkHandler)
}
