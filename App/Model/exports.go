package Model

import (
	"errors"
	"fmt"
	"inventory/App/Boot"
	"inventory/App/Utility"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllExports() []Boot.EscapeExport {
	var Export []Boot.Export
	var EscapeExport []Boot.EscapeExport
	db := Boot.DB()

	if err := db.Model(&Boot.Export{}).Select("*").Scan(&Export).Error; err != nil {
		log.Println("❌ err in GetAllExports", err)
	}
	EscapeExport = make([]Boot.EscapeExport, len(Export))

	for i, value := range Export {
		var escapeExport Boot.EscapeExport

		escapeExport.ID = value.ID
		escapeExport.Name = value.Name
		escapeExport.Address = value.Address
		escapeExport.Number = value.Number
		escapeExport.Phonenumber = value.Phonenumber
		escapeExport.Tax = value.Tax
		escapeExport.InventoryNumber = int32(value.InventoryID)
		escapeExport.ExportProducts = value.ExportProducts
		escapeExport.CreatedAt = value.CreatedAt
		escapeExport.TotalPrice = value.TotalPrice
		// Add other fields here...
		EscapeExport[i] = escapeExport
	}

	return EscapeExport
}

func GetAllExportsByPaginate(offset int, limit int, draft bool) []Boot.EscapeExport {
	var Export []Boot.EscapeExport
	var EscapeExport []Boot.EscapeExport
	db := Boot.DB()
	if draft {
		if err := db.Model(&Boot.Export{}).Select("*").Where("draft", true).Offset(offset).Limit(limit).Order("id DESC").Scan(&Export).Error; err != nil {
			log.Println("❌ err in GetAllExportsByPaginate", err)
		}
	} else {
		if err := db.Model(&Boot.Export{}).Select("*").Where("draft", false).Offset(offset).Limit(limit).Order("id DESC").Scan(&Export).Error; err != nil {
			log.Println("❌ err in GetAllExportsByPaginate", err)
		}
	}

	EscapeExport = make([]Boot.EscapeExport, len(Export))

	for i, value := range Export {
		var escapeExport Boot.EscapeExport
		escapeExport.ID = value.ID
		escapeExport.Name = value.Name
		escapeExport.Address = value.Address
		escapeExport.Number = value.Number
		escapeExport.Phonenumber = value.Phonenumber
		escapeExport.Tax = value.Tax
		escapeExport.InventoryNumber = value.InventoryNumber
		escapeExport.ExportProducts = value.ExportProducts
		escapeExport.CreatedAt = value.CreatedAt
		escapeExport.TotalPrice = value.TotalPrice
		// Add other fields here...
		EscapeExport[i] = escapeExport
	}

	return EscapeExport
}
func GetAllPaymentsByPaginate(offset int, limit int) []Boot.Payments {
	var Payments []Boot.Payments
	Boot.DB().Model(&Boot.Payments{}).
		Select("*").
		Preload("Export").
		Order("id DESC"). // Sort by creation date, newest first
		Offset(offset).
		Limit(limit).
		Scan(&Payments)

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("❌ err in GetAllPaymentsByPaginate ", err)
	}
	defer sqlDB.Close()
	return Payments
}

func GetAllPaymentsWithExportNumberAndUser(offset int, limit int, status string) ([]Boot.PaymentWithExportAndUser, error) {
	var result []Boot.PaymentWithExportAndUser

	query := Boot.DB().Table("payments").
		Select("payments.*, exports.number as export_number, users.name as user_name").
		Joins("LEFT JOIN exports ON exports.id = payments.export_id").
		Joins("LEFT JOIN users ON users.id = payments.user_id"). // Assuming there's a user_id foreign key in payments
		Order("payments.id DESC").
		Offset(offset).
		Limit(limit)

	if status != "" && (status == "pending" || status == "rejected" || status == "collected") {
		query = query.Where("payments.status = ?", status)
	}

	err := query.Find(&result).Error
	if err != nil {
		log.Println("❌ Error fetching payments:", err)
		return nil, fmt.Errorf("Error fetching payments: %v", err)
	}

	return result, nil
}
func GetAllPaymentsWithExportNumberByUserId(offset int, limit int, status string, user_id uint64) ([]Boot.PaymentWithExportAndUser, error) {
	var result []Boot.PaymentWithExportAndUser

	query := Boot.DB().Table("payments").
		Where("payments.user_id = ?", user_id).
		Select("payments.*, exports.number as export_number, users.name as user_name").
		Joins("LEFT JOIN exports ON exports.id = payments.export_id").
		Joins("LEFT JOIN users ON users.id = payments.user_id").
		Order("payments.id DESC").
		Offset(offset).
		Limit(limit)

	if status != "" && (status == "pending" || status == "rejected" || status == "collected") {
		query = query.Where("payments.status = ?", status)
	}

	err := query.Find(&result).Error
	if err != nil {
		log.Println("❌ Error fetching payments:", err)
		return nil, fmt.Errorf("Error fetching payments: %v", err)
	}
	return result, nil
}

