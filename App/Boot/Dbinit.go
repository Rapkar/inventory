package Boot

import (
	"fmt"
	"inventory/App/Utility"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {

	dsn := "root:0311121314@tcp(127.0.0.1:3306)/Inventory?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Inventory Has problem With connect to Database")
	}

	return db

}

func Init() {
	if !DB().Migrator().HasTable(Users{}) {
		DB().Migrator().CreateTable(Users{})
		a, _ := Utility.HashPassword("0000")
		User := Users{Name: "hossein Soltanian", Email: "hosseinbidar7@gmail.com", Password: a, Role: "Admin", Phonenumber: "09125174854"}
		DB().Create(&User)
	} else {
		fmt.Println("users table found.# ")
	}

	//Migrate Inventory DB and Smaple row

	if !DB().Migrator().HasTable(Inventory{}) {
		DB().Migrator().CreateTable(Inventory{})
		Inventory := Inventory{Name: "ایزوگام شرق", Number: "10", RolePrice: 99250, MeterPrice: 102500, Count: 100, InventoryNumber: 1}
		DB().Create(&Inventory)
	} else {
		fmt.Println("Inventory table found. #")
	}

	//Migrate Inventory DB and Smaple row
	ExportProduct := []ExportProducts{}
	if !DB().Migrator().HasTable(ExportProducts{}) {
		DB().Migrator().CreateTable(ExportProducts{})
		ExportProduct = []ExportProducts{{Name: "ایزوگام شرق", Number: "10", RolePrice: 99250, MeterPrice: 102500, Count: 100, InventoryNumber: 1, TotalPrice: 2000000, Meter: 10}}
		DB().Create(&ExportProduct)
	} else {
		fmt.Println("ExportProducts table found. #")
	}

	//Migrate Inventory DB and Smaple row

	if !DB().Migrator().HasTable(Export{}) {
		DB().Migrator().CreateTable(Export{})
		Export := Export{Name: "رضا توانگر", Number: "9283422", Phonenumber: "09199656725", Address: "کرج -کرج=-ایران -سیسی", TotalPrice: 10000000, Tax: 10, ExportProducts: ExportProduct, InventoryNumber: 1, CreatedAt: Utility.CurrentTime()}
		DB().Create(&Export)
	} else {
		fmt.Println("Export table found. #")
	}
}
