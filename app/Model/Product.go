package model

import "inventory/boot"

func GetAllProductsByInventory(inventory int32) []boot.Inventory {
	var products []boot.Inventory
	switch inventory {
	case 1:
		boot.DB().Model(&boot.Users{}).Select("*").Where("inventory = ? ", "1").Scan(&products)
	case 2:
		boot.DB().Model(&boot.Users{}).Select("*").Where("inventory = ? ", "2").Scan(&products)
	}

	return products
}
