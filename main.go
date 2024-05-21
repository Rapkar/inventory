package main

import (
	"fmt"
	"html/template"
	auth "inventory/App/Auth"
	"inventory/App/Boot"
	boot "inventory/App/Boot"
	controller "inventory/App/Controller"
	model "inventory/App/Model"
	"inventory/App/Utility"
	"inventory/App/middleware"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	const postperpage int = 100
	boot.Init()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("assets", "./assets")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Defualt  Routes
	r.GET("/", middleware.AuthMiddleware(), func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, Utility.HomeUrl()+"/Dashboard")
	})
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404.html", gin.H{
			"title": "Main website",
		})
	})

	// auth Route
	// Like (admin user page or Crud pages)
	v1 := r.Group("/auth")
	{
		//defult login page every user can see this page in open login Utility.HomeUrl
		v1.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "login.html", gin.H{
				"title": "ورود به حساب",
			})
		})
		// login page after submit form
		v1.GET("/login", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, Utility.HomeUrl()+"/Dashboard")
		})
		// login page after submit form
		v1.POST("/login", func(c *gin.Context) {
			var login boot.Login
			login.Email = c.PostForm("email")
			login.Password = c.PostForm("pass")
			authorized, name := auth.CheckAuth(login)
			// fmt.Print(authorized, c.PostForm("email"), c.b, authorized)
			if authorized {
				session := sessions.Default(c)
				if session.Get("Auth") != "logedin" {
					session.Set("Auth", "logedin")
					session.Set("UserName", name)
					session.Save()
				}
				c.SetCookie("email", login.Email, 3600, "/Dashboard/", Utility.HomeUrl(), false, true)
				c.SetCookie("pass", login.Password, 3600, "/Dashboard/", Utility.HomeUrl(), false, true)

				// c.JSON(http.StatusOK, gin.H{"message": "success"})
				c.Redirect(http.StatusMovedPermanently, Utility.HomeUrl()+"/Dashboard")
			} else {

				c.Redirect(http.StatusMovedPermanently, Utility.HomeUrl()+"/auth")
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
				"users":    model.GetAllUsersByRole("guest"),
				"exports":  model.GetAllExportsByPaginate(0, 5),
			})
		})

		v2.GET("/signout", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			session.Delete("Auth")
			c.Redirect(http.StatusMovedPermanently, Utility.HomeUrl()+"/auth")

		})
		v2.GET("/users", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			// result := "a"
			c.HTML(http.StatusOK, "users.html", gin.H{

				"Username": session.Get("UserName"),
				"title":    "کاربران",
				"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfUsers()/1, "user-list")),
				"users":    model.GetAllUsersByPaginate(0, postperpage),
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
			user.Password, _ = Utility.HashPassword(user.Password)
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
		v2.POST("/user-list", middleware.AuthMiddleware(), func(c *gin.Context) {
			var data struct {
				Page   string `json:"page"`
				Offset string `json:"offset"`
			}
			if err := c.BindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			page, _ := strconv.ParseInt(data.Page, 10, 8)
			offset := int(page) * int(1)

			result := []Boot.Users{}
			if page == 1 {
				result = model.GetAllUsersByPaginate(0, postperpage)

			} else {
				result = model.GetAllUsersByPaginate(offset, postperpage)

			}
			c.JSON(http.StatusOK, gin.H{"message": result})
		})

		v2.GET("/edituser", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			currentusrt := model.GetCurrentUser(c)
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"Username": session.Get("UserName"),
				"title":    "ویرایش کاربر",
				"user":     currentusrt,
				"action":   "edituser",
			})
		})
		v2.POST("/edituser", middleware.AuthMiddleware(), func(c *gin.Context) {
			currentusrt := model.GetCurrentUser(c)
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
				"InventoryNumber": Utility.GetCurrentInventory(c),
			})
		})
		v2.POST("/addproduct", middleware.AuthMiddleware(), func(c *gin.Context) {
			var product boot.Inventory
			product.Name = c.PostForm("Name")
			product.Number = c.PostForm("Number")
			product.RolePrice = Utility.StringToFloat(c.PostForm("RolePrice"))
			product.MeterPrice = Utility.StringToFloat(c.PostForm("MeterPrice"))
			product.Count = Utility.StringToInt(c.PostForm("Count"))
			product.InventoryNumber = Utility.StringToInt32(c.PostForm("InventoryNumber"))
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
				"products": model.GetAllProductsByInventory(Utility.GetCurrentInventory(c)),
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
		v2.GET("/deleteExport", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)

			if model.RemoveCurrentExport(c) {
				c.HTML(http.StatusOK, "export_list.html", gin.H{
					"Username": session.Get("UserName"),
					"title":    "فاکتورها",
					"message":  boot.Messages("Export removed success"),
					"success":  true,
					"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfExports()/1, "export-list")),
					"exports":  model.GetAllExportsByPaginate(0, postperpage),
				})
			} else {
				c.HTML(http.StatusOK, "export_list.html", gin.H{
					"Username": session.Get("UserName"),
					"title":    "فاکتورها",
					"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfExports()/1, "export-list")),
					"exports":  model.GetAllExportsByPaginate(0, postperpage),
				})
			}
		})
		v2.POST("/getproductbyinventory", func(c *gin.Context) {
			var data struct {
				Name string `json:"name"`
				Id   string `json:"id"`
			}
			if err := c.BindJSON(&data); err != nil {
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
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			val, _ := strconv.ParseInt(data.Id, 10, 8)
			product := model.GetProductById(int(val))
			c.JSON(http.StatusOK, gin.H{"result": product})
		})
		v2.POST("/export", func(c *gin.Context) {

			// make  product  struct  for bind and for data struct

			type Product struct {
				ID              string `gorm:"primaryKey"`
				ProductId       string `gorm:"size:255;"`
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
			// make  Data  struct  for bind
			var data struct {
				Name     string    `json:"Name"`
				Content  string    `json:"Content"`
				Products []Product `json:"Products"`
			}
			// bind data from ajax to Data

			if err := c.BindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			fmt.Println("serialize", data.Content, "serialize")

			// bind data struct to  ExportProducts for make row in db

			exportproducts := make([]boot.ExportProducts, len(data.Products))
			Ids := make(map[int64]int64)
			for a, _ := range data.Products {
				ids, _ := strconv.ParseInt(data.Products[a].ProductId, 10, 64)

				exportproducts[a].ExportID, _ = strconv.ParseUint(data.Products[a].ExportID, 10, 64)
				exportproducts[a].Name = data.Products[a].Name
				exportproducts[a].Number = data.Products[a].ExportID
				exportproducts[a].RolePrice, _ = strconv.ParseFloat(data.Products[a].RolePrice, 64)
				exportproducts[a].MeterPrice, _ = strconv.ParseFloat(data.Products[a].MeterPrice, 64)
				Count, _ := strconv.ParseInt(data.Products[a].Count, 10, 8)
				exportproducts[a].Count = int8(Count)
				InventoryNumber, _ := strconv.ParseInt(data.Products[a].InventoryNumber, 10, 32)

				exportproducts[a].InventoryNumber = int32(InventoryNumber)
				Ids[int64(ids)] = int64(Count)

			}

			// bind result " data.Content " from ajax to  Export for make row in db
			User := boot.Users{}

			Export := boot.Export{}
			result := Utility.Unserialize(data.Content)
			User.Name, Export.Name = result["Name"], result["Name"]
			Export.Number = result["Number"]
			User.Phonenumber, Export.Phonenumber = result["Phonenumber"], result["Phonenumber"]
			User.Address, Export.Address = result["Address"], result["Address"]
			Export.TotalPrice = Utility.StringToFloat(result["TotalPrice"])
			Export.Tax = Utility.StringToFloat(result["Tax"])
			Export.CreatedAt = string(Utility.CurrentTime())
			Export.InventoryNumber = Utility.StringToInt32(result["InventoryNumber"])
			// fmt.Println("inve", result["InventoryNumber"], Export.InventoryNumber, "/inve")
			Export.ExportProducts = exportproducts
			// output: Export,exportproducts
			User.Role = "guest"
			boot.DB().Create(&User)
			if boot.DB().Create(&exportproducts).RowsAffected > 0 && boot.DB().Create(&Export).RowsAffected > 0 {
				controller.InventoryCalculation(Ids)
				session := sessions.Default(c)

				c.HTML(http.StatusOK, "exportshow.html", gin.H{
					"Username": session.Get("UserName"),
					"title":    "فاکتورها",
					"exports":  model.GetAllExports(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "invalid request"})
			}

		})
		v2.GET("/export-list", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "export_list.html", gin.H{
				"Username": session.Get("UserName"),
				"title":    "فاکتورها",
				"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfExports()/1, "export-list")),
				"exports":  model.GetAllExportsByPaginate(0, postperpage),
			})
		})

		v2.POST("/export-list", middleware.AuthMiddleware(), func(c *gin.Context) {
			var data struct {
				Page   string `json:"page"`
				Offset string `json:"offset"`
			}
			if err := c.BindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			page, _ := strconv.ParseInt(data.Page, 10, 8)
			offset := int(page) * int(1)

			result := []Boot.EscapeExport{}
			if page == 1 {
				result = model.GetAllExportsByPaginate(0, postperpage)

			} else {
				result = model.GetAllExportsByPaginate(offset, postperpage)

			}
			c.JSON(http.StatusOK, gin.H{"message": result})
		})
		v2.POST("/export-find", middleware.AuthMiddleware(), func(c *gin.Context) {
			var data struct {
				Term string `json:"term"`
			}
			if err := c.BindJSON(&data); err != nil {

				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			result := model.GetAllExportsByPhoneAndName(data.Term)
			c.JSON(http.StatusOK, gin.H{"message": result})
		})
		v2.GET("/exportshow", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "exportshow.html", gin.H{
				"Username": session.Get("UserName"),
				"title":    "فاکتورها",
				"exports":  model.GetAllExports(),
			})
		})

	}
	// Dashboard Route

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
