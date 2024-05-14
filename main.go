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
	// var users Users
	var user = boot.Users{}
	// boot.DB().Where("email = ?", "hosseinbidar7@gmail.com").First(&user)
	boot.DB().Where("email = ?", "hosseinbidar7@gmail.com").First(&user)
	var users []boot.Users
	boot.DB().Model(&boot.Users{}).Select("*").Where("Email = ?", "hosseinbidar7@gmail.com").Scan(&users)
	pass := ""
	for _, user := range users {
		pass = user.Password
	}
	fmt.Println(pass)

	result := boot.DB().Migrator().CreateTable(boot.Users{})
	result2 := boot.DB().Migrator().CreateTable(boot.Inventory{})

	// result := db.Create(&Users)
	if result == nil {
		a, _ := utility.HashPassword("0000")
		User := boot.Users{Name: "hossein Soltanian", Email: "hosseinbidar7@gmail.com", Password: a, Role: "Admin", Phonenumber: "09125174854"}
		boot.DB().Create(&User)
	}
	if result2 == nil {
		Inventory := boot.Inventory{Name: "ایزوگام شرق", Number: "10", RolePrice: 99.250, MeterPrice: 102.500, Count: 100, InventoryNumber: 1}
		boot.DB().Create(&Inventory)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Defualt  Routes
	r.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, utility.HomeUrl()+"/Dashboard")
	})

	// auth Route
	// Like (admin user page or Crud pages)
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
				"title":  "Main website",
				"action": "add_user",
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
				"title":  "Main website",
				"user":   currentusrt,
				"action": "edituser",
			})
		})
		v2.POST("/edituser", func(c *gin.Context) {
			currentusrt := utility.GetCurrentUser(c)
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"title": "Main website",
				"user":  currentusrt,
			})
		})

		// Product
		v2.GET("/addproduct", func(c *gin.Context) {

			c.HTML(http.StatusOK, "add_product.html", gin.H{
				"title":  "Main website",
				"action": "addproduct",
			})
		})
		v2.POST("/addproduct", middleware.AuthMiddleware(), func(c *gin.Context) {
			var product boot.Inventory
			product.Name = c.PostForm("Name")
			product.Number = c.PostForm("Number")
			product.RolePrice = utility.StringToFloat(c.PostForm("RolePrice"))
			product.MeterPrice = utility.StringToFloat(c.PostForm("MeterPrice"))
			product.Count = utility.StringToInt(c.PostForm("Count"))
			product.InventoryNumber = utility.GetCurrentInventory(c)
			boot.DB().Create(&product)
			fmt.Println(product)
			c.Redirect(http.StatusMovedPermanently, utility.HomeUrl()+"/Dashboard/inventory")

		})
		// inventory
		v2.GET("/inventory", func(c *gin.Context) {

			c.HTML(http.StatusOK, "inventory.html", gin.H{
				"title":    "Main website",
				"products": model.GetAllProductsByInventory(1),
			})
		})

	}
	// Dashboard Route

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
