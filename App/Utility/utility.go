package Utility

import (

	// Model "inventory/app/Model"

	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	ptime "github.com/yaa110/go-persian-calendar"
	"golang.org/x/crypto/bcrypt"
)

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

func GetCurrentInventory(c *gin.Context) int32 {
	Id := c.Request.URL.Query().Get("inventory")
	Inventoryid, _ := strconv.ParseInt(Id, 10, 32)

	return int32(Inventoryid)
}
func StringToFloat(value string) float64 {
	NewValue, _ := strconv.ParseFloat(value, 64)
	return float64(NewValue)
}
func StringToInt(value string) int8 {
	newValue, _ := strconv.ParseInt(value, 10, 8)
	return int8(newValue)
}
func StringToInt32(value string) int32 {
	newValue, _ := strconv.ParseInt(value, 10, 32)
	return int32(newValue)
}

// get the Persian Current Time
func CurrentTime() string {

	var pt = ptime.Now()
	exportFormat := pt.Format("yyyy/MM/dd E hh:mm:ss a")

	return exportFormat
}

// Unserialize string value to array with "&"" char
func Unserialize(value string) map[string]string {

	datas := make(map[string]string)
	pairs := strings.Split(value, "&")
	for _, val := range pairs {
		kv := strings.SplitN(val, "=", 2)
		if len(kv) == 2 {
			datas[kv[0]] = kv[1]
		}

	}
	return datas
}

//	func TableFound(table string) bool {
//		if db.Query("select * from "+table+";") == nil {
//			return false
//		} else {
//			return true
//		}
//	}
func FloatToString(value float64) string {
	val, err := json.Marshal(value)
	vals := ""
	if err == nil {
		vals = string(val)
	}
	return vals
}

func MakePaginate(value int64, url string) string {
	var paginate string
	for i := 1; i < int(value); i++ {
		if i == 1 {
			paginate += "<li class='page-item active'><a class='page-link' attr-page='" + fmt.Sprintf("%d", i) + "' href='./" + url + "/?page=" + fmt.Sprintf("%d", i) + "'> " + fmt.Sprintf("%d", i) + "</a></li> "
		} else {
			paginate += "<li class='page-item'><a class='page-link' attr-page='" + fmt.Sprintf("%d", i) + "' href='./" + url + "/?page=" + fmt.Sprintf("%d", i) + "'> " + fmt.Sprintf("%d", i) + "</a></li> "

		}
	}
	return paginate
}
