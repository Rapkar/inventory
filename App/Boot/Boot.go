package Boot

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"pass" binding:"required"`
}
type Users struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string `gorm:"size:255;"`
	Email       string `gorm:"size:255;index:unique"`
	Password    string `gorm:"type:varchar(255)"`
	Phonenumber string `gorm:"size:255;unique"`
	Address     string `gorm:"size:255;"`
	Role        string `gorm:"size:255;"`
}
type ResponseUsers struct {
	ID          uint64 `gorm:"primaryKey"`
	Name        string `gorm:"size:255;"`
	Email       string `gorm:"size:255;index:unique"`
	Phonenumber string `gorm:"size:255;unique"`
	Address     string `gorm:"size:255;"`
}

//	type Inventory struct {
//		ID              uint64 `gorm:"primaryKey"`
//		Name            string `gorm:"type:varchar(100)"`
//		Number          string `gorm:"size:255;"`
//		RolePrice       int64  `gorm:"type:float"`
//		MeterPrice      int64  `gorm:"type:float"`
//		Count           int64  `gorm:"size:255;"`
//		Meter           int64  `gorm:"size:255;"`
//		InventoryNumber int32  `gorm:"size:255;"`
//	}
type Inventory struct {
	ID   uint64 `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(100)"`
}
type Product struct {
	ID                uint64    `gorm:"primaryKey"`
	Name              string    `gorm:"type:varchar(100)"`
	RollePrice        float64   `gorm:"type:float"`
	MeterPrice        float64   `gorm:"type:float"`
	WeightPrice       float64   `gorm:"type:float"`
	CountPrice        float64   `gorm:"type:float"`
	BarrelPrice       float64   `gorm:"type:float"`
	Roll              int64     `gorm:"size:255;"`
	Meter             float64   `gorm:"type:float"`
	Weight            float64   `gorm:"type:float"`
	Count             int64     `gorm:"size:255;"`
	Barrel            int64     `gorm:"size:255;"`
	InventoryID       uint64    `gorm:"index"`
	MeasurementSystem string    `gorm:"type:varchar(100)"`
	Inventory         Inventory `gorm:"foreignKey:InventoryID;references:ID"`
}
type Export struct {
	ID             uint64           `gorm:"primaryKey"`
	Name           string           `gorm:"type:varchar(100)"`
	Number         string           `gorm:"varchar(255),unique"`
	Phonenumber    string           `gorm:"size:255;"`
	UserID         uint64           `gorm:"index;"`
	CreatorName    string           `gorm:"size:255;"`
	Address        string           `gorm:"size:255;"`
	TotalPrice     float64          `gorm:"type:float"`
	Tax            int64            `gorm:"size:255;"`
	Describe       string           `gorm:"size:255;"`
	CreatedAt      string           `json:"CreatedAt"` // assign the format to a string
	Draft          bool             `gorm:"type:boolean"`
	InventoryID    uint64           `gorm:"index;"`
	ExportProducts []ExportProducts `gorm:"foreignKey:ExportID;references:ID"`
	Payments       []Payments       `gorm:"foreignKey:ExportID"`
	User           Users            `gorm:"foreignKey:UserID;references:ID"`
	// Inventory      Inventory        `gorm:"foreignKey:InventoryNumber"`
	Inventory Inventory `gorm:"foreignKey:InventoryID;references:ID"`
}
type Payments struct {
	ID          uint64    `gorm:"primaryKey"`
	Method      string    `gorm:"type:varchar(100)"`
	Number      string    `gorm:"varchar(255),unique"`
	Name        string    `gorm:"type:varchar(100)"`
	TotalPrice  float64   `gorm:"type:float"`
	Describe    string    `gorm:"size:255;"`
	CreatedAt   string    `json:"CreatedAt"` // assign the format to a string
	ExportID    uint64    `gorm:"index"`
	UserID      uint64    `gorm:"index"`
	InventoryID uint64    `gorm:"index"`
	Export      Export    `gorm:"foreignKey:ExportID"`
	Status      string    `gorm:"type:varchar(100)"`
	User        Users     `gorm:"foreignKey:UserID"`
	Inventory   Inventory `gorm:"foreignKey:InventoryID;references:ID"`
}

