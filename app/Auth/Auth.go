package auth

import (
	"fmt"
	model "inventory/app/Model"
	"inventory/app/utility"
	"inventory/boot"
)

func CheckAuth(login boot.Login) (bool, string) {
	pass := ""
	name := ""
	CurentUser := model.GetUserByEmail(login)
	for _, user := range CurentUser {
		pass = user.Password
		name = user.Name
	}
	fmt.Println("aaAAAAAAAAAAAAAAAA", name)
	dbpass := login.Password
	result := false
	if utility.CheckPasswordHash(dbpass, pass) {

		result = true
	}
	return result, name

}
