package Model

import (
	"fmt"
	"inventory/App/Boot"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProductsByInventory(inventory int32) []Boot.Inventory {
	var products []Boot.Inventory
	switch inventory {
	case 1:
		Boot.DB().Model(&Boot.Inventory{}).Select("*").Where("inventory_number= ? ", "1").Scan(&products)
	case 2:
		Boot.DB().Model(&Boot.Inventory{}).Select("*").Where("inventory_number = ? ", "2").Scan(&products)
	}

	return products
}
func GetProductById(id int) []Boot.Inventory {
	var products []Boot.Inventory
	Boot.DB().Model(&Boot.Inventory{}).Select("*").Where("id= ? ", id).Scan(&products)

	return products
}
func RemoveCurrentProduct(c *gin.Context) bool {
	Id := c.Request.URL.Query().Get("product-id")
	fmt.Println(Id)
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
