package main

import (
	"fmt"
	"inventory/app/middleware"
	utility "inventory/app/utility"
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

// make home page utility.HomeUrl

// Validate user pass
func db() *gorm.DB {

	dsn := "root:0311121314@tcp(127.0.0.1:3306)/Inventory?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db

}

// get all user with rol [Admin , guest]

func GetAllUsersByRole() []Users {
	var users []Users
	db().Model(&Users{}).Select("*").Scan(&users)
	return users
}

// get user By email [Admin , guest]

func GetUserByEmail(login Login) []Users {
	var users []Users
	db().Model(&Users{}).Select("*").Where("Email = ?", "hosseinbidar7@gmail.com").Scan(&users)
	return users
}

// get user By email [Admin , guest]

func checkAuth(login Login) bool {
	pass := ""
	for _, user := range GetUserByEmail(login) {
		pass = user.Password
	}
	dbpass := login.pass
	result := false
	if CheckPasswordHash(dbpass, pass) {
		result = true
	}
	return result

}

/*  helper  */

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

	// fmt.Println("pass is:", CheckPasswordHash("!HS0311121314", a), "pass \n")

	// var users Users
	var user = Users{}
	// db().Where("email = ?", "hosseinbidar7@gmail.com").First(&user)
	db().Where("email = ?", "hosseinbidar7@gmail.com").First(&user)
	// rows, _ := db().Model(&Users{}).Rows()
	// defer rows.Close()
	// fmt.Println(rows.Next())
	var users []Users
	db().Model(&Users{}).Select("*").Where("Email = ?", "hosseinbidar7@gmail.com").Scan(&users)
	pass := ""
	for _, user := range users {
		pass = user.Password
	}
	fmt.Println(pass)

	result := db().Migrator().CreateTable(Users{})

	// result := db.Create(&Users)
	if result == nil {
		a, _ := HashPassword("0000")
		User := Users{Name: "hossein Soltanian", Email: "hosseinbidar7@gmail.com", Password: a, Role: "Admin", Phonenumber: "09125174854"}
		db().Create(&User)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Defualt  Routes
	r.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, utility.HomeUrl()+"/Dashboard")
	})
	v1 := r.Group("/auth")
	{
		//defult login page every user can see this page in open login utility.HomeUrl
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "Main website",
			})
		})
		// login page after submit form
		v1.GET("/login", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, utility.HomeUrl()+"/Dashboard")
		})
		// login page after submit form
		v1.POST("/login", func(c *gin.Context) {
			var login Login
			login.email = c.PostForm("email")
			login.pass = c.PostForm("pass")
			if checkAuth(login) {
				c.SetCookie("Auth", "logedin", 3600, "/Dashboard/", utility.HomeUrl(), false, true)
				c.Redirect(http.StatusMovedPermanently, utility.HomeUrl()+"/Dashboard")
			}

		})
	}
	// logins Route

	// Dashboard Route
	// Like (admin user page or Crud pages)
	v2 := r.Group("/Dashboard")
	{

		v2.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
			c.HTML(http.StatusOK, "dashboard.html", gin.H{
				"title": "Main website",
			})
		})

		v2.GET("/users", middleware.AuthMiddleware(), func(c *gin.Context) {

			// result := "a"
			c.HTML(http.StatusOK, "users.html", gin.H{
				"title": "Main website",
				"users": GetAllUsersByRole(),
			})
		})
		v2.GET("/admin_users", middleware.AuthMiddleware(), func(c *gin.Context) {
			c.HTML(http.StatusOK, "admins.html", gin.H{
				"title":   "Main website",
				"user_id": 1,
			})
		})
		v2.GET("/add_user", middleware.AuthMiddleware(), func(c *gin.Context) {
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"title":   "Main website",
				"user_id": 1,
			})
		})
		v2.POST("/add_user", middleware.AuthMiddleware(), func(c *gin.Context) {
			var user Users
			user.Name = c.PostForm("Name")
			user.Email = c.PostForm("Email")
			user.Phonenumber = c.PostForm("Phonenumber")
			user.Password = c.PostForm("Password")
			user.Role = c.PostForm("Role")
			user.Password, _ = HashPassword(user.Password)
			db().Create(&user)
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
