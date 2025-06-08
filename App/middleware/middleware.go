package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/* Middleware */

// Auth Middleware
func AuthMiddleware(role ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		auth := session.Get("Auth")
		userRole, _ := session.Get("UserRole").(string)

		// اگر لاگین نبود، redirect کن
		if auth != "logedin" {
			c.Redirect(http.StatusFound, "/auth/")
			c.Abort()
			return
		}

		// اگر role مشخص شده بود، چک کن که کاربر مجاز است
		if len(role) > 0 {
			allowedRoles := map[string]bool{
				"Admin":  true,
				"Author": true,
			}

			// اگر نقش کاربر مجاز نبود، redirect کن
			if !allowedRoles[userRole] || (role[0] != "" && userRole != role[0]) {
				c.Redirect(http.StatusFound, "/auth/")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

/* Middleware */
