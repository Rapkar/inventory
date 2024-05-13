package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/* Middleware */

// Auth Middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Cookie, err := c.Cookie("Auth")
		session := sessions.Default(c)
		session.Get("Auth")
		if session.Get("Auth") == "logedin" {

			c.Next()
		} else {

			c.Redirect(http.StatusMovedPermanently, "/auth/")
		}
		// if Cookie == "logedin" && err == nil {

		// 	c.Next()
		// } else {

		// 	c.Redirect(http.StatusMovedPermanently, "/auth/")
		// }
	}
}

/* Middleware */
