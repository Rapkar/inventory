package auth

import (
	"fmt"
	"inventory/App/Boot"
	"inventory/App/Model"
	"inventory/App/Utility"
)

func CheckAuth(login Boot.Login) (bool, string) {
	pass := ""
	name := ""
	CurentUser := Model.GetUserByEmail(login)
	for _, user := range CurentUser {
		pass = user.Password
		name = user.Name
	}
	fmt.Println("aaAAAAAAAAAAAAAAAA", name)
	dbpass := login.Password
	result := false
	if Utility.CheckPasswordHash(dbpass, pass) {

		result = true
	}
	return result, name

}
