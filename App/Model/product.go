package Model

import (
	"inventory/App/Boot"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProductsByInventory(inventory int32) []Boot.Inventory {
	var products []Boot.Inventory

	Boot.DB().Model(&Boot.Inventory{}).Select("*").Where("inventory_number= ? ", inventory).Scan(&products)

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return products
}
func GetAllProductsByInventoryAndPaginate(offset int, limit int, inventory int32) []Boot.Inventory {
	var products []Boot.Inventory

	Boot.DB().Model(&Boot.Inventory{}).Select("*").Where("inventory_number= ? ", inventory).Offset(offset).
		Limit(limit).Scan(&products)

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return products
}
func GetProductById(id int) []Boot.Inventory {
	var products []Boot.Inventory
	Boot.DB().Model(&Boot.Inventory{}).Select("*").Where("id= ? ", id).Scan(&products)
	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return products
}
func RemoveCurrentProduct(c *gin.Context) bool {
	Id := c.Request.URL.Query().Get("product-id")
	ProductID, err := strconv.ParseUint(Id, 10, 64)
	if err != nil {
		// handle the error
		return false
	}
	result := Boot.DB().Delete(&Boot.Inventory{}, ProductID)
	if result.RowsAffected == 0 {
		// if no rows were affected, the deletion failed
		return false
	}
	return true
}
