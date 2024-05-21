package Model

import (
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
func GetAllUsersByPaginate(offset int, limit int) []Boot.Users {
	var Users []Boot.Users
	Boot.DB().Model(&Boot.Users{}).Select("*").Offset(offset).Limit(limit).Scan(&Users)
	return Users
}
