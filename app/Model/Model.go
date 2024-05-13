package model

import (
	"inventory/boot"
)

type Users struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;index:idx_name,unique"`
	Email       string `gorm:"size:255;"`
	Password    string `gorm:"type:varchar(255)"`
	Phonenumber string `gorm:"size:255;"`
	Role        string `gorm:"size:255;"`
}

// get All user by role , arg is [Admin , guest]

func GetAllUsersByRole(role string) []Users {
	var users []Users
	switch role {
	case "Admin":
		boot.DB().Model(&Users{}).Select("*").Where("role = ? ", "Admin").Scan(&users)
	case "guest":
		boot.DB().Model(&Users{}).Select("*").Where("role = ? ", "guest").Scan(&users)
	}

	return users
}

// get user By email [hosseinbidar7@gmail.com]

func GetUserByEmail(login boot.Login) []Users {
	var users []Users
	boot.DB().Model(&Users{}).Select("*").Where("Email = ?", login.Email).Scan(&users)
	return users
}
