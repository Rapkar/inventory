package Model

import (
	"inventory/App/Boot"
	"inventory/App/Utility"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllExports() []Boot.EscapeExport {
	var Export []Boot.Export
	var EscapeExport []Boot.EscapeExport
	Boot.DB().Model(&Boot.Export{}).Select("*").Scan(&Export)
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()

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
		escapeExport.TotalPrice = Utility.IntT64ToString(value.TotalPrice)
		// Add other fields here...
		EscapeExport[i] = escapeExport
	}

	return EscapeExport
}

func GetAllExportsByPaginate(offset int, limit int) []Boot.EscapeExport {
	var Export []Boot.EscapeExport
	var EscapeExport []Boot.EscapeExport
	Boot.DB().Model(&Boot.Export{}).Select("*").Offset(offset).Limit(limit).Scan(&Export)
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()

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
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return Payments
}

type PaymentWithExport struct {
	Boot.Payments
	ExportNumber string `json:"export_number"`
}

func GetAllPaymentsWithExportNumber(offset int, limit int, status string) []PaymentWithExport {
	var result []PaymentWithExport

	query := Boot.DB().Table("payments").
		Select("payments.*, exports.number as export_number").
		Joins("LEFT JOIN exports ON exports.id = payments.export_id").
		Order("payments.id DESC").
		Offset(offset).
		Limit(limit)

	// Add status filter if provided and valid
	if status != "" && (status == "pending" || status == "rejected" || status == "collected") {
		query = query.Where("payments.status = ?", status)
	}

	query.Scan(&result)

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return result
}
func GetPaymentNumberById(c *gin.Context) []Boot.Payments {
	Id := c.Request.URL.Query().Get("ExportId")
	PaymentID, _ := strconv.ParseUint(Id, 10, 64)
	var Payment []Boot.Payments
	Boot.DB().Model(&Boot.Export{}).Select("number").Where("id = ?", PaymentID).Scan(&Payment)
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return Payment
}
func GetExportById(c *gin.Context) ([]Boot.EscapeExport, []Boot.EscapeExportProducts) {
	Id := c.Request.URL.Query().Get("ExportId")
	ExportID, _ := strconv.ParseUint(Id, 10, 64)
	var Export []Boot.Export

	Boot.DB().Model(&Boot.Export{}).Where("id = ?", ExportID).Scan(&Export)
	var ExportProducts []Boot.ExportProducts
	var EscapeExport []Boot.EscapeExport

	EscapeExport = make([]Boot.EscapeExport, len(Export))
	for i, value := range Export {
		var escapeExport Boot.EscapeExport
		escapeExport.Name = value.Name
		escapeExport.Address = value.Address
		escapeExport.Number = value.Number
		escapeExport.Phonenumber = value.Phonenumber
		escapeExport.Tax = value.Tax
		escapeExport.Describe = value.Describe
		escapeExport.InventoryNumber = value.InventoryNumber
		escapeExport.ExportProducts = value.ExportProducts
		escapeExport.CreatedAt = value.CreatedAt
		escapeExport.TotalPrice = Utility.IntT64ToString(value.TotalPrice)
		// Add other fields here...
		EscapeExport[i] = escapeExport
	}

	for _, e := range Export {
		Boot.DB().Model(&Boot.ExportProducts{}).Where("export_id = ?", uint64(e.ID)).Find(&ExportProducts)

	}
	var EscapeExportProducts []Boot.EscapeExportProducts
	EscapeExportProducts = make([]Boot.EscapeExportProducts, len(ExportProducts))
	for i, value := range ExportProducts {
		var escapeExport Boot.EscapeExportProducts
		escapeExport.ID = value.ID
		escapeExport.Name = value.Name
		escapeExport.ExportID = value.ExportID
		escapeExport.Number = value.Number
		escapeExport.RolePrice = Utility.IntT64ToString(value.RolePrice)
		escapeExport.MeterPrice = Utility.IntT64ToString(value.MeterPrice)
		escapeExport.InventoryNumber = value.InventoryNumber
		escapeExport.TotalPrice = Utility.IntT64ToString(value.TotalPrice)
		escapeExport.Count = Utility.IntT64ToString(value.Count)
		// Add other fields here...
		EscapeExportProducts[i] = escapeExport
	}
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return EscapeExport, EscapeExportProducts
}
func GetAllExportsByPhoneAndName(searchTerm string) []Boot.EscapeExport {
	var Export []Boot.Export
	var EscapeExport []Boot.EscapeExport

	Boot.DB().Model(&Boot.Export{}).Where("name LIKE ? OR phonenumber LIKE ? OR number LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%", "%"+searchTerm+"%").Find(&Export)
	// Boot.DB().Model(&Boot.Export{}).Where("name = ?", searchTerm).Limit(3).Find(&Export)
	// if len(Export) == 0 {
	// 	return []Boot.EscapeExport{}
	// }

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
		escapeExport.TotalPrice = Utility.IntT64ToString(value.TotalPrice)
		// Add other fields here...
		EscapeExport[i] = escapeExport
	}
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return EscapeExport
}
func GetAllPaymentsByAttribiute(searchTerm string) []PaymentWithExport {
	var result []PaymentWithExport

	query := Boot.DB().Table("payments").
		Select("payments.*, exports.number as export_number").
		Joins("LEFT JOIN exports ON exports.id = payments.export_id").
		Where("payments.created_at LIKE ? OR payments.number LIKE ?", "%"+searchTerm+"%", "%"+searchTerm+"%")

	query.Scan(&result)

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return result
}
func GetCountOfExports() int64 {
	var count int64
	Boot.DB().Model(&[]Boot.Export{}).Find(&[]Boot.Export{}).Count(&count)

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return count
}
func RemoveCurrentExport(c *gin.Context) bool {
	Id := c.Request.URL.Query().Get("ExportId")
	ExportID, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		// handle the error
		return false
	}
	result := Boot.DB().Delete(&Boot.Export{}, ExportID)
	if result.RowsAffected == 0 {
		// if no rows were affected, the deletion failed
		return false
	}

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return true
}

func RemoveCurrentPayments(c *gin.Context) bool {
	Id := c.Request.URL.Query().Get("PaymentId")
	PaymentID, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		// handle the error
		return false
	}
	result := Boot.DB().Delete(&Boot.Payments{}, PaymentID)
	if result.RowsAffected == 0 {
		// if no rows were affected, the deletion failed
		return false
	}

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return true
}

func CheckExportNumberFound(value string) bool {
	var Export []Boot.Export

	Boot.DB().Model(&Boot.Export{}).Where("number LIKE ?", "%"+value+"%").Find(&Export)

	if Export != nil {
		return false
	}

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return true
}
