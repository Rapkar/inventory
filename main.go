package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// login Structure
type Login struct {
	email string `json:"email" binding:"required"`
	pass  string `json:"pass" binding:"required"`
}
type Users struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;index:idx_name,unique"`
	Email       string `gorm:"size:255;"`
	Password    string `gorm:"type:varchar(255)"`
	Phonenumber string `gorm:"size:255;"`
	Role        string `gorm:"size:255;"`
}

/*  helper  */

// make home page URL
func URL() string {
	return "http://127.0.0.1:8080"
}

// Validate user pass
func checkAuth(login Login) bool {
	a, _ := HashPassword("0000")
	dbpass := login.pass
	result := false
	fmt.Println("sssssssssssssssssssssssssssssss")
	if CheckPasswordHash(dbpass, a) {
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

// password hashing

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {

	dsn := "root:0311121314@tcp(127.0.0.1:3306)/Inventory?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// fmt.Println("pass is:", CheckPasswordHash("!HS0311121314", a), "pass \n")
	if err == nil {
		fmt.Print("Connection success : ", db)

		result := db.Migrator().CreateTable(Users{})

		// result := db.Create(&Users)
		if result == nil {
			a, _ := HashPassword("0000")
			User := Users{Name: "hossein Soltanian", Email: "hosseinbidar7@gmail.com", Password: a, Role: "Admin", Phonenumber: "09125174854"}
			db.Create(&User)
		}
	} else {
		fmt.Println(err)
	}

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
