package Model

import (
	"fmt"
	"inventory/App/Boot"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get All user by role , arg is [Admin , guest]

func GetAllUsersByRole(role string) []Boot.Users {
	var users []Boot.Users
	switch role {
	case "Admin":
		Boot.DB().Model(&Boot.Users{}).Select("*").Where("role = ? ", "Admin").Scan(&users)
	case "Author":
		Boot.DB().Model(&Boot.Users{}).Select("*").Where("role = ? ", "Author").Scan(&users)
	case "guest":
		Boot.DB().Model(&Boot.Users{}).Select("*").Where("role = ? ", "guest").Scan(&users)
	}

	return users
}

// get user By email [hosseinbidar7@gmail.com]

func GetUserByEmail(login Boot.Login) []Boot.Users {
	var users []Boot.Users
	Boot.DB().Model(Boot.Users{}).Select("*").Where("Email = ?", login.Email).Scan(&users)
	return users
}

// get user by id
func GetUserById(userid Boot.Users) []Boot.Users {
	var user []Boot.Users
	Boot.DB().Model(&Boot.Users{}).Select("*").Where("ID = ?", userid.ID).Scan(&user)
	return user
}
func GetUserRoleById(userid Boot.Users) string {
	var role string
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
	Boot.DB().Model(&[]Boot.Users{}).Find(&[]Boot.Users{}).Count(&count)
	return count
}
func GetAllUsersByPaginate(offset int, limit int, role string) []Boot.Users {
	var Users []Boot.Users
	Boot.DB().Model(&Boot.Users{}).Select("*").Where("role = ?", role).Offset(offset).Limit(limit).Scan(&Users)
	return Users
}

func RemoveCurrentUser(c *gin.Context) bool {
	Id := c.Request.URL.Query().Get("user-id")
	fmt.Println(Id)
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

	Boot.DB().Model(&Boot.Users{}).Where("name LIKE ? OR phonenumber LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%").Find(&Users)
	return Users
}