func GetTotalprice(userid uint64) (float64, error) {
	var result float64

	// Execute the query and scan the result
	err := Boot.DB().Table("payments").
		Where("user_id = ?", userid). // Fixed variable name from user_id to userid
		Select("SUM(total_price)").   // Assuming you want to sum all payments, otherwise you might need a different approach
		Scan(&result).                // Need to actually execute the query and scan the result
		Error

	if err != nil {
		return 0, err // Return the error to handle it properly
	}

	return result, nil
}
func GetCountOfUserExports(userid uint64) (int64, error) {
	var result int64

	// Execute the query and scan the result
	err := Boot.DB().Table("exports").
		Where("user_id = ?", userid). // Fixed variable name from user_id to userid
		Select("COUNT(total_price)"). // Assuming you want to sum all payments, otherwise you might need a different approach
		Scan(&result).                // Need to actually execute the query and scan the result
		Error

	if err != nil {
		return 0, err // Return the error to handle it properly
	}

	return result, nil
}
func GetUserTotalPrice(userid uint64) (float64, error) {
	var result float64

	// Execute the query and scan the result
	err := Boot.DB().Table("exports").
		Where("user_id = ?", userid). // Fixed variable name from user_id to userid
		Select("SUM(total_price)").   // Assuming you want to sum all payments, otherwise you might need a different approach
		Scan(&result).                // Need to actually execute the query and scan the result
		Error

	if err != nil {
		return 0, err // Return the error to handle it properly
	}

	return result, nil
}
func GetUserTotalPaid(userid uint64) (float64, error) {
	var result *float64 // Use a pointer to handle NULL values

	// Execute the query and scan the result
	err := Boot.DB().Table("payments").
		Where("user_id = ?", userid).
		Select("SUM(total_price) as total").
		Scan(&result). // Scan into a pointer
		Error

	if err != nil {
		return 0, err
	}

	// If result is nil (NULL in database), return 0
	if result == nil {
		return 0, nil
	}

	return *result, nil
}
func GetPaymentNumberById(c *gin.Context) ([]Boot.Payments, error) {
	// دریافت PaymentId از Query Param
	Id := c.DefaultQuery("PaymentId", "")
	PaymentID, err := strconv.ParseUint(Id, 10, 64)
	if err != nil || PaymentID == 0 {
		log.Println("❌ Invalid PaymentId:", err)
		return nil, fmt.Errorf("Invalid PaymentId")
	}

	// جستجو در جدول Payments برای دریافت شماره پرداخت
	var payments []Boot.Payments
	err = Boot.DB().Model(&Boot.Payments{}).Where("id = ?", PaymentID).Find(&payments).Error
	if err != nil || len(payments) == 0 {
		log.Println("❌ Error fetching payment by ID:", err)
		return nil, fmt.Errorf("Payment not found")
	}

	return payments, nil
}

const (
	RollColumn   int16 = 1 << 0
	MeterColumn  int16 = 1 << 1
	WeightColumn int16 = 1 << 2
	CountColumn  int16 = 1 << 3
	BarrelColumn int16 = 1 << 4
)

