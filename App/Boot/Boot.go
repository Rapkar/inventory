package Boot

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

type Export struct {
	ID              uint64           `gorm:"primaryKey"`
	Name            string           `gorm:"type:varchar(100)"`
	Number          string           `gorm:"size:255;"`
	Phonenumber     string           `gorm:"size:255;"`
	Address         string           `gorm:"size:255;"`
	TotalPrice      float64          `gorm:"type:decimal(10,2);"`
	Tax             float64          `gorm:"type:float"`
	CreatedAt       string           `json:"start_date"` // assign the format to a string
	InventoryNumber int32            `gorm:"size:255;"`
	ExportProducts  []ExportProducts `gorm:"foreignKey:ExportID"`
}
type EscapeExport struct {
	ID              uint64           `gorm:"primaryKey"`
	Name            string           `gorm:"type:varchar(100)"`
	Number          string           `gorm:"size:255;"`
	Phonenumber     string           `gorm:"size:255;"`
	Address         string           `gorm:"size:255;"`
	TotalPrice      string           `gorm:"type:string"`
	Tax             float64          `gorm:"type:float"`
	CreatedAt       string           `json:"start_date"` // assign the format to a string
	InventoryNumber int32            `gorm:"size:255;"`
	ExportProducts  []ExportProducts `gorm:"foreignKey:ExportID"`
}
type ExportProducts struct {
	ID              uint64  `gorm:"primaryKey"`
	ExportID        uint64  `gorm:"size:255;"`
	Name            string  `gorm:"type:varchar(100)"`
	Number          string  `gorm:"size:255;"`
	RolePrice       float64 `gorm:"type:float"`
	MeterPrice      float64 `gorm:"type:float"`
	Count           int8    `gorm:"size:255;"`
	Meter           int8    `gorm:"size:255;"`
	TotalPrice      float64 `gorm:"size:255;"`
	InventoryNumber int32   `gorm:"size:255;"`
}
