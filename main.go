package main

import (
	"fmt"
	auth "inventory/app/Auth"
	model "inventory/app/Model"
	"inventory/app/middleware"
	utility "inventory/app/utility"
	"inventory/boot"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {

	// fmt.Println("pass is:", CheckPasswordHash("!HS0311121314", a), "pass \n")

	// var users Users
	var user = boot.Users{}
	// boot.DB().Where("email = ?", "hosseinbidar7@gmail.com").First(&user)
	boot.DB().Where("email = ?", "hosseinbidar7@gmail.com").First(&user)
	// rows, _ := boot.DB().Model(&boot.Users{}).Rows()
	// defer rows.Close()
	// fmt.Println(rows.Next())
	var users []boot.Users
	boot.DB().Model(&boot.Users{}).Select("*").Where("Email = ?", "hosseinbidar7@gmail.com").Scan(&users)
	pass := ""
	for _, user := range users {
		pass = user.Password
	}
	fmt.Println(pass)

	result := boot.DB().Migrator().CreateTable(boot.Users{})

	// result := db.Create(&Users)
	if result == nil {
		a, _ := utility.HashPassword("0000")
		User := boot.Users{Name: "hossein Soltanian", Email: "hosseinbidar7@gmail.com", Password: a, Role: "Admin", Phonenumber: "09125174854"}
		boot.DB().Create(&User)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

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
			var login boot.Login
			login.Email = c.PostForm("email")
			login.Password = c.PostForm("pass")
			authorized, name := auth.CheckAuth(login)
			if authorized {
				session := sessions.Default(c)

				if session.Get("Auth") != "logedin" {
					session.Set("Auth", "logedin")
					session.Set("UserName", &name)
					session.Save()
				}
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
			session := sessions.Default(c)
			fmt.Println("session:", session.Get("UserName"))
			c.HTML(http.StatusOK, "dashboard.html", gin.H{
				"title":    "Main website",
				"Username": session.Get("UserName"),
			})
		})

		v2.GET("/users", middleware.AuthMiddleware(), func(c *gin.Context) {

			// result := "a"
			c.HTML(http.StatusOK, "users.html", gin.H{
				"title": "Main website",
				"users": model.GetAllUsersByRole("guest"),
			})
		})
		v2.GET("/admin_users", middleware.AuthMiddleware(), func(c *gin.Context) {
			c.HTML(http.StatusOK, "admins.html", gin.H{
				"title": "Main website",
				"users": model.GetAllUsersByRole("Admin"),
			})
		})
		v2.GET("/add_user", middleware.AuthMiddleware(), func(c *gin.Context) {
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"title":   "Main website",
				"user_id": 1,
			})
		})
		v2.POST("/add_user", middleware.AuthMiddleware(), func(c *gin.Context) {
			var user boot.Users
			user.Name = c.PostForm("Name")
			user.Email = c.PostForm("Email")
			user.Phonenumber = c.PostForm("Phonenumber")
			user.Password = c.PostForm("Password")
			user.Role = c.PostForm("Role")
			user.Password, _ = utility.HashPassword(user.Password)
			boot.DB().Create(&user)
		})
		v2.GET("/edituser", func(c *gin.Context) {
			currentusrt := utility.GetCurrentUser(c)
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"title": "Main website",
				"user":  currentusrt,
			})
		})

	}
	// Dashboard Route

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