func GetExportProductsColumns(products []Boot.EscapeExportProducts) int16 {
	var columns int16 = 0

	for _, p := range products {
		if p.Roll > 0 {
			columns |= RollColumn
		}
		if p.Meter > 0 {
			columns |= MeterColumn
		}
		if p.Weight > 0 {
			columns |= WeightColumn
		}
		if p.Count > 0 {
			columns |= CountColumn
		}
		if p.Barrel > 0 {
			columns |= BarrelColumn
		}
	}
	return columns
}
func GetExportById2(c *gin.Context) (Boot.EscapeExport, []Boot.EscapeExportProducts, []Boot.Payments) {
	// دریافت ID از Query Param
	Id := c.DefaultQuery("ExportId", "")
	ExportID, err := strconv.ParseUint(Id, 10, 64)
	if err != nil || ExportID == 0 {
		log.Println("❌ Invalid ExportId:", err)
		return Boot.EscapeExport{}, nil, nil
	}

	var export Boot.Export
	// جستجو در جدول Export با پیش‌بارگذاری روابط
	err = Boot.DB().Preload("Payments").Where("id = ?", ExportID).First(&export).Error
	if err != nil {
		log.Println("❌ Error fetching export by ID:", err)
		return Boot.EscapeExport{}, nil, nil
	}

	escapeExport := Boot.EscapeExport{
		ID:              export.ID,
		Name:            export.Name,
		Address:         export.Address,
		Number:          export.Number,
		Phonenumber:     export.Phonenumber,
		Tax:             export.Tax,
		Describe:        export.Describe,
		InventoryNumber: int32(export.InventoryID),
		ExportProducts:  export.ExportProducts,
		CreatedAt:       export.CreatedAt,
		CreatorName:     export.CreatorName,
		TotalPrice:      export.TotalPrice,
		Draft:           export.Draft,
	}

	// جستجو برای محصولات مربوط به Export
	var exportProducts []Boot.ExportProducts
	err = Boot.DB().Where("export_id = ?", ExportID).Find(&exportProducts).Error
	if err != nil {
		log.Println("❌ Error fetching export products:", err)
		return escapeExport, nil, export.Payments
	}

	var escapeExportProducts []Boot.EscapeExportProducts
	for _, exportProduct := range exportProducts {
		escapeExportProduct := Boot.EscapeExportProducts{
			ID:              exportProduct.ID,
			Name:            exportProduct.Name,
			ExportID:        exportProduct.ExportID,
			InventoryNumber: int32(exportProduct.InventoryID),
			TotalPrice:      exportProduct.TotalPrice,
			Roll:            exportProduct.Roll,
			Meter:           exportProduct.Meter,
			Count:           exportProduct.Count,
			Weight:          exportProduct.Weight,
			Barrel:          exportProduct.Barrel,
			RollePrice:      exportProduct.RollePrice,
			MeterPrice:      exportProduct.MeterPrice,
			CountPrice:      exportProduct.CountPrice,
			ProductID:       exportProduct.ProductID,
			BarrelPrice:     exportProduct.BarrelPrice,
		}
		escapeExportProducts = append(escapeExportProducts, escapeExportProduct)
	}

	return escapeExport, escapeExportProducts, export.Payments
}
func GetExportById(c *gin.Context) ([]Boot.EscapeExport, []Boot.EscapeExportProducts) {
	// دریافت ID از Query Param
	Id := c.DefaultQuery("ExportId", "")
	ExportID, err := strconv.ParseUint(Id, 10, 64)
	if err != nil || ExportID == 0 {
		log.Println("❌ Invalid ExportId:", err)
		return nil, nil
	}

	var exports []Boot.Export
	// جستجو در جدول Export
	err = Boot.DB().Model(&Boot.Export{}).Where("id = ?", ExportID).Find(&exports).Error
	if err != nil || len(exports) == 0 {
		log.Println("❌ Error fetching export by ID:", err)
		return nil, nil
	}

	var escapeExports []Boot.EscapeExport
	for _, export := range exports {
		escapeExport := Boot.EscapeExport{
			Name:            export.Name,
			Address:         export.Address,
			Number:          export.Number,
			Phonenumber:     export.Phonenumber,
			Tax:             export.Tax,
			Describe:        export.Describe,
			InventoryNumber: int32(export.InventoryID),
			ExportProducts:  export.ExportProducts,
			CreatedAt:       export.CreatedAt,
			CreatorName:     export.CreatorName,
			TotalPrice:      export.TotalPrice,
		}
		escapeExports = append(escapeExports, escapeExport)
	}

	// جستجو برای محصولات مربوط به Export
	var exportProducts []Boot.ExportProducts
	err = Boot.DB().Model(&Boot.ExportProducts{}).Where("export_id = ?", ExportID).Find(&exportProducts).Error
	if err != nil {
		log.Println("❌ Error fetching export products:", err)
		return escapeExports, nil
	}

	var escapeExportProducts []Boot.EscapeExportProducts
	for _, exportProduct := range exportProducts {
		escapeExportProduct := Boot.EscapeExportProducts{
			ID:              exportProduct.ID,
			Name:            exportProduct.Name,
			ExportID:        exportProduct.ExportID,
			InventoryNumber: int32(exportProduct.ID),
			TotalPrice:      exportProduct.TotalPrice,
			Roll:            exportProduct.Roll,
			Meter:           exportProduct.Meter,
			Count:           exportProduct.Count,
			Weight:          exportProduct.Weight,
			Barrel:          exportProduct.Barrel,
			RollePrice:      exportProduct.RollePrice,
			MeterPrice:      exportProduct.MeterPrice,
			CountPrice:      exportProduct.CountPrice,
			BarrelPrice:     exportProduct.BarrelPrice,
		}
		escapeExportProducts = append(escapeExportProducts, escapeExportProduct)
	}

	return escapeExports, escapeExportProducts
}

