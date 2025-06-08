package auth

import (
	"fmt"
	"inventory/App/Boot"
	"inventory/App/Model"
	"inventory/App/Utility"
	"log"
)

func CheckAuth(login Boot.Login) (bool, string, string, uint64, error) {
	users, err := Model.GetUserByEmail(login.Email)
	if err != nil {
		log.Printf("❌ خطا در دریافت کاربر با ایمیل %s: %v", login.Email, err)
		return false, "", "", 0, err
	}

	if len(users) == 0 {
		return false, "", "", 0, fmt.Errorf("کاربری با این ایمیل یافت نشد")
	}

	user := users[0] // فرض بر اینکه ایمیل یونیک باشه

	// بررسی صحت رمز عبور
	if !Utility.CheckPasswordHash(login.Password, user.Password) {
		return false, "", "", 0, fmt.Errorf("رمز عبور اشتباه است")
	}

	return true, user.Name, user.Role, user.ID, nil
}
