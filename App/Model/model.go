package Model

import (
	"fmt"
	"inventory/App/Boot"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// get All user by role , arg is [Admin , guest]

func GetAllUsersByRole(role string) []Boot.Users {
	var users []Boot.Users
	db := Boot.DB()

	if err := db.Where("role = ?", role).Find(&users).Error; err != nil {
		log.Println("❌ err in  GetAllUsersByRole", role, ":", err)
	}

	return users
}

func GetUserRoleById(userID uint64) (string, error) {
	var role string
	db := Boot.DB()

	err := db.Model(&Boot.Users{}).
		Select("role").
		Where("id = ?", userID).
		Take(&role).Error

	if err != nil {
		log.Printf("❌ err in GetUserRoleById %d: %v", userID, err)
		return "", err
	}

	return role, nil
}

// get user By email [hosseinbidar7@gmail.com]

func GetUserByEmail(email string) ([]Boot.Users, error) {
	var users []Boot.Users
	db := Boot.DB()

	err := db.Where("email = ?", email).Find(&users).Error
	if err != nil {
		log.Printf("❌ err in GetUserByEmail %s: %v", email, err)
		return nil, err
	}

	return users, nil
}

// get user by id
func GetUserById(userID uint64) (*Boot.Users, error) {
	var user Boot.Users
	db := Boot.DB()

	err := db.First(&user, userID).Error
	if err != nil {
		log.Printf("❌ err in GetUserById %d: %v", userID, err)
		return nil, err
	}

	return &user, nil
}

func GetCurrentUser(c *gin.Context) (*Boot.Users, error) {
	//get id from url
	id := c.DefaultQuery("user-id", "")
	if id == "" {
		return nil, fmt.Errorf("❌ user-id is required")
	}

	//id to unit64
	userIdUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		//err in convert
		return nil, fmt.Errorf("❌ invalid user-id: %v", err)
	}

	//user by id

	currentUser, err := GetUserById(userIdUint)
	if err != nil {
		// if user not found
		return nil, err
	}

	return currentUser, nil
}

func GetCountOfUsers() (int64, error) {
	var count int64
	db := Boot.DB()

	if err := db.Model(&Boot.Users{}).Count(&count).Error; err != nil {
		return 0, fmt.Errorf("❌ error counting users: %v", err)
	}

	return count, nil
}

func GetCountOfProduct(id int32) int64 {
	var count int64
	db := Boot.DB()
	if err := db.Model(&[]Boot.Product{}).Where("inventory_id= ? ", id).Count(&count).Error; err != nil {
		log.Println("❌  err in  GetCountOfProduct ", err)
		return 0
	}
	return count
}
func GetAllUsersByPaginate(offset int, limit int, role string) []Boot.Users {
	var users []Boot.Users
	db := Boot.DB()

	err := db.Model(&Boot.Users{}).
		Where("role = ?", role).
		Offset(offset).
		Limit(limit).
		Find(&users).Error

	if err != nil {
		log.Printf("❌ err in get GetAllUsersByPaginate with role %s: %v", role, err)
	}

	return users
}

func RemoveCurrentUser(c *gin.Context) bool {
	// get user id from url
	id := c.DefaultQuery("user-id", "")
	if id == "" {
		//if id is empty
		log.Println("❌ user-id is required")
		return false
	}

	//convert user id to unit
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		//
		log.Printf("❌ invalid user-id: %v", err)
		return false
	}

	//delete user by id
	result := Boot.DB().Delete(&Boot.Users{}, userID)
	if result.Error != nil {
		// if delete faild
		log.Printf("❌ error deleting user with id %d: %v", userID, result.Error)
		return false
	}

	if result.RowsAffected == 0 {
		// if  RowsAffected faild
		log.Printf("❌ no user found with id %d", userID)
		return false
	}

	// if success
	return true
}

type UserFullDetails struct {
	User           Boot.Users
	Exports        []Boot.Export
	ExportProducts []Boot.ExportProducts
	Payments       []Boot.Payments
}

