package utility

import (
	"fmt"
	model "inventory/app/Model"
	"inventory/boot"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func TeataSay() {

	fmt.Println("asd")
}
func HomeUrl() string {
	return "http://127.0.0.1:8080"
}

// password hashing

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func GetCurrentUser(c *gin.Context) []boot.Users {
	Id := c.Request.URL.Query().Get("user-id")
	userIdUint, _ := strconv.ParseUint(Id, 10, 64)
	USERID := boot.Users{ID: userIdUint}
	currentUser := model.GetUserById(USERID)
	return currentUser
}
