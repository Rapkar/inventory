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
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func main() {

	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("❌ Failed to open log file: %v", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	log.Println("🚀 Logger initialized")

	log.Println("🔧 Booting application...")
	const postperpage int = 20
	boot.Init()

	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("🔥 Panic in ScheduleBackups:", r)
			}
		}()
		boot.ScheduleBackups()
		log.Println("📦 ScheduleBackups started")
	}()

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

			authorized, name, role, err := auth.CheckAuth(login)

			if err != nil || !authorized {
				// اگر خطایی رخ داد، کاربر را به صفحه لاگین با پیام خطا هدایت کنیم
				log.Println("❌ Error in CheckAuth:", err)
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"message": "ایمیل یا رمز عبور اشتباه است",
				})
				return
			}

			if authorized {
				// اگر لاگین موفقیت‌آمیز بود
				session := sessions.Default(c)
				if session.Get("Auth") != "logedin" {
					session.Clear()
					session.Set("Auth", "logedin")
					session.Set("UserRole", role)
					session.Set("UserName", name)
					session.Save()
				}

				// ریدایرکت به داشبورد پس از ورود موفق
				c.Redirect(http.StatusMovedPermanently, Utility.HomeUrl()+"/Dashboard")
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

			session.Delete("mysession")
			session.Delete("Auth")
			session.Delete("UserRole")
			session.Delete("UserName")
			session.Clear()

			c.SetCookie("mysession", "", -1, "/", Utility.HomeUrl(), false, true)

			c.Redirect(http.StatusMovedPermanently, Utility.HomeUrl()+"/auth")
		})

		v2.GET("/users", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			// result := "a"
			pageStr := c.DefaultQuery("page", "1")
			page, _ := strconv.Atoi(pageStr)
			if page < 1 {
				page = 1
			}
			offset := (page - 1) * postperpage

			totalItems, err := model.GetCountOfUsers()
			if err != nil {
				log.Println("❌ Error fetching GetCountOfUsers:", err)
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"message": "خطا در دریافت اطلاعات کاربر. لطفاً دوباره تلاش کنید.",
				})
				return
			}

			c.HTML(http.StatusOK, "users.html", gin.H{

				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "کاربران",
				// "Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfUsers()/int64(postperpage), "user-list")),
				"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "users")),
				"users":       model.GetAllUsersByPaginate(offset, postperpage, "guest"),
				"CurrentPage": page,
			})
		})
		v2.GET("/authors", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			// result := "a"

			pageStr := c.DefaultQuery("page", "1")
			page, _ := strconv.Atoi(pageStr)
			if page < 1 {
				page = 1
			}
			offset := (page - 1) * postperpage

			totalItems, err := model.GetCountOfUsers()
			if err != nil {
				log.Println("❌ Error fetching GetCountOfUsers:", err)
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"message": "خطا در دریافت اطلاعات کاربر. لطفاً دوباره تلاش کنید.",
				})
				return
			}
			c.HTML(http.StatusOK, "users.html", gin.H{

				"Username":    session.Get("UserName"),
				"UserRole":    session.Get("UserRole"),
				"title":       "کاربران",
				"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "Author")),
				"users":       model.GetAllUsersByPaginate(offset, postperpage, "Author"),
				"CurrentPage": page,
			})
		})
		v2.GET("/admin_users", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)
			pageStr := c.DefaultQuery("page", "1")
			page, _ := strconv.Atoi(pageStr)
			if page < 1 {
				page = 1
			}
			offset := (page - 1) * postperpage

			totalItems, err := model.GetCountOfUsers()
			if err != nil {
				log.Println("❌ Error fetching GetCountOfUsers:", err)
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"message": "خطا در دریافت اطلاعات کاربر. لطفاً دوباره تلاش کنید.",
				})
				return
			}

			c.HTML(http.StatusOK, "admins.html", gin.H{
				"Username":    session.Get("UserName"),
				"UserRole":    session.Get("UserRole"),
				"title":       "کاربران ادمین",
				"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "admin_users")),
				"users":       model.GetAllUsersByPaginate(offset, postperpage, "Admin"),
				"CurrentPage": page,
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
			pageStr := c.DefaultQuery("page", "1")
			category := c.DefaultQuery("user", "users")
			page, _ := strconv.Atoi(pageStr)
			if page < 1 {
				page = 1
			}
			offset := (page - 1) * postperpage

			totalItems, err := model.GetCountOfUsers()
			if err != nil {
				log.Println("❌ Error fetching GetCountOfUsers:", err)
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"message": "خطا در دریافت اطلاعات کاربر. لطفاً دوباره تلاش کنید.",
				})
				return
			}

			if model.RemoveCurrentUser(c) {
				c.HTML(http.StatusOK, "users.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "کاربران",
					"message":  boot.Messages("user remove success"),
					"success":  true,
					// "Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfUsers()/1, "user-list")),
					"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "users")),
					"users":       model.GetAllUsersByPaginate(offset, postperpage, category),
					"CurrentPage": page,
				})
			} else {
				c.HTML(http.StatusOK, "users.html", gin.H{
					"Username":    session.Get("UserName"),
					"UserRole":    session.Get("UserRole"),
					"title":       "فاکتورها",
					"success":     false,
					"message":     boot.Messages("user remove faild"),
					"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "users")),
					"users":       model.GetAllUsersByPaginate(offset, postperpage, category),
					"CurrentPage": page,
				})
			}
		})
		v2.GET("/edituser", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)
			user, err := model.GetCurrentUser(c)
			fmt.Println(user)
			if err != nil {
				log.Println("❌ Error fetching current user:", err)
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"message": "خطا در دریافت اطلاعات کاربر. لطفاً دوباره تلاش کنید.",
				})
				return
			}
			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "ویرایش کاربر",
				"user":     user,
				"action":   "edituser",
			})

		})
		v2.GET("/user/details", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)
			user, err := model.GetCurrentUser(c)
			if err != nil {
				log.Println("❌ Error fetching current user:", err)
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"message": "خطا در دریافت اطلاعات کاربر. لطفاً دوباره تلاش کنید.",
				})
				return
			}
			UserFullDetails, err := model.GetUserFullDetailsByID(user.ID)

			if err != nil {
				log.Println("❌ Error fetching current user detail faild:", err)
			}

			type UserCalculations struct {
				ExportTotalprice float64
				TotalPaid        float64
				ExportsCount     int64
				DebtAmount       float64 // بهتر است با حرف بزرگ شروع شود (قابلیت export)
				CreditAmount     float64 // بهتر است با حرف بزرگ شروع شود
			}
			Totalprice, _ := model.GetUserTotalPrice(user.ID)
			TotalPaid, _ := model.GetUserTotalPaid(user.ID)
			var CreditAmount float64
			var DebtAmount float64

			difference := Totalprice - TotalPaid

			if difference > 0 {
				// مشتری بدهکار است (باید بیشتر پرداخت کند)
				CreditAmount = difference
				DebtAmount = 0
			} else if difference < 0 {
				// مشتری بستانکار است (پرداخت بیش از حد انجام داده)
				DebtAmount = -difference // استفاده از مقدار مطلق
				CreditAmount = 0
			} else {
				// تسویه کامل شده
				CreditAmount = 0
				DebtAmount = 0
			}
			var userCalc UserCalculations
			userCalc.ExportTotalprice, _ = model.GetUserTotalPrice(user.ID)
			userCalc.TotalPaid, _ = model.GetUserTotalPaid(user.ID)
			userCalc.ExportsCount, _ = model.GetCountOfUserExports(user.ID)
			userCalc.DebtAmount = DebtAmount
			userCalc.CreditAmount = CreditAmount
			c.HTML(http.StatusOK, "details_user.html", gin.H{
				"Username":         session.Get("UserName"),
				"UserRole":         session.Get("UserRole"),
				"title":            "ویرایش کاربر",
				"details":          UserFullDetails,
				"UserCalculations": userCalc,
				"action":           "edituser",
			})

		})
		v2.POST("/edituser", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)

			// دریافت شناسه کاربر
			userId := c.PostForm("ID")
			if userId == "" {
				c.HTML(http.StatusBadRequest, "edit_user.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "ویرایش کاربر",
					"action":   "edit_user",
					"message":  "شناسه کاربر الزامی است",
					"success":  false,
				})
				return
			}

			// یافتن کاربر موجود
			var user boot.Users
			if err := boot.DB().First(&user, userId).Error; err != nil {
				c.HTML(http.StatusNotFound, "edit_user.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "ویرایش کاربر",
					"action":   "edit_user",
					"message":  "کاربر یافت نشد",
					"success":  false,
				})
				return
			}

			// اعتبارسنجی و به‌روزرسانی فیلدها
			if name := strings.TrimSpace(c.PostForm("Name")); name != "" {
				if len(name) < 3 {
					c.HTML(http.StatusBadRequest, "edit_user.html", gin.H{
						"Username": session.Get("UserName"),
						"UserRole": session.Get("UserRole"),
						"title":    "ویرایش کاربر",
						"action":   "edit_user",
						"message":  "نام باید حداقل ۳ کاراکتر باشد",
						"success":  false,
					})
					return
				}
				user.Name = name
			}

			if email := strings.TrimSpace(c.PostForm("Email")); email != "" {
				if !Utility.IsValidEmail(email) {
					c.HTML(http.StatusBadRequest, "edit_user.html", gin.H{
						"Username": session.Get("UserName"),
						"UserRole": session.Get("UserRole"),
						"title":    "ویرایش کاربر",
						"action":   "edit_user",
						"message":  "فرمت ایمیل نامعتبر است",
						"success":  false,
					})
					return
				}
				user.Email = email
			}

			if phone := strings.TrimSpace(c.PostForm("Phonenumber")); phone != "" {
				if !Utility.IsValidPhoneNumber(phone) {
					c.HTML(http.StatusBadRequest, "edit_user.html", gin.H{
						"Username": session.Get("UserName"),
						"UserRole": session.Get("UserRole"),
						"title":    "ویرایش کاربر",
						"action":   "edit_user",
						"message":  "فرمت شماره تلفن نامعتبر است",
						"success":  false,
					})
					return
				}
				user.Phonenumber = phone
			}

			if role := strings.TrimSpace(c.PostForm("Role")); role != "" {
				if role != "Admin" && role != "Author" && role != "guest" {
					c.HTML(http.StatusBadRequest, "edit_user.html", gin.H{
						"Username": session.Get("UserName"),
						"UserRole": session.Get("UserRole"),
						"title":    "ویرایش کاربر",
						"action":   "edit_user",
						"message":  "نقش کاربر نامعتبر است",
						"success":  false,
					})
					return
				}
				user.Role = role
			}

			if address := strings.TrimSpace(c.PostForm("Address")); address == "" {

				c.HTML(http.StatusBadRequest, "edit_user.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "ویرایش کاربر",
					"action":   "edit_user",
					"message":  "آدرس نامعتبر است",
					"success":  false,
				})
				return

			} else {
				user.Address = address
			}

			if password := c.PostForm("Password"); password != "" {
				if len(password) < 8 {
					c.HTML(http.StatusBadRequest, "edit_user.html", gin.H{
						"Username": session.Get("UserName"),
						"UserRole": session.Get("UserRole"),
						"title":    "ویرایش کاربر",
						"action":   "edit_user",
						"message":  "رمز عبور باید حداقل ۸ کاراکتر باشد",
						"success":  false,
					})
					return
				}
				hashedPassword, err := Utility.HashPassword(password)
				if err != nil {
					c.HTML(http.StatusInternalServerError, "edit_user.html", gin.H{
						"Username": session.Get("UserName"),
						"UserRole": session.Get("UserRole"),
						"title":    "ویرایش کاربر",
						"action":   "edit_user",
						"message":  "خطا در پردازش رمز عبور",
						"success":  false,
					})
					return
				}
				user.Password = hashedPassword
			}

			// ذخیره تغییرات
			if err := boot.DB().Save(&user).Error; err != nil {
				c.HTML(http.StatusInternalServerError, "edit_user.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "ویرایش کاربر",
					"action":   "edit_user",
					"message":  "خطا در به‌روزرسانی کاربر",
					"success":  false,
				})
				return
			}

			c.HTML(http.StatusOK, "edit_user.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "ویرایش کاربر",
				"action":   "edit_user",
				"message":  "تغییرات با موفقیت ذخیره شدند",
				"success":  true,
			})
		})

		v2.POST("/users-find", middleware.AuthMiddleware(), func(c *gin.Context) {
			var data struct {
				Name  string `json:"name"`
				Phone string `json:"phone"`
			}
			if err := c.BindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			result, message := model.GetUsersByNameAndPhone(data.Name, data.Phone)
			c.JSON(http.StatusOK, gin.H{
				"message": message,
				"users":   result,
			})
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
			var product boot.Product
			product.Name = c.PostForm("Name")
			product.RolePrice, _ = Utility.StringToFloat64(c.PostForm("RolePrice"))
			product.MeterPrice, _ = Utility.StringToFloat64(c.PostForm("MeterPrice"))
			product.Count, _ = Utility.StringToInt64(c.PostForm("Count"))
			product.Meter, _ = Utility.StringToFloat64(c.PostForm("Meter"))
			product.Weight, _ = Utility.StringToFloat64(c.PostForm("Weight"))
			product.WeightPrice, _ = Utility.StringToFloat64(c.PostForm("WeightPrice"))
			product.InventoryID, _ = Utility.StringToUnit64(c.PostForm("InventoryNumber"))

			res := boot.DB().Create(&product)

			if res.RowsAffected > 0 {
				c.HTML(http.StatusOK, "add_product.html", gin.H{
					"title":   "محصول",
					"action":  "addproduct",
					"message": boot.Messages("product_add_success"),
					"success": true,
				})
			} else {
				c.HTML(http.StatusOK, "add_product.html", gin.H{
					"title":   "محصول",
					"action":  "addproduct",
					"message": boot.Messages("product_add_failed"),
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

			// دریافت مقادیر از فرم
			Id := c.PostForm("Id")
			ProductID, err := strconv.ParseInt(Id, 10, 8)
			if err != nil {
				c.HTML(http.StatusBadRequest, "add_product.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "ویرایش محصول",
					"error":    "شناسه محصول نامعتبر است",
					"formData": c.Request.PostForm,
				})
				return
			}

			// دریافت و تبدیل مقادیر
			rolePrice, _ := Utility.StringToFloat64(c.PostForm("RolePrice"))
			meterPrice, _ := Utility.StringToFloat64(c.PostForm("MeterPrice"))
			count, _ := Utility.StringToInt64(c.PostForm("Count"))
			meter, _ := Utility.StringToFloat64(c.PostForm("Meter"))
			Weight, _ := Utility.StringToFloat64(c.PostForm("Weight"))
			WeightPrice, _ := Utility.StringToFloat64(c.PostForm("WeightPrice"))
			inventoryID, _ := Utility.StringToUnit64(c.PostForm("InventoryNumber"))

			// اعتبارسنجی: حداقل دو فیلد باید مقدار مثبت داشته باشند
			validFields := 0
			if rolePrice > 0 {
				validFields++
			}
			if meterPrice > 0 {
				validFields++
			}
			if count > 0 {
				validFields++
			}
			if meter > 0 {
				validFields++
			}
			if Weight > 0 {
				validFields++
			}
			if WeightPrice > 0 {
				validFields++
			}

			if validFields < 2 {
				c.HTML(http.StatusBadRequest, "add_product.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "ویرایش محصول",
					"error":    "حداقل دو مورد از مقادیر (قیمت رول، قیمت متر، تعداد، متراژ) باید پر شوند",
					"formData": c.Request.PostForm,
					"products": model.GetProductById(int(ProductID)),
				})
				return
			}

			// آماده‌سازی محصول برای آپدیت
			product := boot.Product{
				Name:        c.PostForm("Name"),
				RolePrice:   rolePrice,
				MeterPrice:  meterPrice,
				Count:       count,
				Meter:       meter,
				Weight:      Weight,
				WeightPrice: WeightPrice,
				InventoryID: inventoryID,
			}

			// انجام عملیات آپدیت
			res := boot.DB().Model(&boot.Product{}).Where("id = ?", ProductID).Updates(&product)
			currentProduct := model.GetProductById(int(ProductID))

			if res.RowsAffected > 0 {
				c.HTML(http.StatusOK, "add_product.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "ویرایش محصول",
					"products": currentProduct,
					"action":   "editproduct",
					"message":  "محصول با موفقیت ویرایش شد",
					"success":  true,
				})
			} else {
				c.HTML(http.StatusOK, "add_product.html", gin.H{
					"Username": session.Get("UserName"),
					"UserRole": session.Get("UserRole"),
					"title":    "ویرایش محصول",
					"products": currentProduct,
					"action":   "editproduct",
					"error":    "خطا در ویرایش محصول",
					"formData": c.Request.PostForm,
				})
			}
		})
		// inventory
		v2.GET("/inventory", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			pageStr := c.DefaultQuery("page", "1")
			inventoryStr := c.DefaultQuery("inventory", "1")

			page, _ := strconv.Atoi(pageStr)
			if page < 1 {
				page = 1
			}

			inventoryID, _ := strconv.Atoi(inventoryStr)
			offset := (page - 1) * postperpage

			totalItems := model.GetCountOfProduct(1)

			c.HTML(http.StatusOK, "inventory.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "انبار",
				"Paginate": template.HTML(Utility.MakeinventoryPaginate(
					int64(totalItems),
					int64(postperpage),
					int64(page),
					"inventory",
					int32(inventoryID),
				)),
				"products":    model.GetAllProductsByInventoryAndPaginate(offset, postperpage, int32(inventoryID)),
				"CurrentPage": page,
				"InventoryID": inventoryID, // Pass to template if needed
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
			var product boot.Product
			var oldproduct boot.Product
			product.ID, _ = strconv.ParseUint(c.PostForm("ProductName"), 10, 64)

			res := boot.DB().Model(&product).Where("id = ? ", product.ID).Scan(&oldproduct)
			if res.RowsAffected > 0 {
				product.Name = oldproduct.Name
				product.RolePrice, _ = Utility.StringToFloat64(c.PostForm("RolePrice"))
				product.MeterPrice, _ = Utility.StringToFloat64(c.PostForm("MeterPrice"))
				pCount, _ := Utility.StringToInt64(c.PostForm("ProductsCount"))
				product.Count = oldproduct.Count + pCount
				pMeter, _ := Utility.StringToFloat64(c.PostForm("Weight"))
				product.Weight, _ = Utility.StringToFloat64(c.PostForm("WeightPrice"))
				product.WeightPrice, _ = Utility.StringToFloat64(c.PostForm("ProductMeter"))
				product.Meter = oldproduct.Meter + pMeter
				product.InventoryID = 1

				res := boot.DB().Model(&product).Where("id = ? ", product.ID).Updates(&product)
				if res.RowsAffected > 0 {
					c.HTML(http.StatusOK, "production.html", gin.H{
						"title": "محصول",
						// "action":   "addproduct",
						"message":  boot.Messages("product made success"),
						"success":  true,
						"products": model.GetAllProductsByInventory(int32(1)),
					})
				} else {
					c.HTML(http.StatusOK, "production.html", gin.H{
						"title": "محصول",
						// "action":   "addproduct",
						"message":  boot.Messages("product made success"),
						"success":  false,
						"products": model.GetAllProductsByInventory(int32(1)),
					})
				}
			}

		})

		v2.POST("/updatepayment", middleware.AuthMiddleware(), func(c *gin.Context) {
			var Payments boot.Payments
			var oldPayments boot.Payments

			Payments.ID, _ = strconv.ParseUint(c.PostForm("PaymentID"), 10, 64)
			res := boot.DB().Model(&Payments).Where("id = ? ", Payments.ID).Scan(&oldPayments)
			if res.RowsAffected > 0 {
				Payments.Method = oldPayments.Method
				Payments.Number = c.PostForm("PaymentNumber")
				Payments.Name = c.PostForm("PaymentName")
				Payments.TotalPrice, _ = Utility.StringToFloat64(c.PostForm("PaymentTotalPrice"))
				Payments.Describe = c.PostForm("PaymentDescribe")
				if c.PostForm("CreatedAt") != "" {
					Payments.CreatedAt = c.PostForm("CreatedAt")
				} else {
					Payments.CreatedAt = Utility.CurrentTime()

				}
				Payments.Status = c.PostForm("PaymentStatus")

				res := boot.DB().Model(&Payments).Where("id = ? ", Payments.ID).Updates(&Payments)
				session := sessions.Default(c)
				status := c.Query("status")

				if res.RowsAffected > 0 {
					res, _ := model.GetAllPaymentsWithExportNumberAndUser(0, postperpage, status)
					c.HTML(http.StatusOK, "payments.html", gin.H{
						"Username": session.Get("UserName"),
						"UserRole": session.Get("UserRole"),
						"title":    "پرداخت ها",
						"success":  true,

						// "Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfExports()/1, "export-list")),
						"Payments": res,
					})
				} else {

					c.HTML(http.StatusOK, "payments.html", gin.H{
						"Username": session.Get("UserName"),
						"UserRole": session.Get("UserRole"),
						"title":    "پرداخت ها",
						"success":  false,

						// "Paginate": template.HTML(Utility.MakePaginate(model.GetCountOfExports()/1, "export-list")),
						"Payments": res,
					})
				}
			}

		})
		// inventory  Production

		v2.GET("/export", middleware.AuthMiddleware(), func(c *gin.Context) {

			uniqueString, _ := model.GenerateUniqueExportID()
			// if !model.CheckExportNumberFound(uniqueString) {
			// 	uniqueString = Utility.MakeRandValue()
			// } else {
			// 	return
			// }
			session := sessions.Default(c)

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
			pageStr := c.DefaultQuery("page", "1")
			page, _ := strconv.Atoi(pageStr)
			if page < 1 {
				page = 1
			}
			offset := (page - 1) * postperpage

			totalItems := model.GetCountOfExports()

			if model.RemoveCurrentExport(c) {
				c.HTML(http.StatusOK, "export_list.html", gin.H{
					"Username":    session.Get("UserName"),
					"UserRole":    session.Get("UserRole"),
					"title":       "فاکتورها",
					"message":     boot.Messages("Export removed success"),
					"success":     true,
					"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "export-list")),
					"exports":     model.GetAllExportsByPaginate(offset, postperpage),
					"CurrentPage": page,
				})
			} else {
				c.HTML(http.StatusOK, "export_list.html", gin.H{
					"Username":    session.Get("UserName"),
					"UserRole":    session.Get("UserRole"),
					"title":       "فاکتورها",
					"message":     boot.Messages("Export removed success"),
					"success":     false,
					"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "export-list")),
					"exports":     model.GetAllExportsByPaginate(offset, postperpage),
					"CurrentPage": page,
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
			// Product struct for binding and data handling
			// type Product struct {
			// 	ID          string `gorm:"primaryKey"`
			// 	ProductId   string `gorm:"size:255;"`
			// 	ExportID    string `gorm:"size:255;"`
			// 	Name        string `gorm:"type:varchar(100)" json:"name"`
			// 	Number      string `gorm:"size:255;"`
			// 	RolePrice   string `gorm:"type:float"`
			// 	MeterPrice  string `gorm:"type:float"`
			// 	Count       string `gorm:"size:255;"`
			// 	Meter       string `gorm:"size:255;"`
			// 	TotalPrice  string `gorm:"index"`
			// 	InventoryID uint64 `gorm:"size:255;"`
			// }
			type ExportProducts struct {
				ID          uint64         `gorm:"primaryKey"`
				ExportID    uint64         `gorm:"index"`
				Name        string         `gorm:"type:varchar(100)"`
				Number      string         `gorm:"size:255;"`
				RolePrice   float64        `gorm:"type:float" json:",string"`
				MeterPrice  float64        `gorm:"type:float" json:",string"`
				Count       int64          `gorm:"size:255;" json:",string"`
				Meter       float64        `gorm:"type:float"`
				TotalPrice  float64        `gorm:"type:float" json:",string"`
				InventoryID uint64         `gorm:"index" json:",string"`
				Export      Boot.Export    `gorm:"foreignKey:ExportID;references:ID"`
				Inventory   Boot.Inventory `gorm:"foreignKey:InventoryID;references:ID"`
			}
			// type Product struct {
			// 	ID          uint64    `gorm:"primaryKey"`
			// 	Name        string    `gorm:"type:varchar(100)"`
			// 	Number      string    `gorm:"size:255;"`
			// 	RolePrice   float64   `gorm:"type:float"`
			// 	MeterPrice  float64   `gorm:"type:float"`
			// 	Count       int64     `gorm:"size:255;"`
			// 	Meter       float64   `gorm:"type:float"`
			// 	Weight      float64   `gorm:"type:float"`
			// 	InventoryID uint64    `gorm:"index"`
			// 	Inventory   Inventory `gorm:"foreignKey:InventoryID;references:ID"`
			// }
			// type PaymentRequest struct {
			// 	Method     string  `gorm:"size:255;"`
			// 	Number     string  `gorm:"size:255;"`
			// 	Name       string  `json:"Name"`
			// 	TotalPrice float64 `gorm:"type:float" json:",string"`
			// 	CreatedAt  string  `gorm:"size:255;"`
			// 	Status     string  `gorm:"size:255;"`
			// }
			type Payments struct {
				ID          uint64         `gorm:"primaryKey"`
				Method      string         `gorm:"type:varchar(100)"`
				Number      string         `gorm:"varchar(255),unique"`
				Name        string         `gorm:"type:varchar(100)"`
				TotalPrice  float64        `gorm:"type:float" json:",string"`
				Describe    string         `gorm:"size:255;"`
				CreatedAt   string         `json:"CreatedAt"` // assign the format to a string
				ExportID    uint64         `gorm:"index"`
				UserID      uint64         `gorm:"index"`
				InventoryID uint64         `gorm:"index"`
				Export      Boot.Export    `gorm:"foreignKey:ExportID"`
				Status      string         `gorm:"type:varchar(100)"`
				User        Boot.Users     `gorm:"foreignKey:UserID"`
				Inventory   Boot.Inventory `gorm:"foreignKey:InventoryID;references:ID"`
			}
			// Data struct for binding
			var data struct {
				Name       string           `json:"Name"`
				TotalPrice float64          `json:"TotalPrice"`
				Content    string           `json:"Content"`
				Products   []ExportProducts `json:"Products"`
				Payments   []Payments       `json:"Payments"`
			}
			// Bind data from request
			if err := c.BindJSON(&data); err != nil {
				fmt.Println(err)

				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Convert products to ExportProducts
			exportproducts := make([]boot.ExportProducts, len(data.Products))
			Ids := make(map[int64]int64)

			for a := range data.Products {
				// ids, _ := strconv.ParseInt(data.Products[a].ProductId, 10, 64)
				exportproducts[a].ExportID = data.Products[a].ExportID
				exportproducts[a].Name = data.Products[a].Name
				exportproducts[a].Number = data.Products[a].Number
				exportproducts[a].RolePrice = data.Products[a].RolePrice
				exportproducts[a].MeterPrice = data.Products[a].MeterPrice
				exportproducts[a].Count = data.Products[a].Count
				exportproducts[a].Meter = data.Products[a].Meter
				exportproducts[a].TotalPrice = data.TotalPrice
				exportproducts[a].InventoryID = data.Products[a].InventoryID
				exportproducts[a].TotalPrice = data.Products[a].TotalPrice
			}
			PaymentRequest := make([]boot.Payments, len(data.Payments))
			result := Utility.Unserialize(data.Content)
			// Phonenumber := result["Phonenumber"]

			invNum, err := Utility.StringToUnit64(result["InventoryNumber"])
			for a, payment := range data.Payments {
				totalPrice := payment.TotalPrice

				createdAt := payment.CreatedAt
				if createdAt == "" {
					createdAt = Utility.CurrentTime()
				}

				PaymentRequest[a].Method = payment.Method
				PaymentRequest[a].Number = payment.Number
				PaymentRequest[a].Name = payment.Name
				PaymentRequest[a].TotalPrice = totalPrice
				PaymentRequest[a].CreatedAt = createdAt
				PaymentRequest[a].Status = payment.Status
				PaymentRequest[a].InventoryID = invNum

			}
			// Process user and export data

			// if err != nil {
			// 	c.JSON(http.StatusBadRequest, gin.H{"error": "invalid total price"})
			// 	return
			// }
			tprice, _ := Utility.StringToFloat64(result["ExportTotalPrice"])
			tax, _ := Utility.StringToInt64(result["Tax"])

			Export := boot.Export{}
			// Database transaction
			err = boot.DB().Transaction(func(tx *gorm.DB) error {

				User := boot.Users{
					Name:        result["Name"],
					Phonenumber: result["Phonenumber"],
					Address:     result["Address"],
					Role:        "guest"}
				if err := tx.Where("phonenumber = ?", User.Phonenumber).First(&User).Error; err != nil {
					if err := tx.Create(&User).Error; err != nil {
						return err
					}
				}

				Export.UserID = User.ID
				Export.Name = result["Name"]
				Export.Number = result["ExportID"]
				Export.Phonenumber = result["Phonenumber"]
				Export.Address = result["Address"]
				Export.TotalPrice = tprice
				Export.Tax = tax
				Export.CreatedAt = string(Utility.CurrentTime())
				Export.ProductID = invNum
				Export.Describe = result["describe"]
				Export.ExportProducts = exportproducts

				if !model.CheckExportNumberFound(Export.Number) {
					newID, err := model.GenerateUniqueExportID()
					if err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate export ID"})
						return err
					}
					Export.Number = newID
				}

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "invalid inventory number"})
					return err
				}

				// Create user

				// Create export
				if err := tx.Where("id = ?", Export.ID).First(&Export).Error; err != nil {
					if err := tx.Create(&Export).Error; err != nil {
						return err
					}
				}
				// Update export products with the correct ExportID
				for i := range exportproducts {
					exportproducts[i].ExportID = Export.ID
				}
				if err := tx.Where("id = ?", Export.ID).First(&exportproducts).Error; err != nil {
					// Create export products
					if err := tx.Create(&exportproducts).Error; err != nil {
						return err
					}
				}

				// Process inventory
				controller.InventoryCalculation(Ids)

				// Create payments
				for i := range PaymentRequest {
					PaymentRequest[i].ExportID = Export.ID
					PaymentRequest[i].UserID = User.ID
				}
				if err := tx.Create(&PaymentRequest).Error; err != nil {
					return err
				}

				return nil
			})

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": "transaction failed",
					"error":   err.Error(),
				})
				return
			}
			log.Print("new Export submited")
			c.JSON(http.StatusOK, gin.H{"message": "sucess", "id": Export.ID})
		})
		v2.GET("/export-list", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			pageStr := c.DefaultQuery("page", "1")
			page, _ := strconv.Atoi(pageStr)
			if page < 1 {
				page = 1
			}
			offset := (page - 1) * postperpage

			totalItems := model.GetCountOfExports()

			c.HTML(http.StatusOK, "export_list.html", gin.H{
				"Username":    session.Get("UserName"),
				"UserRole":    session.Get("UserRole"),
				"title":       "فاکتورها",
				"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "export-list")),
				"exports":     model.GetAllExportsByPaginate(offset, postperpage),
				"CurrentPage": page,
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
			// fmt.Println(offset, page)
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
		v2.POST("/payment-find", middleware.AuthMiddleware(), func(c *gin.Context) {
			var data struct {
				Term string `json:"term"`
			}
			if err := c.BindJSON(&data); err != nil {

				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			result := model.GetAllPaymentsByAttribiute(data.Term)
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
			status := c.Query("status")
			pageStr := c.DefaultQuery("page", "1")
			page, _ := strconv.Atoi(pageStr)
			if page < 1 {
				page = 1
			}
			offset := (page - 1) * postperpage
			var res []boot.PaymentWithExportAndUser
			userid, err := Utility.StringToUnit64(c.Query("user_id"))
			if err == nil {
				res, err = model.GetAllPaymentsWithExportNumberByUserId(offset, postperpage, status, userid)
			} else {
				res, err = model.GetAllPaymentsWithExportNumberAndUser(offset, postperpage, status)
			}

			if err != nil {
				c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "خطا در دریافت پرداخت‌ها"})
				return
			}
			totalItems := model.GetCountOfPayments()
			c.HTML(http.StatusOK, "payments.html", gin.H{
				"Username":    session.Get("UserName"),
				"UserRole":    session.Get("UserRole"),
				"title":       "پرداخت ها",
				"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "payments")),
				"Payments":    res,
				"CurrentPage": page,
			})

		})
		v2.GET("/backup", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			// fmt.Println(resExport, Export.ID)
			fs := afero.NewOsFs()

			baseURL := viper.GetString("LOCAL_URL")
			backups, _ := boot.GetBackupList(fs, baseURL)
			c.HTML(http.StatusOK, "backups.html", gin.H{
				"Username": session.Get("UserName"),
				"UserRole": session.Get("UserRole"),
				"title":    "بک آپ ها",
				"backups":  backups,
			})
		})
		r.GET("/backups/:filename", func(c *gin.Context) {
			filename := c.Param("filename")
			fs := afero.NewOsFs()

			// بررسی وجود فایل با استفاده از afero.Fs
			fileInfo, err := fs.Stat(filename)
			if err != nil {
				if os.IsNotExist(err) {
					c.JSON(http.StatusNotFound, gin.H{"error": "فایل یافت نشد"})
					return
				}
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// باز کردن فایل با استفاده از afero
			file, err := fs.Open(filename)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer file.Close()

			// تنظیم هدرهای مناسب برای دانلود
			c.Header("Content-Description", "File Transfer")
			c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
			c.Header("Content-Type", "application/octet-stream")
			c.Header("Content-Length", fmt.Sprint(fileInfo.Size()))

			// ارسال فایل
			_, err = io.Copy(c.Writer, file)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
		})
		v2.GET("/charts", middleware.AuthMiddleware(), func(c *gin.Context) {
			session := sessions.Default(c)
			c.HTML(http.StatusOK, "charts.html", gin.H{
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
		v2.GET("/deletePayments", middleware.AuthMiddleware("Admin"), func(c *gin.Context) {
			session := sessions.Default(c)
			status := c.Query("status")

			pageStr := c.DefaultQuery("page", "1")
			page, _ := strconv.Atoi(pageStr)
			if page < 1 {
				page = 1
			}
			offset := (page - 1) * postperpage

			totalItems := model.GetCountOfPayments()
			res, _ := model.GetAllPaymentsWithExportNumberAndUser(offset, postperpage, status)

			if model.RemoveCurrentPayments(c) {
				c.HTML(http.StatusOK, "payments.html", gin.H{
					"Username":    session.Get("UserName"),
					"UserRole":    session.Get("UserRole"),
					"title":       "پرداختی ها",
					"message":     boot.Messages("payments removed success"),
					"success":     true,
					"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "payments")),
					"Payments":    res,
					"CurrentPage": page,
				})
			} else {
				c.HTML(http.StatusOK, "payments.html", gin.H{
					"Username":    session.Get("UserName"),
					"UserRole":    session.Get("UserRole"),
					"title":       "پرداختی ها",
					"message":     boot.Messages("payments removed faild"),
					"success":     false,
					"Paginate":    template.HTML(Utility.MakePaginate(int64(totalItems), int64(postperpage), int64(page), "payments")),
					"Payments":    res,
					"CurrentPage": page,
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