func GetUserFullDetailsByID(userID uint64) (UserFullDetails, error) {
	var result UserFullDetails
	db := Boot.DB()

	// دریافت اطلاعات کاربر
	if err := db.Where("id = ?", userID).First(&result.User).Error; err != nil {
		log.Printf("❌ خطا در یافتن کاربر با ID=%d : %v", userID, err)
		return result, err
	}

	// دریافت فاکتورهای کاربر
	if err := db.Where("user_id = ?", userID).Find(&result.Exports).Error; err != nil {
		log.Printf("❌ خطا در دریافت فاکتورهای کاربر با ID=%d : %v", userID, err)
		return result, err
	}

	// دریافت محصولات و پرداخت‌ها
	var exportIDs []uint64
	for _, export := range result.Exports {
		exportIDs = append(exportIDs, export.ID)
	}

	if len(exportIDs) > 0 {
		if err := db.Where("export_id IN ?", exportIDs).Find(&result.ExportProducts).Error; err != nil {
			log.Printf("❌ خطا در دریافت محصولات فاکتورهای کاربر با ID=%d : %v", userID, err)
			return result, err
		}

		if err := db.Where("export_id IN ?", exportIDs).Find(&result.Payments).Error; err != nil {
			log.Printf("❌ خطا در دریافت پرداخت‌های کاربر با ID=%d : %v", userID, err)
			return result, err
		}
	}

	return result, nil
}
func GetTotalPayment(c *gin.Context) (float64, error) {

	var total float64
	if err := Boot.DB().Model(&Boot.Payments{}).Select("SUM(total_price)").Scan(&total).Error; err != nil {
		log.Fatal("GetTotalPayment err:", err)
		return 0, err
	}

	return total, nil
}
func GetTotalIncome(c *gin.Context) (float64, error) {

	var totalIncome float64
	err := Boot.DB().Model(&Boot.Payments{}).
		Select("SUM(total_price)").
		Where("status = ?", "collected").
		Scan(&totalIncome).Error

	if err != nil {
		log.Printf("GetTotalIncome error: %v", err)
		return 0, err
	}

	return totalIncome, nil
}
func GetTotalPrices(c *gin.Context) (float64, error) {

	var TotalPrice float64
	err := Boot.DB().Model(&Boot.Export{}).
		Select("SUM(total_price)").
		Scan(&TotalPrice).Error

	if err != nil {
		log.Printf("GetTotalIncome error: %v", err)
		return 0, err
	}

	return TotalPrice, nil
}
func GetTotalPaymentByDateRange(c *gin.Context, fromDate, toDate string) (float64, error) {
	var total float64
	err := Boot.DB().Model(&Boot.Payments{}).
		Select("SUM(total_price)").
		Where("created_at BETWEEN ? AND ?", fromDate, toDate).
		Scan(&total).Error

	if err != nil {
		log.Printf("GetTotalPaymentByDateRange error: %v", err)
		return 0, err
	}

	return total, nil
}
func GetTotalIncomeByDateRange(c *gin.Context, fromDate, toDate string) (float64, error) {
	var totalIncome float64
	err := Boot.DB().Model(&Boot.Payments{}).
		Select("SUM(total_price)").
		Where("status = ? AND created_at BETWEEN ? AND ?", "collected", fromDate, toDate).
		Scan(&totalIncome).Error

	if err != nil {
		log.Printf("GetTotalIncomeByDateRange error: %v", err)
		return 0, err
	}

	return totalIncome, nil
}
func GetTotalPriceByDateRange(c *gin.Context, fromDate, toDate string) (float64, error) {
	var totalIncome float64
	err := Boot.DB().Model(&Boot.Export{}).
		Select("SUM(total_price)").
		Where("created_at BETWEEN ? AND ?", fromDate, toDate).
		Scan(&totalIncome).Error

	if err != nil {
		log.Printf("GetTotalPriceByDateRange error: %v", err)
		return 0, err
	}

	return totalIncome, nil
}

func ChangeExportDraftStatus(id uint64, value bool) bool {
	// result := Boot.DB().Model(&Boot.Export{}).Where("id = ?", id).Update("draft", value)

	// revece value
	Draftvalue := true

	if value {
		Draftvalue = false
	}
	result := Boot.DB().Model(&[]Boot.Export{}).Where("id = ?", id).Update("draft", Draftvalue)
	if result.Error != nil {
		log.Printf("❌ خطا در بروزرسانی وضعیت پیش‌نویس فاکتور ID=%d : %v", id, result.Error)
		return false
	}

	if result.RowsAffected == 0 {
		log.Printf("⚠️ هیچ رکوردی با ID=%d یافت نشد", id)
		return false
	}

	return true
}
