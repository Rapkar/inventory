package Boot

import (
	"errors"
	"fmt"
	"inventory/App/Utility"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/afero"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() *gorm.DB {

	viper.SetConfigFile(".env")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	MODE := viper.Get("MODE")

	var (
		DATABASENAME     string
		DATABASEUSER     string
		DATABASEUSERPASS string
		DATABASEPORT     string
		DATABASEHOST     string
	)
	if MODE == "DEVELOP" {
		DATABASENAME = viper.GetString("MYSQL_LOCAL_DATABASE")
		DATABASEUSER = viper.GetString("MYSQL_LOCAL_USERNAME")
		DATABASEUSERPASS = viper.GetString("MYSQL_LOCAL_PASS")
		DATABASEPORT = viper.GetString("MYSQL_LOCAL_PORT")
		DATABASEHOST = viper.GetString("MYSQL_LOCAL_HOST")
	} else {
		DATABASENAME = viper.GetString("MYSQL_DATABASE")
		DATABASEUSER = viper.GetString("MYSQL_USERNAME")
		DATABASEUSERPASS = viper.GetString("MYSQL_PASS")
		DATABASEPORT = viper.GetString("MYSQL_PORT")
		DATABASEHOST = viper.GetString("MYSQL_HOST")
	}

	dsn := DATABASEUSER + ":" + DATABASEUSERPASS + "@tcp(" + DATABASEHOST + ":" + DATABASEPORT + ")/" + DATABASENAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, err := db.DB()

	if err != nil {
		log.Printf("Inventory Has problem With connect to Database")
		// sqlDB.Close()

	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(6)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)
	// go func() {
	// 	ticker := time.NewTicker(5 * time.Minute)
	// 	defer ticker.Stop()

	// 	for range ticker.C {
	// 		sqlDB, err := DB().DB()
	// 		if err != nil {
	// 			log.Printf("Pool stats error: %v", err)
	// 			continue
	// 		}

	// 		stats := sqlDB.Stats()
	// 		log.Printf("DB Pool Stats: OpenConnections=%d InUse=%d Idle=%d",
	// 			stats.OpenConnections, stats.InUse, stats.Idle)
	// 	}
	// }()

	return db

}

func Init() {

	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f)
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

	if !DB().Migrator().HasTable(Payments{}) {
		DB().Migrator().CreateTable(Payments{})

		// First create sample ExportProducts for the Export
		exportProducts := []ExportProducts{
			{
				Name:            "ایزوگام شرق",
				Number:          "10",
				RolePrice:       99250,
				MeterPrice:      102500,
				Count:           100,
				Meter:           10,
				TotalPrice:      2000000,
				InventoryNumber: 1,
			},
		}

		// Create an Export to associate with
		export := Export{
			Name:            "رضا توانگر",
			Number:          "9283422",
			Phonenumber:     "09199656725",
			Address:         "کرج -کرج=-ایران -سیسی",
			TotalPrice:      10000000,
			Tax:             10,
			ExportProducts:  exportProducts,
			InventoryNumber: 1,
			CreatedAt:       Utility.CurrentTime(),
		}
		DB().Create(&export)

		// Now create the Payment associated with this Export
		Payment := Payments{
			Method:     "مستقیم",
			Number:     "9283422",
			TotalPrice: 9000,
			Name:       "ملی",
			Describe:   "کرج -کرج=-ایران -سیسی",
			CreatedAt:  Utility.CurrentTime(),
			ExportID:   export.ID, // Associate with the export
		}
		DB().Create(&Payment)
	} else {
		fmt.Println("Payments table found. #")
	}

}

func TakeBackup(fs afero.Fs, is int) {
	username := "root"
	dbName := "Inventory"
	password := "0311121314" // replace with your actual password
	viper.SetConfigFile(".env")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	i := 1
	t := time.Now().Add(-time.Hour * 24 * time.Duration(i))
	backupName := fmt.Sprintf("backup-%s.sql", t.Format("2006-01-02-3-4-5"))
	if is == 1 && viper.GetString("LAST_BS") == "" {
		viper.Set("LAST_BS", backupName)
	}
	cmd := exec.Command("mysqldump", "-u", username, dbName)
	cmd.Env = append(os.Environ(), fmt.Sprintf("MYSQL_PWD=%s", password))
	file, err := os.Create(backupName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	cmd.Stdout = file
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Backup saved to backup.sql", backupName, "round ", is)
	if is == 1 {

		err = fs.Remove(RemoveFileName(t, backupName, 2))
		// fmt.Println(fs.Stat(backupName))
		if err != nil {
			// log.Printf(err)
		}
		// is = 0
		// viper.Set("LAST_BS", "")
	} else if is == 2 {

		err = fs.Remove(RemoveFileName(t, backupName, 2))
		if err != nil {
			// log.Fatal(err)
		}

	} else if is == 3 {

		err = fs.Remove(RemoveFileName(t, backupName, 2))
		if err != nil {
			// log.Fatal(err)
		}
		// viper.Set("LAST_BS", "")
	} else if is == 4 {
		err = fs.Remove(RemoveFileName(t, backupName, 2))
		if err != nil {
			// log.Fatal(err)
		}
		is = 0
		// viper.Set("LAST_BS", "")
	}

}
func TakeBackup2(fs afero.Fs, backupType int) error {
	// خواندن تنظیمات از محیط
	viper.SetConfigFile(".env")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	MODE := viper.Get("MODE")

	var (
		dbName string
		dbUser string
		dbPass string
	)
	if MODE == "DEVELOP" {
		dbName = viper.GetString("MYSQL_LOCAL_DATABASE")
		dbUser = viper.GetString("MYSQL_LOCAL_USERNAME")
		dbPass = viper.GetString("MYSQL_LOCAL_PASS")

	} else {
		dbName = viper.GetString("MYSQL_DATABASE")
		dbUser = viper.GetString("MYSQL_USERNAME")
		dbPass = viper.GetString("MYSQL_PASS")

	}

	if dbUser == "" || dbName == "" || dbPass == "" {
		return errors.New("تنظیمات دیتابیس ناقص است")
	}

	// ایجاد نام فایل بکاپ
	backupTime := time.Now().Add(-time.Hour * 24 * time.Duration(backupType))
	backupName := fmt.Sprintf("backup-%s.sql", backupTime.Format("2006-01-02-15-04-05"))

	// اجرای دستور mysqldump
	cmd := exec.Command("mysqldump", "-u", dbUser, dbName)
	cmd.Env = append(os.Environ(), fmt.Sprintf("MYSQL_PWD=%s", dbPass))

	file, err := fs.Create(backupName)
	if err != nil {
		return fmt.Errorf("error in make backup file: %v", err)
	}
	defer file.Close()

	cmd.Stdout = file
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("eror in  mysqldump: %v", err)
	}

	// مدیریت فایل‌های قدیمی
	if err := manageOldBackups(fs, backupType, backupTime); err != nil {
		log.Printf("error in manage old backup files: %v", err)
	}

	log.Printf("✅ created backup success %s", backupName)
	return nil
}

