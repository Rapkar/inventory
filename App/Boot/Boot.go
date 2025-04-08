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
	Phonenumber string `gorm:"size:255,unique"`
	Address     string `gorm:"size:255;"`
	Role        string `gorm:"size:255;"`
}

type Inventory struct {
	ID              uint64 `gorm:"primaryKey"`
	Name            string `gorm:"type:varchar(100)"`
	Number          string `gorm:"size:255;"`
	RolePrice       int64  `gorm:"type:float"`
	MeterPrice      int64  `gorm:"type:float"`
	Count           int64  `gorm:"size:255;"`
	Meter           int64  `gorm:"size:255;"`
	InventoryNumber int32  `gorm:"size:255;"`
}

type Export struct {
	ID              uint64           `gorm:"primaryKey"`
	Name            string           `gorm:"type:varchar(100)"`
	Number          string           `gorm:"varchar(255),unique"`
	Phonenumber     string           `gorm:"size:255;"`
	Address         string           `gorm:"size:255;"`
	TotalPrice      int64            `gorm:"size:255;"`
	Tax             int64            `gorm:"size:255;"`
	Describe        string           `gorm:"size:255;"`
	CreatedAt       string           `json:"CreatedAt"` // assign the format to a string
	InventoryNumber int32            `gorm:"size:255;"`
	ExportProducts  []ExportProducts `gorm:"foreignKey:ExportID"`
}
type ExportProducts struct {
	ID              uint64 `gorm:"primaryKey"`
	ExportID        uint64 `gorm:"index"`
	Name            string `gorm:"type:varchar(100)"`
	Number          string `gorm:"size:255;"`
	RolePrice       int64  `gorm:"size:255;"`
	MeterPrice      int64  `gorm:"size:255;"`
	Count           int64  `gorm:"size:255;"`
	Meter           int64  `gorm:"size:255;"`
	TotalPrice      int64  `gorm:"size:255;"`
	InventoryNumber int32  `gorm:"size:255;"`
}

type EscapeExport struct {
	ID              uint64           `gorm:"primaryKey"`
	Name            string           `gorm:"type:varchar(100)"`
	Number          string           `gorm:"size:255;"`
	Phonenumber     string           `gorm:"size:255;"`
	Address         string           `gorm:"size:255;"`
	TotalPrice      string           `gorm:"type:string"`
	Tax             int64            `gorm:"size:255;"`
	Describe        string           `gorm:"size:255;"`
	CreatedAt       string           `json:"CreatedAt"` // assign the format to a string
	InventoryNumber int32            `gorm:"size:255;"`
	ExportProducts  []ExportProducts `gorm:"foreignKey:ExportID"`
}
type EscapeExportProducts struct {
	ID              uint64 `gorm:"primaryKey"`
	ExportID        uint64 `gorm:"index"`
	Name            string `gorm:"type:varchar(100)"`
	Number          string `gorm:"size:255;"`
	RolePrice       string `gorm:"size:255;"`
	MeterPrice      string `gorm:"size:255;"`
	Count           string `gorm:"size:255;"`
	Meter           string `gorm:"size:255;"`
	TotalPrice      string `gorm:"size:255;"`
	InventoryNumber int32  `gorm:"size:255;"`
}
