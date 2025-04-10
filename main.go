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
	const postperpage int = 50
	boot.Init()
	go boot.PeroudBackup()

	r := gin.Default()
	r.LoadHTMLGlob("Views/templates/*")
	r.Static("assets", "./assets")
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// Defualt  Routes
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "ایزوگام شرق و دلیجان",
		})
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
			authorized, name, role := auth.CheckAuth(login)

			// fmt.Print(authorized, c.PostForm("email"), c.b, authorized)
			if authorized {
				session := sessions.Default(c)
				if session.Get("Auth") != "logedin" {
					session.Clear()
					session.Set("Auth", "logedin")
					session.Set("UserRole", role)
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
				"title":      "صفحه اصلی",
				"Username":   session.Get("UserName"),
				"UserRole":   session.Get("UserRole"),
				"message":    boot.Messages("login success"),
				"success":    true,
				"users":      model.GetAllUsersByRole("guest"),
				"exports":    model.GetAllExportsByPaginate(0, 5),
				"allexports": model.GetAllExports(),
			})
		})
		v2.GET("/api/allexports", middleware.AuthMiddleware(), func(c *gin.Context) {
			c.JSON(http.StatusOK,
				model.GetAllExports(),
			)
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
				"UserRole": session.Get("UserRole"),
				"title":    "کاربران",
				"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfUsers()/int64(postperpage), "user-list")),
				"users":    model.GetAllUsersByPaginate(0, postperpage, "guest"),
			})
		})
		v2.GET("/authors", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			// result := "a"
			c.HTML(http.StatusOK, "users.html", gin.H{

				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "کاربران",
				"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfUsers()/int64(postperpage), "user-list")),
				"users":    model.GetAllUsersByPaginate(0, postperpage, "author"),
			})
		})
		v2.GET("/admin_users", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "admins.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "کاربران ادمین",
				"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfUsers()/int64(postperpage), "user-list")),
				"users":    model.GetAllUsersByPaginate(0, postperpage, "Admin"),
			})
		})
		v2.GET("/add_user", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "کاربران",
				"action":   "add_user",
			})
		})
		v2.POST("/add_user", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
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
					"UserRole": session.Get("UserRole"),
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
		v2.POST("/user-list", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
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
				result = model.GetAllUsersByPaginate(0, postperpage, "Admin")

			} else {
				result = model.GetAllUsersByPaginate(offset, postperpage, "Admin")

			}
			c.JSON(http.StatusOK, gin.H{"message": result})
		})
		v2.GET("/deleteuser", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)

			if model.RemoveCurrentUser(c) {
				c.HTML(http.StatusOK, "users.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "کاربران",
					"message":  boot.Messages("user remove success"),
					"success":  true,
					"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfUsers()/1, "user-list")),
					"users":    model.GetAllUsersByPaginate(0, postperpage, "Admin"),
				})
			} else {
				c.HTML(http.StatusOK, "users.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "فاکتورها",
					"success":  false,
					"message":  boot.Messages("user remove faild"),
					"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfUsers()/1, "user-list")),
					"users":    model.GetAllExportsByPaginate(0, postperpage),
				})
			}
		})
		v2.GET("/edituser", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)
			currentusrt := model.GetCurrentUser(c)
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "ویرایش کاربر",
				"user":     currentusrt,
				"action":   "edituser",
			})
		})
		v2.POST("/edituser", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			currentusrt := model.GetCurrentUser(c)
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"title":   "ویرایش کاربر",
				"action":  "add_user",
				"message": "کاربر با موفقیت  اصلاح شد",
				"success": true,
				"user":    currentusrt,
			})
		})
		v2.POST("/users-find", middleware.AuthMiddleware(), func(c *gin.Context) {
			var data struct {
				Term string `json:"term"`
			}
			if err := c.BindJSON(&data); err != nil {

				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			result := model.GetAllUsersByPhoneAndName(data.Term)
			c.JSON(http.StatusOK, gin.H{"message": result})
		})
		// Product
		v2.GET("/addproduct", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "add_product.html", gin.H{
				"Username":        session.Get("UserName"),
				"UserRole":        session.Get("UserRole"),
				"title":           "محصول",
				"action":          "addproduct",
				"InventoryNumber": Utility.GetCurrentInventory(c),
			})
		})
		v2.POST("/addproduct", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			var product boot.Inventory
			product.Name = c.PostForm("Name")
			product.Number = c.PostForm("Number")
			product.RolePrice = Utility.StringToInt64(c.PostForm("RolePrice"))
			product.MeterPrice = Utility.StringToInt64(c.PostForm("MeterPrice"))
			product.Count = Utility.StringToInt64(c.PostForm("Count"))
			product.Meter = Utility.StringToInt64(c.PostForm("Meter"))
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
		v2.GET("/deleteproduct", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			// session := sessions.Default(c)
			inventory := c.Request.URL.Query().Get("inventory")
			// InventoryID, _ := strconv.ParseUint(inventory, 10, 64)

			if model.RemoveCurrentProduct(c) {
				c.Redirect(301, "inventory?inventory="+inventory)

			} else {
				c.Redirect(301, "inventory?inventory="+inventory)

			}
		})
		v2.GET("/editproduct", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)
			Id := c.Request.URL.Query().Get("product-id")
			ProductID, err := strconv.ParseInt(Id, 10, 8)
			if err != nil {
				// handle the error
				c.Redirect(301, "inventory?inventory=1")
				panic("product not found")
			}
			currentProduct := model.GetProductById(int(ProductID))

			c.HTML(http.StatusOK, "add_product.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "ویرایش کاربر",
				"products": currentProduct,
				"action":   "editproduct",
			})
		})
		v2.POST("/editproduct", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)

			var product boot.Inventory
			Id := c.PostForm("Id")
			ProductID, _ := strconv.ParseInt(Id, 10, 8)
			product.Name = c.PostForm("Name")
			product.RolePrice = Utility.StringToInt64(c.PostForm("RolePrice"))
			product.MeterPrice = Utility.StringToInt64(c.PostForm("MeterPrice"))
			product.Count = Utility.StringToInt64(c.PostForm("Count"))
			product.Meter = Utility.StringToInt64(c.PostForm("Meter"))
			product.InventoryNumber = Utility.StringToInt32(c.PostForm("InventoryNumber"))
			res := boot.DB().Model(&product).Where("id = ? ", ProductID).Updates(&product)
			currentProduct := model.GetProductById(int(ProductID))
			if res.RowsAffected > 0 {
				c.HTML(http.StatusOK, "add_product.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "ویرایش کاربر",
					"products": currentProduct,
					"action":   "editproduct",
					"message":  boot.Messages("product made success"),
					"success":  true,
				})
			} else {
				c.HTML(http.StatusOK, "add_product.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "ویرایش کاربر",
					// "products": currentProduct,
					"action":  "editproduct",
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
				"UserRole": session.Get("UserRole"),
				"title":    "انبار",
				"products": model.GetAllProductsByInventory(Utility.GetCurrentInventory(c)),
			})
		})

		// inventory  Production
		v2.GET("/production", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "production.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "تولید",
				"action":   "updateproduct",
				"products": model.GetAllProductsByInventory(int32(1)),
			})
		})
		v2.POST("/updateproduct", middleware.AuthMiddleware(), func(c *gin.Context) {
			var product boot.Inventory
			var oldproduct boot.Inventory
			product.ID, _ = strconv.ParseUint(c.PostForm("ProductName"), 10, 64)
			res := boot.DB().Model(&product).Where("id = ? ", product.ID).Scan(&oldproduct)
			if res.RowsAffected > 0 {
				product.Name = oldproduct.Name
				product.RolePrice = Utility.StringToInt64(c.PostForm("RolePrice"))
				product.MeterPrice = Utility.StringToInt64(c.PostForm("MeterPrice"))
				product.Count = oldproduct.Count + Utility.StringToInt64(c.PostForm("ProductsCount"))
				fmt.Println(oldproduct.Count, c.PostForm("ProductsCount"), product.Count)
				product.Meter = oldproduct.Meter + Utility.StringToInt64(c.PostForm("ProductMeter"))
				fmt.Println(oldproduct.Name, product.Name)
				product.InventoryNumber = 1

				res := boot.DB().Model(&product).Where("id = ? ", product.ID).Updates(&product)
				if res.RowsAffected > 0 {
					c.HTML(http.StatusOK, "production.html", gin.H{
						"title":   "محصول",
						"action":  "addproduct",
						"message": boot.Messages("product made success"),
						"success": true,
					})
				} else {
					c.HTML(http.StatusOK, "production.html", gin.H{
						"title":   "محصول",
						"action":  "addproduct",
						"message": boot.Messages("product made faild"),
						"success": false,
					})
				}
			}

		})
		// inventory  Production

		v2.GET("/export", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)

			uniqueString := Utility.MakeRandValue()
			if !model.CheckExportNumberFound(uniqueString) {
				uniqueString = Utility.MakeRandValue()
			} else {
				return
			}

			c.HTML(http.StatusOK, "export.html", gin.H{

				"Username":     session.Get("UserName"),
				"UserRole":     session.Get("UserRole"),
				"action":       "export",
				"title":        "فاکتور",
				"date":         Utility.CurrentTime(),
				"exportnumber": uniqueString,
				"products":     model.GetAllProductsByInventory(1),
			})

		})
		v2.GET("/deleteExport", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)

			if model.RemoveCurrentExport(c) {
				c.HTML(http.StatusOK, "export_list.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "فاکتورها",
					"message":  boot.Messages("Export removed success"),
					"success":  true,
					"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfExports()/1, "export-list")),
					"exports":  model.GetAllExportsByPaginate(0, postperpage),
				})
			} else {
				c.HTML(http.StatusOK, "export_list.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
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
			type PaymentRequest struct {
				Method     string `json:"Method"`
				Number     string `json:"Number"`
				Name       string `json:"Name"`
				TotalPrice string `json:"TotalPrice"`
				CreatedAt  string `json:"CreatedAt"`
				Status     string `json:"Status"`
			}
			// make  Data  struct  for bind
			var data struct {
				Name     string           `json:"Name"`
				Content  string           `json:"Content"`
				Products []Product        `json:"Products"`
				Payments []PaymentRequest `json:"Payments"`
			}
			// bind data from ajax to Data

			if err := c.BindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// fmt.Println("serialize", "serialize")
			// // bind data struct to  ExportProducts for make row in db
			exportproducts := make([]boot.ExportProducts, len(data.Products))
			Ids := make(map[int64]int64)
			for a, _ := range data.Products {
				ids, _ := strconv.ParseInt(data.Products[a].ProductId, 10, 64)

				exportproducts[a].ExportID, _ = strconv.ParseUint(data.Products[a].ExportID, 10, 64)
				fmt.Println(data.Products[a].ExportID)
				exportproducts[a].Name = data.Products[a].Name
				exportproducts[a].Number = data.Products[a].ExportID
				exportproducts[a].RolePrice, _ = strconv.ParseInt(data.Products[a].RolePrice, 10, 64)
				exportproducts[a].MeterPrice, _ = strconv.ParseInt(data.Products[a].MeterPrice, 10, 64)
				Count, _ := strconv.ParseInt(data.Products[a].Count, 10, 8)
				exportproducts[a].Count = int64(Count)
				InventoryNumber, _ := strconv.ParseInt(data.Products[a].InventoryNumber, 10, 32)
				exportproducts[a].TotalPrice, _ = strconv.ParseInt(data.Products[a].TotalPrice, 10, 64)
				fmt.Println(exportproducts[a].TotalPrice)
				exportproducts[a].InventoryNumber = int32(InventoryNumber)
				Ids[int64(ids)] = int64(Count)

			}

			// bind result " data.Content " from ajax to  Export for make row in db
			User := boot.Users{}
			fmt.Println("data", data)
			Export := boot.Export{}

			result := Utility.Unserialize(data.Content)
			User.Name, Export.Name = result["Name"], result["Name"]
			Export.Number = result["ExportID"]
			User.Phonenumber, Export.Phonenumber = result["Phonenumber"], result["Phonenumber"]
			User.Address, Export.Address = result["Address"], result["Address"]
			Tprice := Utility.StringToInt64(result["ExportTotalPrice"])
			Export.TotalPrice = Tprice
			Export.Tax = Utility.StringToInt64(result["Tax"])

			fmt.Println(result["Tax"], Utility.StringToFloat(result["Tax"]), Export.Tax)
			Export.CreatedAt = string(Utility.CurrentTime())
			Export.InventoryNumber = Utility.StringToInt32(result["InventoryNumber"])
			Export.Describe = result["describe"]

			Export.ExportProducts = exportproducts
			User.Role = "guest"
			if !model.CheckExportNumberFound(Export.Number) {
				Export.Number = Utility.MakeRandValue()
			} else {
				return
			}

			boot.DB().Create(&User)
			resexportproducts := boot.DB().Create(&exportproducts)
			resExport := boot.DB().Create(&Export)
			// fmt.Println("ddddddddddd", resExport, resexportproducts)
			if resexportproducts.RowsAffected > 0 && resExport.RowsAffected > 0 {
				controller.InventoryCalculation(Ids)
				resExport := boot.DB().Last(&Export)
				fmt.Println(resExport, Export.ID)
				for _, payment := range data.Payments {
					// Convert and create payment records
					totalPrice, _ := strconv.ParseInt(payment.TotalPrice, 10, 64)
					dbPayment := boot.Payments{
						Method:     payment.Method,
						Number:     payment.Number,
						Name:       payment.Name,
						TotalPrice: totalPrice,
						CreatedAt:  payment.CreatedAt,
						Status:     payment.Status,
						ExportID:   Export.ID,
					}
					boot.DB().Create(&dbPayment)

				}

				c.JSON(http.StatusOK, gin.H{"message": "sucess", "id": Export.ID})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "invalid request"})
			}

		})
		v2.GET("/export-list", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "export_list.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
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
			offset := int(page) * int(postperpage)
			fmt.Println(offset, page)
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
			exports, products := model.GetExportById(c)
			c.HTML(http.StatusOK, "exportshow.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "فاکتورها",
				"exports":  exports,
				"products": products,
			})
		})
		v2.GET("/payments", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "payments.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "پرداخت ها",
				"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfExports()/1, "export-list")),
				"Payments": model.GetAllPaymentsWithExportNumber(0, postperpage),
			})
		})
		v2.GET("/deletePayments", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)

			if model.RemoveCurrentPayments(c) {
				c.HTML(http.StatusOK, "payments.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "فاکتورها",
					"message":  boot.Messages("payments removed success"),
					"success":  true,
					"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfExports()/1, "export-list")),
					"Payments": model.GetAllPaymentsByPaginate(0, postperpage),
				})
			} else {
				c.HTML(http.StatusOK, "payments.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "فاکتورها",
					"Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfExports()/1, "export-list")),
					"Payments": model.GetAllPaymentsByPaginate(0, postperpage),
				})
			}
		})
		// v2.GET("/Download", middleware.AuthMiddleware(), func(c *gin.Context) {
		// 	session := sessions.Default(c)
		// 	// Id := c.Request.URL.Query().Get("ExportId")
		// 	exports, products := model.GetExportById(c)
		// 	Utility.GooglePDF("http://127.0.0.1:8080/Dashboard/exportshow?ExportId=3", "file1")
		// 	c.HTML(http.StatusOK, "exportshow.html", gin.H{
		// 		"Username": session.Get("UserName"),
		// 		"title":    "فاکتورها",
		// 		"exports":  exports,
		// 		"products": products,
		// 	})
		// })

	}
	// Dashboard Route

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
