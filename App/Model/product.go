package Model

import (
	"fmt"
	"inventory/App/Boot"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProductsByInventory(inventory int32) []Boot.Product {
	var products []Boot.Product

	Boot.DB().Model(&Boot.Product{}).Select("*").Where("inventory_id= ? ", inventory).Scan(&products)

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return products
}
func GetAllProductsByInventoryAndPaginate(offset int, limit int, inventory int32) []Boot.Product {
	var products []Boot.Product

	Boot.DB().Model(&Boot.Product{}).Select("*").Where("inventory_id= ? ", inventory).Offset(offset).
		Limit(limit).Scan(&products)

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
	return products
}
func GetProductById(id int) []Boot.Product {
	var products []Boot.Product
	Boot.DB().Model(&Boot.Product{}).Select("*").Where("id= ? ", id).Scan(&products)
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
	fmt.Println("remove", ProductID)
	result := Boot.DB().Delete(&Boot.Product{}, ProductID)
	if result.RowsAffected == 0 {
		// if no rows were affected, the deletion failed
		return true
	}
	fmt.Println("remove", result.RowsAffected)

	return true
}
