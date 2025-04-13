package Model

import (
	"inventory/App/Boot"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get All user by role , arg is [Admin , guest]

func GetAllUsersByRole(role string) []Boot.Users {
	var users []Boot.Users
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	Boot.DB().Model(&Boot.Users{}).Select("*").Where("role = ? ", role).Scan(&users)

	return users
}

// get user By email [hosseinbidar7@gmail.com]

func GetUserByEmail(login Boot.Login) []Boot.Users {
	var users []Boot.Users
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	Boot.DB().Model(Boot.Users{}).Select("*").Where("Email = ?", login.Email).Scan(&users)

	return users
}

// get user by id
func GetUserById(userid Boot.Users) []Boot.Users {
	var user []Boot.Users
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	Boot.DB().Model(&Boot.Users{}).Select("*").Where("ID = ?", userid.ID).Scan(&user)
	return user
}
func GetUserRoleById(userid Boot.Users) string {
	var role string
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	Boot.DB().Model(&Boot.Users{}).Select("role").Where("ID = ?", userid.ID).Scan(&role)
	return role
}
func GetCurrentUser(c *gin.Context) []Boot.Users {
	Id := c.Request.URL.Query().Get("user-id")
	userIdUint, _ := strconv.ParseUint(Id, 10, 64)
	USERID := Boot.Users{ID: userIdUint}
	currentUser := GetUserById(USERID)
	return currentUser
}
func GetCountOfUsers() int64 {
	var count int64
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
		return 0
	}
	defer sqlDB.Close()
	Boot.DB().Model(&[]Boot.Users{}).Find(&[]Boot.Users{}).Count(&count)
	return count
}
func GetAllUsersByPaginate(offset int, limit int, role string) []Boot.Users {
	var Users []Boot.Users
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	Boot.DB().Model(&Boot.Users{}).Select("*").Where("role = ?", role).Offset(offset).Limit(limit).Scan(&Users)
	return Users
}

func RemoveCurrentUser(c *gin.Context) bool {
	Id := c.Request.URL.Query().Get("user-id")
	ExportID, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		// handle the error
		return false
	}
	result := Boot.DB().Delete(&Boot.Users{}, ExportID)
	if result.RowsAffected == 0 {
		// if no rows were affected, the deletion failed
		return false
	}
	return true
}
func GetAllUsersByPhoneAndName(searchTerm string) []Boot.Users {
	var Users []Boot.Users
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	Boot.DB().Model(&Boot.Users{}).Where("name LIKE ? OR phonenumber LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%").Find(&Users)
	return Users
}
