package Model

import (
	"fmt"
	"inventory/App/Boot"
	"inventory/App/Utility"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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
		escapeExport.InventoryNumber = int32(value.ProductID)
		escapeExport.ExportProducts = value.ExportProducts
		escapeExport.CreatedAt = value.CreatedAt
		escapeExport.TotalPrice = value.TotalPrice
		// Add other fields here...
		EscapeExport[i] = escapeExport
	}

	return EscapeExport
}

func GetAllExportsByPaginate(offset int, limit int) []Boot.EscapeExport {
	var Export []Boot.EscapeExport
	var EscapeExport []Boot.EscapeExport
	db := Boot.DB()

	if err := db.Model(&Boot.Export{}).Select("*").Offset(offset).Limit(limit).Scan(&Export).Error; err != nil {
		log.Println("❌ err in GetAllExportsByPaginate", err)
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

func GetAllPaymentsWithExportNumber(offset int, limit int, status string) ([]Boot.PaymentWithExport, error) {
	var result []Boot.PaymentWithExport

	query := Boot.DB().Table("payments").
		Select("payments.*, exports.number as export_number").
		Joins("LEFT JOIN exports ON exports.id = payments.export_id").
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
func GetAllPaymentsWithExportNumberByUserId(offset int, limit int, status string, user_id uint64) ([]Boot.PaymentWithExport, error) {
	var result []Boot.PaymentWithExport

	query := Boot.DB().Table("payments").
		Where("payments.user_id = ?", user_id).
		Select("payments.*, exports.number as export_number").
		Joins("LEFT JOIN exports ON exports.id = payments.export_id").
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
	fmt.Println(result)
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
			InventoryNumber: int32(export.ProductID),
			ExportProducts:  export.ExportProducts,
			CreatedAt:       export.CreatedAt,
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
			Number:          exportProduct.Number,
			RolePrice:       exportProduct.RolePrice,
			MeterPrice:      exportProduct.MeterPrice,
			InventoryNumber: int32(exportProduct.ID),
			TotalPrice:      exportProduct.TotalPrice,
			Count:           Utility.IntT64ToString(exportProduct.Count),
			Meter:           exportProduct.Meter,
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
			InventoryNumber: int32(value.ProductID),
			ExportProducts:  value.ExportProducts,
			CreatedAt:       value.CreatedAt,
			TotalPrice:      value.TotalPrice,
			// اگر فیلد دیگه‌ای هست اضافه کن
		}
	}

	return result
}

func GetAllPaymentsByAttribiute(searchTerm string) []Boot.PaymentWithExport {
	var result []Boot.PaymentWithExport

	db := Boot.DB()
	err := db.Table("payments").
		Select("payments.*, exports.number as export_number").
		Joins("LEFT JOIN exports ON exports.id = payments.export_id").
		Where("payments.created_at LIKE ? OR payments.number LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%").
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
	id := c.Query("ExportId") // ساده‌تر از c.Request.URL.Query().Get
	exportID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println("❌ Export Convert has conflict:", err)
		return false
	}

	result := Boot.DB().Delete(&Boot.Export{}, exportID)
	if result.Error != nil {
		log.Println("❌ err in Export delete :", result.Error)
		return false
	}

	if result.RowsAffected == 0 {
		log.Println("⚠️ Payment record id not found")
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