func GetAllExportsByPhoneAndName(searchTerm string) []Boot.EscapeExport {
	var exports []Boot.Export
	var result []Boot.EscapeExport

	db := Boot.DB()
	err := db.Model(&Boot.Export{}).
		Where("name LIKE ? OR phonenumber LIKE ? OR number LIKE ?",
			"%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%").
		Find(&exports).Error

	if err != nil {
		log.Println("❌ Error fetching exports:", err)
		return result
	}

	result = make([]Boot.EscapeExport, len(exports))
	for i, value := range exports {
		result[i] = Boot.EscapeExport{
			ID:              value.ID,
			Name:            value.Name,
			Address:         value.Address,
			Number:          value.Number,
			Phonenumber:     value.Phonenumber,
			Tax:             value.Tax,
			InventoryNumber: int32(value.InventoryID),
			ExportProducts:  value.ExportProducts,
			CreatedAt:       value.CreatedAt,
			TotalPrice:      value.TotalPrice,
			// اگر فیلد دیگه‌ای هست اضافه کن
		}
	}

	return result
}

func GetAllPaymentsByAttribiute(searchTerm string) []Boot.PaymentWithExportAndUser {
	var result []Boot.PaymentWithExportAndUser

	db := Boot.DB()
	err := db.Table("payments").
		Select("payments.*, exports.number as export_number, users.name as user_name").
		Joins("LEFT JOIN exports ON exports.id = payments.export_id").
		Joins("LEFT JOIN users ON users.id = payments.user_id").
		Where("payments.created_at LIKE ? OR payments.number LIKE ?  OR users.name LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%").
		Scan(&result).Error

	if err != nil {
		log.Println("❌ Error fetching payments by attribute:", err)
	}

	return result
}

func GetCountOfExports() int64 {
	var count int64

	err := Boot.DB().Model(&Boot.Export{}).Count(&count).Error
	if err != nil {
		log.Println("❌ Error counting exports:", err)
		return 0
	}

	return count
}

func GetCountOfPayments() int64 {
	var count int64

	err := Boot.DB().Model(&Boot.Payments{}).Count(&count).Error
	if err != nil {
		log.Println("❌ Error counting payments:", err)
		return 0
	}

	return count
}

func RemoveCurrentExport(c *gin.Context) bool {
	id := c.Query("ExportId")
	exportID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println("❌ Export ID conversion error:", err)
		return false
	}

	// Start a database transaction
	err = Boot.DB().Transaction(func(tx *gorm.DB) error {
		// 1. First delete related products
		if err := tx.Delete(&Boot.ExportProducts{}, "export_id = ?", exportID).Error; err != nil {
			return fmt.Errorf("failed to delete export products: %v", err)
		}

		// 2. Then delete related payments
		if err := tx.Delete(&Boot.Payments{}, "export_id = ?", exportID).Error; err != nil {
			return fmt.Errorf("failed to delete payments: %v", err)
		}

		// 3. Finally delete the main export record
		result := tx.Delete(&Boot.Export{}, exportID)
		if result.Error != nil {
			return fmt.Errorf("failed to delete export: %v", result.Error)
		}
		if result.RowsAffected == 0 {
			return fmt.Errorf("no export record found with ID %d", exportID)
		}

		return nil // Return nil to commit the transaction
	})

	if err != nil {
		log.Println("❌ Transaction failed:", err)
		return false
	}

	return true
}

