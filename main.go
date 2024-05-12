package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// login Structure
type Login struct {
	email string `json:"email" binding:"required"`
	pass  string `json:"pass" binding:"required"`
}

/*  helper  */

// make home page URL
func URL() string {
	return "http://127.0.0.1:8080"
}

// Validate user pass
func checkAuth(login Login) bool {
	dbuser := "admin@admin.co"
	dbpass := "0000"
	result := false
	if login.email == dbuser && login.pass == dbpass {
		result = true
	}
	return result

}

/*  helper  */

/* Middleware */

// Auth Middleware
func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		Cookie, err := c.Cookie("Auth")

		if Cookie == "logedin" && err == nil {

			c.Next()
		} else {

			c.Redirect(http.StatusMovedPermanently, "/auth/")
		}
	}
}

/* Middleware */
func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func main() {

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Defualt  Routes
	r.GET("/", authMiddleware(), func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, URL()+"/Dashboard")
	})
	v1 := r.Group("/auth")
	{
		//defult login page every user can see this page in open login URL
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "Main website",
			})
		})
		// login page after submit form
		v1.GET("/login", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, URL()+"/Dashboard")
		})
		// login page after submit form
		v1.POST("/login", func(c *gin.Context) {
			var login Login
			login.email = c.PostForm("email")
			login.pass = c.PostForm("pass")
			if checkAuth(login) {
				c.SetCookie("Auth", "logedin", 3600, "/Dashboard/", URL(), false, true)
				c.Redirect(http.StatusMovedPermanently, URL()+"/Dashboard")
			}

		})
	}
	// logins Route

	// Dashboard Route
	// Like (admin user page or Crud pages)
	v2 := r.Group("/Dashboard")
	{

		v2.GET("/", authMiddleware(), func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard.html", gin.H{
				"title": "Main website",
			})
		})

		v2.GET("/users", authMiddleware(), func(c *gin.Context) {
			c.HTML(http.StatusOK, "users.html", gin.H{
				"title":   "Main website",
				"user_id": 1,
			})
		})
		v2.GET("/edituser", func(c *gin.Context) {

			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"title":   "Main website",
				"user_id": 1,
			})
		})

	}
	// Dashboard Route

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
