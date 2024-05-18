package main

import (
	"fmt"
	auth "inventory/app/Auth"
	model "inventory/app/Model"
	"inventory/app/middleware"
	utility "inventory/app/utility"
	"inventory/boot"
	"log"
	"net/http"
	"strconv"

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
	result3 := boot.DB().Migrator().CreateTable(boot.ExportProducts{})
	result4 := boot.DB().Migrator().CreateTable(boot.Export{})

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
	ExportProduct := []boot.ExportProducts{}
	if result3 == nil {
		ExportProduct = []boot.ExportProducts{{Name: "ایزوگام شرق", Number: "10", RolePrice: 99.250, MeterPrice: 102.500, Count: 100, InventoryNumber: 1, TotalPrice: 2000000, Meter: 10}}
		boot.DB().Create(&ExportProduct)
	}
	if result4 == nil {
		Export := boot.Export{Name: "رضا توانگر", Number: "9283422", Phonenumber: "09199656725", Address: "کرج -کرج=-ایران -سیسی", TotalPrice: 10000000, Tax: 10, ExportProducts: ExportProduct, InventoryNumber: 1, CreatedAt: utility.CurrentTime()}
		boot.DB().Create(&Export)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("assets", "./assets")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Defualt  Routes
	r.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, utility.HomeUrl()+"/Dashboard")
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"title": "Main website",
		})
	})

	// r.GET("assets/images/", func(c *gin.Context) {
	// 	c.File("./assets/images" + c.Param("images"))
	// 	fmt.Println("imageeeeeeee", c.Param("images"))
	// })

	// auth Route
	// Like (admin user page or Crud pages)
	v1 := r.Group("/auth")
	{
		//defult login page every user can see this page in open login utility.HomeUrl
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "ورود به حساب",
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
					session.Set("UserName", name)
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
			c.HTML(http.StatusOK, "dashboard.html", gin.H{
				"title":    "صفحه اصلی",
				"Username": session.Get("UserName"),
				"message":  boot.Messages("login success"),
				"success":  true,
			})
		})

		v2.GET("/users", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			// result := "a"
			c.HTML(http.StatusOK, "users.html", gin.H{

				"Username": session.Get("UserName"),
				"title":    "کاربران",
				"users":    model.GetAllUsersByRole("guest"),
			})
		})
		v2.GET("/admin_users", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "admins.html", gin.H{
				"Username": session.Get("UserName"),
				"title":    "کاربران ادمین",
				"users":    model.GetAllUsersByRole("Admin"),
			})
		})
		v2.GET("/add_user", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"Username": session.Get("UserName"),
				"title":    "کاربران",
				"action":   "add_user",
			})
		})
		v2.POST("/add_user", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			var user boot.Users
			user.Name = c.PostForm("Name")
			user.Email = c.PostForm("Email")
			user.Phonenumber = c.PostForm("Phonenumber")
			user.Password = c.PostForm("Password")
			user.Role = c.PostForm("Role")
			user.Password, _ = utility.HashPassword(user.Password)
			res := boot.DB().Create(&user)
			if res.RowsAffected > 0 {
				c.HTML(http.StatusOK, "edit_user.html", gin.H{
					"Username": session.Get("UserName"),
					"title":    "کاربران",
					"action":   "add_user",
					"message":  boot.Messages("user made success"),
					"success":  true,
				})
			} else {
				c.HTML(http.StatusOK, "edit_user.html", gin.H{
					"title":   "کاربران",
					"action":  "add_user",
					"message": boot.Messages("user made faild"),
					"success": false,
				})
			}

		})
		v2.GET("/edituser", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			currentusrt := utility.GetCurrentUser(c)
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"Username": session.Get("UserName"),
				"title":    "ویرایش کاربر",
				"user":     currentusrt,
				"action":   "edituser",
			})
		})
		v2.POST("/edituser", middleware.AuthMiddleware(), func(c *gin.Context) {
			currentusrt := utility.GetCurrentUser(c)
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"title":   "ویرایش کاربر",
				"action":  "add_user",
				"message": "کاربر با موفقیت  اصلاح شد",
				"success": true,
				"user":    currentusrt,
			})
		})

		// Product
		v2.GET("/addproduct", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "add_product.html", gin.H{
				"Username":        session.Get("UserName"),
				"title":           "محصول",
				"action":          "addproduct",
				"InventoryNumber": utility.GetCurrentInventory(c),
			})
		})
		v2.POST("/addproduct", middleware.AuthMiddleware(), func(c *gin.Context) {
			var product boot.Inventory
			product.Name = c.PostForm("Name")
			product.Number = c.PostForm("Number")
			product.RolePrice = utility.StringToFloat(c.PostForm("RolePrice"))
			product.MeterPrice = utility.StringToFloat(c.PostForm("MeterPrice"))
			product.Count = utility.StringToInt(c.PostForm("Count"))
			product.InventoryNumber = utility.StringToInt32(c.PostForm("InventoryNumber"))
			res := boot.DB().Create(&product)
			if res.RowsAffected > 0 {
				c.HTML(http.StatusOK, "add_product.html", gin.H{
					"title":   "محصول",
					"action":  "addproduct",
					"message": boot.Messages("product made success"),
					"success": true,
				})
			} else {
				c.HTML(http.StatusOK, "add_product.html", gin.H{
					"title":   "محصول",
					"action":  "addproduct",
					"message": boot.Messages("product made faild"),
					"success": false,
				})
			}

		})
		// inventory
		v2.GET("/inventory", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "inventory.html", gin.H{
				"Username": session.Get("UserName"),
				"title":    "انبار",
				"products": model.GetAllProductsByInventory(utility.GetCurrentInventory(c)),
			})
		})

		v2.GET("/export", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "export.html", gin.H{
				"Username": session.Get("UserName"),
				"action":   "export",
				"title":    "فاکتور",
				"products": model.GetAllProductsByInventory(1),
			})
		})
		v2.POST("/getproductbyinventory", func(c *gin.Context) {
			var data struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			}
			if err := c.BindJSON(&data); err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			val, _ := strconv.ParseInt(data.Id, 10, 32)

			products := model.GetAllProductsByInventory(int32(val))
			c.JSON(http.StatusOK, gin.H{"result": products})
		})
		v2.POST("/getproductbyid", func(c *gin.Context) {

			var data struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			}
			if err := c.BindJSON(&data); err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			val, _ := strconv.ParseInt(data.Id, 10, 8)
			product := model.GetProductById(int(val))
			c.JSON(http.StatusOK, gin.H{"result": product})
		})
		v2.POST("/export", func(c *gin.Context) {
			// mmm := map[string]string
			type Product struct {
				ID              string `gorm:"primaryKey"`
				ExportID        string `gorm:"size:255;"`
				Name            string `gorm:"type:varchar(100)" json:"name"`
				Number          string `gorm:"size:255;"`
				RolePrice       string `gorm:"type:float"`
				MeterPrice      string `gorm:"type:float"`
				Count           string `gorm:"size:255;"`
				Meter           string `gorm:"size:255;"`
				TotalPrice      string `gorm:"size:255;"`
				InventoryNumber string `gorm:"size:255;"`
			}

			var data struct {
				Name     string    `json:"Name"`
				Content  string    `json:"Content"`
				Products []Product `json:"Products"`
			}
			if err := c.BindJSON(&data); err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			exportproducts := make([]boot.ExportProducts, len(data.Products))
			// fmt.Println(data)
			for a, _ := range data.Products {
				exportproducts[a].ExportID, _ = strconv.ParseUint(data.Products[a].ExportID, 10, 64)
				exportproducts[a].Name = data.Products[a].Name
				exportproducts[a].Number = data.Products[a].ExportID
				exportproducts[a].RolePrice, _ = strconv.ParseFloat(data.Products[a].RolePrice, 64)
				exportproducts[a].MeterPrice, _ = strconv.ParseFloat(data.Products[a].MeterPrice, 64)
				Count, _ := strconv.ParseInt(data.Products[a].Count, 10, 8)
				exportproducts[a].Count = int8(Count)
				InventoryNumber, _ := strconv.ParseInt(data.Products[a].InventoryNumber, 10, 32)
				exportproducts[a].InventoryNumber = int32(InventoryNumber)

			}
			result۲ := utility.Unserialize(data.Content)
			fmt.Println("data", result۲, exportproducts)

			// byteSlice, err := json.Marshal(data.Products)
			// if err != nil {
			// 	panic(err)
			// }
			// // fmt.Println(reflect.TypeOf(data.Products), data.Products, byteSlice)
			// var products []boot.ExportProducts
			// json(data.Products)
			// fmt.Println(reflect.TypeOf(data.Products), data.Products)
			// for va, in := range data.Products {
			// 	fmt.Println(in, va)
			// }

			// fmt.Println(re, result۲)
			// Export:=[]boot.Export{
			// 	Name:result[""]
			// }

			// fmt.Println(html.EscapeString(result["Name"]))

			// datas["Address"]
			c.JSON(http.StatusOK, gin.H{"message": "data"})
		})
		v2.GET("/export-list", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			// Export:=model.GetAllExports()
			// Export.TotalPrice=model.FloatToString(Export.TotalPrice)
			fmt.Println(model.GetAllExports())
			c.HTML(http.StatusOK, "export_list.html", gin.H{
				"Username": session.Get("UserName"),
				"title":    "فاکتورها",
				"exports":  model.GetAllExports(),
			})
		})

	}
	// Dashboard Route

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