func RemoveCurrentPayments(c *gin.Context) bool {
	id := c.Request.URL.Query().Get("PaymentId")
	paymentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println("❌ err in convert Payment id:", err)
		return false
	}

	result := Boot.DB().Delete(&Boot.Payments{}, paymentID)
	if result.Error != nil {
		log.Println("❌ err in remove Payment:", result.Error)
		return false
	}

	if result.RowsAffected == 0 {
		log.Println("⚠️ can't find Payment whith this id")
		return false
	}

	return true
}

func CheckExportNumberFound(value string) bool {
	var exports []Boot.Export

	err := Boot.DB().Model(&Boot.Export{}).Where("number = ?", value).Find(&exports).Error
	if err != nil {
		log.Println("❌ Err in CheckExportNumberFound :", err)
		return false
	}

	return len(exports) > 0
}

func GenerateUniqueExportID() (string, error) {
	const maxAttempts = 10
	for i := 0; i < maxAttempts; i++ {
		uniqueString := Utility.MakeRandValue()
		if !CheckExportNumberFound(uniqueString) {
			return uniqueString, nil
		}
	}
	return "", fmt.Errorf("failed to generate unique export ID after %d attempts", maxAttempts)
}
func GetUsersByNameAndPhone(nameTerm string, phoneTerm string) ([]Boot.ResponseUsers, string) {
	var result []Boot.ResponseUsers
	message := ""
	db := Boot.DB()

	// پاکسازی ورودی‌ها
	nameTerm = strings.TrimSpace(nameTerm)
	phoneTerm = strings.TrimSpace(phoneTerm)

	// بررسی وجود شماره تلفن (بدون در نظر گرفتن نام)
	var phoneExists bool
	if len(phoneTerm) >= 10 {
		var count int64
		db.Model(&Boot.Users{}).Where("phonenumber = ?", phoneTerm).Count(&count)
		phoneExists = count > 0
	}

	// بررسی وجود نام (بدون در نظر گرفتن شماره)
	var nameExists bool
	if len(nameTerm) >= 2 {
		var count int64
		db.Model(&Boot.Users{}).Where("name LIKE ?", "%"+nameTerm+"%").Count(&count)
		nameExists = count > 0
	}

	// ساخت کوئری اصلی بر اساس ورودی‌ها
	query := db.Model(&Boot.Users{})
	if len(nameTerm) >= 2 {
		query = query.Where("name LIKE ?", "%"+nameTerm+"%")
	}
	if len(phoneTerm) >= 10 {
		query = query.Where("phonenumber = ?", phoneTerm)
	}

	// اجرای کوئری
	err := query.Select("id, name, phonenumber, address").Find(&result).Error
	if err != nil {
		log.Println("❌ Error fetching users:", err)
		return nil, "خطا در دریافت اطلاعات"
	}

	// تعیین پیام مناسب
	if len(phoneTerm) >= 10 && phoneExists {
		message = "این شماره تلفن قبلاً ثبت شده است"
	} else if len(nameTerm) >= 2 && nameExists || phoneExists {
		message = "کاربری با این نام قبلاً ثبت شده است"
	} else if len(result) > 0 {
		message = "نتایج جستجو:"
	} else {
		message = "موردی یافت نشد"
	}

	return result, message
}
func GetPaymentNumberByExportId(ExportNumber string) ([]Boot.Payments, []Boot.ExportProducts, error) {
	if ExportNumber == "" {
		return nil, nil, fmt.Errorf("ExportNumber parameter is required")
	}

	var export Boot.Export
	if err := Boot.DB().Model(&Boot.Export{}).
		Where("number = ?", ExportNumber).
		First(&export).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, fmt.Errorf("export with number %s not found", ExportNumber)
		}
		log.Printf("❌ Database error in finding export: %v", err)
		return nil, nil, fmt.Errorf("database error")
	}

	var payments []Boot.Payments
	if err := Boot.DB().Model(&Boot.Payments{}).
		Where("export_id = ?", export.ID).
		Find(&payments).Error; err != nil {

		log.Printf("❌ Database error in finding payments: %v", err)
		return nil, nil, fmt.Errorf("database error")
	}

	var ExportProducts []Boot.ExportProducts
	if err := Boot.DB().Model(&Boot.ExportProducts{}).
		Where("export_id = ?", export.ID).
		Find(&ExportProducts).Error; err != nil {

		log.Printf("❌ Database error in finding export products: %v", err)
		return nil, nil, fmt.Errorf("database error")
	}

	return payments, ExportProducts, nil
}
