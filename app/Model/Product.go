package model

import "inventory/boot"

func GetAllProductsByInventory(inventory int32) []boot.Inventory {
	var products []boot.Inventory
	switch inventory {
	case 1:
		boot.DB().Model(&boot.Inventory{}).Select("*").Where("inventory_number= ? ", "1").Scan(&products)
	case 2:
		boot.DB().Model(&boot.Inventory{}).Select("*").Where("inventory_number = ? ", "2").Scan(&products)
	}

	return products
}
func GetProductById(id int) []boot.Inventory {
	var products []boot.Inventory
	boot.DB().Model(&boot.Inventory{}).Select("*").Where("id= ? ", id).Scan(&products)

	return products
}
