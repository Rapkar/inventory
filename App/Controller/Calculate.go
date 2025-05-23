package controller

import (
	"fmt"
	"inventory/App/Boot"
	"log"
	"reflect"
)

func InventoryCalculation(id map[int64]int64) {

	for ix, val := range id {
		var count int64
		Boot.DB().Model(&Boot.Inventory{}).Select("count").Where("ID = ?", ix).Scan(&count)
		fmt.Println(reflect.TypeOf(val), val, reflect.TypeOf(count), count)
		res := count - val
		Boot.DB().Model(&Boot.Inventory{}).Where("ID = ?", ix).Update("count", res)
	}

	db := Boot.DB()
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("خطا در دریافت اتصال دیتابیس:", err)
	}
	defer sqlDB.Close()
}