func manageOldBackups(fs afero.Fs, backupType int, backupTime time.Time) error {
	// تعریف سیاست‌های نگهداری بکاپ‌ها
	retentionDays := map[int]int{
		1: 7,   // بکاپ روزانه - نگهداری 7 روز
		2: 30,  // بکاپ هفتگی - نگهداری 30 روز
		3: 365, // بکاپ ماهانه - نگهداری 1 سال
	}

	days, exists := retentionDays[backupType]
	if !exists {
		return nil
	}

	oldBackupTime := backupTime.Add(-time.Hour * 24 * time.Duration(days))
	oldBackupName := fmt.Sprintf("backup-%s.sql", oldBackupTime.Format("2006-01-02-15-04-05"))

	if err := fs.Remove(oldBackupName); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error in remove old backup files: %v", err)
	}

	return nil
}
func ScheduleBackups() {
	// بک‌آپ روزانه
	go func() {
		for {
			TakeBackup2(afero.NewOsFs(), 1)
			time.Sleep(24 * time.Hour)
		}
	}()

	// بک‌آپ هفتگی
	go func() {
		for {
			TakeBackup2(afero.NewOsFs(), 2)
			time.Sleep(7 * 24 * time.Hour)
		}
	}()
}
func GetBackupList(fs afero.Fs, baseURL string) ([]BackupFile, error) {
	var backups []BackupFile

	afero.Walk(fs, ".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasPrefix(info.Name(), "backup-") && strings.HasSuffix(info.Name(), ".sql") {
			timeStr := strings.TrimPrefix(info.Name(), "backup-")
			timeStr = strings.TrimSuffix(timeStr, ".sql")

			backupTime, err := time.Parse("2006-01-02-15-04-05", timeStr)
			if err != nil {
				log.Printf("خطا در تجزیه تاریخ فایل بک‌آپ %s: %v", info.Name(), err)
				return nil
			}

			// ایجاد لینک دانلود
			downloadURL := fmt.Sprintf("%s/%s", strings.TrimRight(baseURL, "/"), info.Name())

			backups = append(backups, BackupFile{
				Name:        info.Name(),
				Size:        info.Size(),
				ModTime:     info.ModTime(),
				BackupTime:  backupTime,
				DownloadURL: downloadURL,
			})
		}
		return nil
	})

	// بقیه کد مانند قبل...
	return backups, nil
}

type BackupFile struct {
	Name        string    `json:"name"`        // نام فایل
	Size        int64     `json:"size"`        // حجم فایل (بایت)
	ModTime     time.Time `json:"modTime"`     // زمان آخرین تغییر
	BackupTime  time.Time `json:"backupTime"`  // زمان بک‌آپ
	DownloadURL string    `json:"downloadUrl"` // لینک دانلود
}

// return old file name
func RemoveFileName(t time.Time, backupName string, ordertiem int) string {
	oldBackupName := ""
	oldBackupName = filepath.Join(filepath.Dir(backupName), "backup-"+t.Add(-time.Hour*48).Format("2006-01-02-3-4-5")+".sql")
	fmt.Println(" so detectedfile is :", oldBackupName, "and removed")
	return oldBackupName
}

func PeroudBackup() {
	is := 1
	// oldbackupName := ""
	ticker := time.NewTicker(24 * time.Hour)
	fs := afero.NewOsFs()
	for range ticker.C {
		fmt.Println("\n in parent round number is =", is)

		if is == 1 {
			TakeBackup2(fs, 1)
		} else if is == 2 {
			TakeBackup2(fs, 2)
		} else if is == 3 {
			TakeBackup2(fs, 3)
		} else if is == 4 {
			TakeBackup2(fs, 4)
			is = 0
		}
		is++
		// fmt.Println("number=", is, " name=", oldbackupName)

	}
}
