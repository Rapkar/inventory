package Model

import (
	"inventory/App/Boot"
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
