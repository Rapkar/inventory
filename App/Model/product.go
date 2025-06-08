package Model

import (
	"errors"
	"fmt"
	"inventory/App/Boot"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllInventories() []Boot.Inventory {
	var inventories []Boot.Inventory
	result := Boot.DB().Model(&Boot.Inventory{}).Scan(&inventories)
	if result.Error != nil {
		log.Fatal("err in GetAllInventories ", result.Error)
	}
	return inventories
}
func DeleteInventory(ID uint64) bool {
	result := Boot.DB().Delete([]Boot.Inventory{}, ID)
	if result.RowsAffected == 0 {
		return true
	}

	return true
}
func UpdateInventory(ID uint64, updateData string) error {
	result := Boot.DB().Model(&Boot.Inventory{}).Where("id = ?", ID).Update("name", updateData)
	if result.Error != nil {
		return fmt.Errorf("error updating inventory: %v", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no inventory found with ID %d", ID)
	}
	return nil
}
func CreateInventory(name string) (*Boot.Inventory, error) {
	if strings.TrimSpace(name) == "" {
		return nil, fmt.Errorf("نام انبار نمی‌تواند خالی باشد")
	}

	inventory := &Boot.Inventory{
		Name: name,
	}

	result := Boot.DB().Create(inventory)
	if result.Error != nil {
		return nil, fmt.Errorf("خطا در ایجاد انبار: %v", result.Error)
	}

	return inventory, nil
}
func GetInventoryByID(ID uint64) (*Boot.Inventory, error) {
	var inventory Boot.Inventory
	result := Boot.DB().First(&inventory, ID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("انبار با ID %d یافت نشد", ID)
		}
		return nil, result.Error
	}

	return &inventory, nil
}
func GetAllProductsByInventory(inventory int32) []Boot.Product {
	var products []Boot.Product

	err := Boot.DB().Model(&Boot.Product{}).Select("*").Where("inventory_id= ? ", inventory).Scan(&products)

	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	return products
}
func GetAllProductsWithInventory() []Boot.Product {
	var products []Boot.Product

	// استفاده از Preload برای لود خودکار رابطه Inventory
	err := Boot.DB().Preload("Inventory").Find(&products).Error
	if err != nil {
		log.Println("خطا در دریافت محصولات:", err)
	}
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
	result := Boot.DB().Delete(&Boot.Product{}, ProductID)
	if result.RowsAffected == 0 {
		// if no rows were affected, the deletion failed
		return true
	}

	return true
}
