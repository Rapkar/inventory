package Utility

import (

	// Model "inventory/app/Model"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	ptime "github.com/yaa110/go-persian-calendar"
	"golang.org/x/crypto/bcrypt"
)

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
func StringToInt64(value string) int64 {
	NewValue, _ := strconv.ParseFloat(value, 64)
	return int64(NewValue)
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
func IntT64ToString(value int64) string {
	return strconv.FormatInt(value, 10)
}
func IntToString(value int8) string {
	val := string(fmt.Sprint(value))
	return val
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
func MakeRandValue() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	letter := letters[rand.Intn(len(letters))]
	num := rand.Intn(10000)
	uniqueString := fmt.Sprintf("%c%05d", letter, num)

	return uniqueString
}

// func ExampleNewPDFGenerator() {

// 	// Create new PDF generator
// 	pdfg, err := NewPDFGenerator()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Set global options
// 	pdfg.Dpi.Set(300)
// 	pdfg.Orientation.Set(OrientationLandscape)
// 	pdfg.Grayscale.Set(true)

// 	// Create a new input page from an URL
// 	page := NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")

// 	// Set options for this page
// 	page.FooterRight.Set("[page]")
// 	page.FooterFontSize.Set(10)
// 	page.Zoom.Set(0.95)

// 	// Add to document
// 	pdfg.AddPage(page)

// 	// Create PDF document in internal buffer
// 	err = pdfg.Create()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Write buffer contents to file on disk
// 	err = pdfg.WriteFile("./simplesample.pdf")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Done")
// 	// Output: Done
// }
