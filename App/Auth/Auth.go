package auth

import (
	"inventory/App/Boot"
	"inventory/App/Model"
	"inventory/App/Utility"
)

func CheckAuth(login Boot.Login) (bool, string, string) {
	pass := ""
	name := ""
	role := ""
	CurentUser := Model.GetUserByEmail(login)
	for _, user := range CurentUser {
		pass = user.Password
		name = user.Name
		role = user.Role
	}
	dbpass := login.Password
	result := false
	if Utility.CheckPasswordHash(dbpass, pass) {

		result = true
	}
	return result, name, role

}
