package model

import (
	"inventory/boot"
)

// get All user by role , arg is [Admin , guest]

func GetAllUsersByRole(role string) []boot.Users {
	var users []boot.Users
	switch role {
	case "Admin":
		boot.DB().Model(&boot.Users{}).Select("*").Where("role = ? ", "Admin").Scan(&users)
	case "guest":
		boot.DB().Model(&boot.Users{}).Select("*").Where("role = ? ", "guest").Scan(&users)
	}

	return users
}

// get user By email [hosseinbidar7@gmail.com]

func GetUserByEmail(login boot.Login) []boot.Users {
	var users []boot.Users
	boot.DB().Model(boot.Users{}).Select("*").Where("Email = ?", login.Email).Scan(&users)
	return users
}

// get user by id
func GetUserById(userid boot.Users) []boot.Users {
	var user []boot.Users
	boot.DB().Model(&boot.Users{}).Select("*").Where("ID = ?", userid.ID).Scan(&user)
	return user
}
