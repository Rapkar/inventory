package boot

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {

	dsn := "root:0311121314@tcp(127.0.0.1:3306)/Inventory?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db

}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"pass" binding:"required"`
}
type Users struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string `gorm:"size:255;index:idx_name,unique"`
	Email       string `gorm:"size:255;"`
	Password    string `gorm:"type:varchar(255)"`
	Phonenumber string `gorm:"size:255;"`
	Role        string `gorm:"size:255;"`
}

type Inventory struct {
	ID              uint64  `gorm:"primaryKey"`
	Name            string  `gorm:"type:varchar(100)"`
	Number          string  `gorm:"size:255;"`
	RolePrice       float64 `gorm:"type:float"`
	MeterPrice      float64 `gorm:"type:float"`
	Count           int8    `gorm:"size:255;"`
	InventoryNumber int32   `gorm:"size:255;"`
}