type ExportProducts struct {
	ID                uint64    `gorm:"primaryKey"`
	ExportID          uint64    `gorm:"index"`
	Name              string    `gorm:"type:varchar(100)"`
	RollePrice        float64   `gorm:"type:float"`
	MeterPrice        float64   `gorm:"type:float"`
	WeightPrice       float64   `gorm:"type:float"`
	CountPrice        float64   `gorm:"type:float"`
	BarrelPrice       float64   `gorm:"type:float"`
	Roll              int64     `gorm:"size:255;"`
	Meter             float64   `gorm:"type:float"`
	Weight            float64   `gorm:"type:float"`
	Count             int64     `gorm:"size:255;"`
	Barrel            int64     `gorm:"size:255;"`
	TotalPrice        float64   `gorm:"type:float"`
	InventoryID       uint64    `gorm:"index"`
	ProductID         uint64    `gorm:"index"`
	MeasurementSystem string    `gorm:"type:varchar(100)"`
	Export            Export    `gorm:"foreignKey:ExportID;references:ID"`
	Inventory         Inventory `gorm:"foreignKey:InventoryID;references:ID"`
	Product           Product   `gorm:"foreignKey:ProductID;references:ID"`
}

type EscapeExport struct {
	ID              uint64           `gorm:"primaryKey"`
	Name            string           `gorm:"type:varchar(100)"`
	Number          string           `gorm:"size:255;"`
	Phonenumber     string           `gorm:"size:255;"`
	Address         string           `gorm:"size:255;"`
	TotalPrice      float64          `gorm:"type:float"`
	Tax             int64            `gorm:"size:255;"`
	Draft           bool             `gorm:"type:boolean"`
	Describe        string           `gorm:"size:255;"`
	CreatedAt       string           `json:"CreatedAt"` // assign the format to a string
	CreatorName     string           `gorm:"size:255;"`
	InventoryNumber int32            `gorm:"size:255;"`
	ExportProducts  []ExportProducts `gorm:"foreignKey:ExportID"`
	InventoryName   string           `gorm:"size:255;"`
}
type EscapeExportProducts struct {
	ID          uint64  `gorm:"primaryKey"`
	ExportID    uint64  `gorm:"index"`
	Name        string  `gorm:"type:varchar(100)"`
	RollePrice  float64 `gorm:"type:float"`
	MeterPrice  float64 `gorm:"type:float"`
	WeightPrice float64 `gorm:"type:float"`
	CountPrice  float64 `gorm:"type:float"`
	BarrelPrice float64 `gorm:"type:float"`
	Roll        int64   `gorm:"size:255;"`
	Meter       float64 `gorm:"type:float"`
	Weight      float64 `gorm:"type:float"`
	Count       int64   `gorm:"size:255;"`
	Barrel      int64   `gorm:"size:255;"`
	TotalPrice  float64 `gorm:"type:float"`
	ProductID   uint64  `gorm:"index"`
	Product     Product `gorm:"foreignKey:ProductID;references:ID"`

	InventoryNumber int32 `gorm:"size:255;"`
	// Inventory       Inventory `gorm:"foreignKey:InventoryNumber"`
}
type PaymentWithExportAndUser struct {
	Payments
	ExportNumber string `json:"export_number"`
	UserName     string `json:"UserName"`
	PhoneNumber  string `json:"PhoneNumber"`
}

type ProductWithInventory struct {
	Product
	Inventory
}
type BalanceAdjustment struct {
	ID            uint64  `gorm:"primaryKey"`
	UserID        uint64  `gorm:"index"`
	OffsetAmount  float64 // مثبت یا منفی
	Reason        string
	CreatedBy     uint64 `gorm:"index"`
	CreatedAt     string
	User          Users `gorm:"foreignKey:UserID"` // کاربر مربوطه
	CreatedByUser Users `gorm:"foreignKey:CreatedBy"`
}
