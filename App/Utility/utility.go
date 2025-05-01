package Utility

import (

	// Model "inventory/app/Model"
	"encoding/json"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ptime "github.com/yaa110/go-persian-calendar"
	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func IsValidPhoneNumber(phone string) bool {
	// الگوی ساده برای شماره تلفن (می‌توانید بر اساس نیاز تغییر دهید)
	phoneRegex := regexp.MustCompile(`^\+?[0-9]{10,15}$`)
	return phoneRegex.MatchString(phone)
}

func HomeUrl() string {
	MODE := viper.Get("MODE")
	var url string
	if MODE == "DEVELOP" {
		url = viper.GetString("LOCAL_URL")
	} else {
		url = viper.GetString("URL")

	}
	return url
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
func StringToFloat64(value string) (float64, error) {
	NewValue, err := strconv.ParseFloat(value, 64)
	return float64(NewValue), err
}
func StringToInt64(value string) (int64, error) {
	NewValue, err := strconv.ParseFloat(value, 64)
	return int64(NewValue), err
}
func StringToBool(value string) (bool, error) {
	return strconv.ParseBool(value)
}
func StringToUnit64(value string) (uint64, error) {
	NewValue, err := strconv.ParseFloat(value, 64)
	return uint64(NewValue), err
}

func StringToInt(value string) int8 {
	newValue, _ := strconv.ParseInt(value, 10, 8)
	return int8(newValue)
}
func StringToInt32(value string) (int32, error) {
	newValue, err := strconv.ParseInt(value, 10, 32)
	return int32(newValue), err
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
func IntT64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}
func IntToString(value int8) string {
	val := string(fmt.Sprint(value))
	return val
}

func MakePaginate(totalItems int64, itemsPerPage int64, currentPage int64, url string) string {
	totalPages := totalItems / itemsPerPage
	if totalItems%itemsPerPage != 0 {
		totalPages++
	}

	var paginate strings.Builder

	// Always show first page
	if currentPage != 1 {
		paginate.WriteString(fmt.Sprintf(
			`<li class='page-item'><a class='page-link' data-page='1' href='./%s/?page=1'>1</a></li>`,
			url,
		))
	}

	// Calculate range to show (current page ±1)
	startPage := currentPage - 1
	if startPage < 2 {
		startPage = 2
	}

	endPage := currentPage + 1
	if endPage > totalPages-1 {
		endPage = totalPages - 1
	}

	// Add ellipsis if needed
	if startPage > 2 {
		paginate.WriteString(`<li class='page-item disabled'><span class='page-link'>...</span></li>`)
	}

	// Show pages in range
	for i := startPage; i <= endPage; i++ {
		if i == currentPage {
			paginate.WriteString(fmt.Sprintf(
				`<li class='page-item active'><a class='page-link' data-page='%d' href='./%s/?page=%d'>%d</a></li>`,
				i, url, i, i,
			))
		} else {
			paginate.WriteString(fmt.Sprintf(
				`<li class='page-item'><a class='page-link' data-page='%d' href='./%s/?page=%d'>%d</a></li>`,
				i, url, i, i,
			))
		}
	}

	// Add ellipsis if needed
	if endPage < totalPages-1 {
		paginate.WriteString(`<li class='page-item disabled'><span class='page-link'>...</span></li>`)
	}

	// Always show last page if different from first
	if totalPages > 1 {
		if currentPage != totalPages {
			paginate.WriteString(fmt.Sprintf(
				`<li class='page-item'><a class='page-link' data-page='%d' href='./%s/?page=%d'>%d</a></li>`,
				totalPages, url, totalPages, totalPages,
			))
		}
	}

	return paginate.String()
}
func MakeinventoryPaginate(totalItems int64, itemsPerPage int64, currentPage int64, url string, inventoryID int32) string {
	totalPages := totalItems / itemsPerPage
	if totalItems%itemsPerPage != 0 {
		totalPages++
	}

	var paginate strings.Builder

	// Helper function to build consistent URLs
	buildURL := func(page int64) string {
		return fmt.Sprintf("./%s/?page=%d&inventory=%d", url, page, inventoryID)
	}

	// Always show first page if not current
	if currentPage != 1 {
		paginate.WriteString(fmt.Sprintf(
			`<li class='page-item'><a class='page-link' data-page='1' href='%s'>1</a></li>`,
			buildURL(1),
		))
	}

	// Calculate range to show (current page ±1)
	startPage := currentPage - 1
	if startPage < 2 {
		startPage = 2
	}

	endPage := currentPage + 1
	if endPage > totalPages-1 {
		endPage = totalPages - 1
	}

	// Add ellipsis if needed
	if startPage > 2 {
		paginate.WriteString(`<li class='page-item disabled'><span class='page-link'>...</span></li>`)
	}

	// Show pages in range
	for i := startPage; i <= endPage; i++ {
		if i == currentPage {
			paginate.WriteString(fmt.Sprintf(
				`<li class='page-item active'><a class='page-link' data-page='%d' href='%s'>%d</a></li>`,
				i, buildURL(i), i,
			))
		} else {
			paginate.WriteString(fmt.Sprintf(
				`<li class='page-item'><a class='page-link' data-page='%d' href='%s'>%d</a></li>`,
				i, buildURL(i), i,
			))
		}
	}

	// Add ellipsis if needed
	if endPage < totalPages-1 {
		paginate.WriteString(`<li class='page-item disabled'><span class='page-link'>...</span></li>`)
	}

	// Always show last page if different from first
	if totalPages > 1 && currentPage != totalPages {
		paginate.WriteString(fmt.Sprintf(
			`<li class='page-item'><a class='page-link' data-page='%d' href='%s'>%d</a></li>`,
			totalPages, buildURL(totalPages), totalPages,
		))
	}

	return paginate.String()
}
func MakeRandValue() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letter := letters[rand.Intn(len(letters))]
	num := rand.Intn(10000)
	uniqueString := fmt.Sprintf("%c%05d", letter, num)

	return uniqueString
}
func BitAnd(a, b int16) bool {
	return a&b != 0
}
