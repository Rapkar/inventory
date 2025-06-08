package Boot

import (
	"compress/gzip"
	"errors"
	"fmt"
	"inventory/App/Utility"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
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

	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(6)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db

}

func Init() {

	f, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f)

	// Auto Make Faker Data With Transaction
	experror := DB().Transaction(func(tx *gorm.DB) error {

		if err := tx.AutoMigrate(&Users{}, &Inventory{}, &Product{}, &Export{}, &ExportProducts{}, &Payments{}, &BalanceAdjustment{}); err != nil {
			log.Fatal("❌ Erro in Create Base Tables . #")
			return err
		}

		var existingUser Users
		a, _ := Utility.HashPassword("0000")

		if err := tx.Where("email = ?", "hosseinbidar7@gmail.com").First(&existingUser).Error; err != nil {
			existingUser = Users{
				Name:        "hossein Soltanian",
				Email:       "hosseinbidar7@gmail.com",
				Password:    a,
				Role:        "Admin",
				Phonenumber: "09125174854"}
			if err := tx.Create(&existingUser).Error; err != nil {
				log.Fatal("❌ Erro in insert data to Users Table . #")
				return err
			}
		} else {
			fmt.Println("base user Found")
		}
		var existingInventory Inventory
		if err := tx.Where("id = ?", 1).First(&existingInventory).Error; err != nil {

			existingInventory = Inventory{
				Name: "انبار اشتهارد"}

			if err := tx.Create(&existingInventory).Error; err != nil {
				log.Fatal("❌ Error in batch insert to Inventory Table. #", err)
				return err
			}
			existingInventory = Inventory{
				Name: "انبار زنجان"}

			if err := tx.Create(&existingInventory).Error; err != nil {
				log.Fatal("❌ Error in batch insert to Inventory Table. #", err)
				return err
			}
		} else {
			fmt.Println("base Inventory Found")
		}
		var existinProduct []Product
		if err := tx.Where("id =?", 1).First(&existinProduct).Error; err != nil {
			existinProduct = []Product{{
				Name:              " ایزوگام شرق صادراتی",
				RollePrice:        99250,
				Roll:              100,
				MeasurementSystem: "roll",
				InventoryID:       existingInventory.ID,
			},
				{
					Name:              "ایزوگام غرب شرق مخصوص",
					RollePrice:        92500,
					Roll:              80,
					MeasurementSystem: "roll",
					InventoryID:       existingInventory.ID,
				},
				{
					Name:              "ایزوگام شمال شرق بدون فویل",
					RollePrice:        110000,
					Roll:              115000,
					MeasurementSystem: "roll",
					InventoryID:       existingInventory.ID,
				},
				{
					Name:              "ایزوگام سپید گام صادراتی",
					RollePrice:        95000,
					Roll:              120,
					MeasurementSystem: "roll",
					InventoryID:       existingInventory.ID,
				},

				{
					Name:              "ایزوگام سپیدگام صادراتی بدون فویل",
					RollePrice:        87500,
					Roll:              80,
					MeasurementSystem: "roll",
					InventoryID:       existingInventory.ID,
				},
				{
					Name:              "ایزوگام اصلاحی درجه ۲",
					MeterPrice:        108000,
					Meter:             50,
					MeasurementSystem: "meter",
					InventoryID:       existingInventory.ID,
				},
				{
					Name:              "ایزوگام شرق طرح دار",
					RollePrice:        95000,
					Roll:              120,
					MeasurementSystem: "roll",
					InventoryID:       existingInventory.ID,
				},
				{
					Name:              "بشکه قیر",
					BarrelPrice:       108000,
					Barrel:            75,
					MeasurementSystem: "barrel",
					InventoryID:       existingInventory.ID,
				},
			}
			for _, item := range existinProduct {
				if err := tx.Create(&item).Error; err != nil {
					log.Fatalf("❌ خطا در ثبت محصول %s در جدول exportProducts: %v", item.Name, err)
					return err
				}
			}
			// if err := tx.Create(&existinProduct).Error; err != nil {
			// 	log.Fatal("❌ Erro in insert data to Inventory Table . #")
			// 	return err
			// }
		} else {
			fmt.Println("base Inventory Found")
		}
		var existinExport Export
		if err := tx.Where("id =?", 1).First(&existinExport).Error; err != nil {
			existinExport = Export{
				Name:        "حسین سلطانیان",
				Number:      "9283422",
				Phonenumber: "09125174854",
				UserID:      existingUser.ID,
				CreatorName: existingUser.Name,
				Address:     "کرج -کرج=-ایران -سیسی",
				TotalPrice:  10000000,
				Tax:         10,
				InventoryID: existingInventory.ID,
				CreatedAt:   Utility.CurrentTime(),
			}

			if err := tx.Create(&existinExport).Error; err != nil {
				log.Fatal("❌ Erro in insert data to newExport Table . #")
				return err
			}
		} else {
			fmt.Println("base Inventory Found")
		}

		var existinExportProducts ExportProducts
		if err := tx.Where("id = ?", 1).First(&existinExportProducts).Error; err != nil {
			products := []ExportProducts{
				{
					ExportID:    existinExport.ID,
					Name:        " ایزوگام شرق صادراتی",
					RollePrice:  99250,
					Roll:        5,
					InventoryID: existingInventory.ID,
					ProductID:   1,
					TotalPrice:  496250,
				},
				{
					ExportID:    existinExport.ID,
					Name:        "ایزوگام غرب شرق مخصوص",
					RollePrice:  87500,
					Roll:        5,
					InventoryID: existingInventory.ID,
					ProductID:   2,
					TotalPrice:  437500,
				},
				{
					ExportID:    existinExport.ID,
					Name:        "ایزوگام شمال شرق بدون فویل",
					RollePrice:  110000,
					Roll:        10,
					InventoryID: existingInventory.ID,
					ProductID:   3,
					TotalPrice:  1100000,
				},
				{
					ExportID:    existinExport.ID,
					Name:        "ایزوگام سپید گام صادراتی",
					RollePrice:  95000,
					Roll:        100,
					InventoryID: existingInventory.ID,
					ProductID:   4,
					TotalPrice:  9500000,
				},
				{
					ExportID:    existinExport.ID,
					Name:        "ایزوگام سپیدگام صادراتی بدون فویل",
					RollePrice:  105000,
					Roll:        10,
					InventoryID: existingInventory.ID,
					ProductID:   5,
					TotalPrice:  1050000,
				},
				{
					ExportID:    existinExport.ID,
					Name:        "ایزوگام اصلاحی درجه ۲",
					MeterPrice:  108000,
					Meter:       10.5,
					InventoryID: existingInventory.ID,
					ProductID:   6,
					TotalPrice:  1134000,
				},
				{
					ExportID:    existinExport.ID,
					Name:        "ایزوگام شرق طرح دار",
					RollePrice:  95000,
					Roll:        75,
					InventoryID: existingInventory.ID,
					ProductID:   7,
					TotalPrice:  7125000,
				},
				{
					ExportID:    existinExport.ID,
					Name:        "بشکه قیر",
					BarrelPrice: 108000,
					Barrel:      75,
					InventoryID: existingInventory.ID,
					ProductID:   8,
					TotalPrice:  8100000,
				},
			}

			// ذخیره همه محصولات در دیتابیس با یک حلقه
			for _, product := range products {
				if err := tx.Create(&product).Error; err != nil {
					log.Fatalf("❌ خطا در ثبت محصول %s در جدول exportProducts: %v", product.Name, err)
					return err
				}
			}
		} else {
			fmt.Println("base Inventory Found")
		}

		var existinPayments Payments
		if err := tx.Where("id = ?", 1).First(&existinExport).Error; err != nil {
			existinPayments = Payments{
				Method:      "مستقیم",
				Number:      "9283422",
				TotalPrice:  9000,
				Name:        "ملی",
				Describe:    "کرج -کرج=-ایران -سیسی",
				CreatedAt:   Utility.CurrentTime(),
				ExportID:    existinExport.ID,
				UserID:      1,
				InventoryID: existingInventory.ID,
				Status:      "collected",
			}

			if err := tx.Create(&existinPayments).Error; err != nil {
				log.Fatal("❌ Erro in insert data to existinPayments Table . #")

				return err
			}
		} else {
			fmt.Println("base Inventory Found")
		}

		var balanceAdjustment BalanceAdjustment
		if err := tx.Where("id =?", 1).First(&balanceAdjustment).Error; err != nil {
			balanceAdjustment = BalanceAdjustment{
				UserID:       existingUser.ID,
				OffsetAmount: 0,
				Reason:       "قیمت اصلاحی",
				CreatedBy:    existingUser.ID,
				CreatedAt:    Utility.CurrentTime(),
			}

			if err := tx.Create(&balanceAdjustment).Error; err != nil {
				log.Fatal("❌ Erro in insert data to balanceAdjustment Table . #")
				return err
			}
		} else {
			fmt.Println("base Inventory Found")
		}

		return err

	})

	if experror != nil {
		fmt.Println("❌ خطا در ساخت دیتابیس ها:", err)
		return
	}

	fmt.Println("✅ Create  base tables sucess . #")

}

// func TakeBackup(fs afero.Fs, is int) {
// 	username := "root"
// 	dbName := "Inventory"
// 	password := "0311121314" // replace with your actual password
// 	viper.SetConfigFile(".env")
// 	viper.SetConfigName("config")
// 	viper.AddConfigPath(".")
// 	viper.ReadInConfig()
// 	i := 1
// 	t := time.Now().Add(-time.Hour * 24 * time.Duration(i))
// 	backupName := fmt.Sprintf("backup-%s.sql", t.Format("2006-01-02-3-4-5"))
// 	if is == 1 && viper.GetString("LAST_BS") == "" {
// 		viper.Set("LAST_BS", backupName)
// 	}
// 	cmd := exec.Command("mysqldump", "-u", username, dbName)
// 	cmd.Env = append(os.Environ(), fmt.Sprintf("MYSQL_PWD=%s", password))
// 	file, err := os.Create(backupName)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
// 	cmd.Stdout = file
// 	cmd.Stderr = os.Stderr
// 	err = cmd.Run()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Backup saved to backup.sql", backupName, "round ", is)
// 	if is == 1 {

// 		err = fs.Remove(RemoveFileName(t, backupName, 2))
// 		// fmt.Println(fs.Stat(backupName))
// 		if err != nil {
// 			// log.Printf(err)
// 		}
// 		// is = 0
// 		// viper.Set("LAST_BS", "")
// 	} else if is == 2 {

// 		err = fs.Remove(RemoveFileName(t, backupName, 2))
// 		if err != nil {
// 			// log.Fatal(err)
// 		}

// 	} else if is == 3 {

// 		err = fs.Remove(RemoveFileName(t, backupName, 2))
// 		if err != nil {
// 			// log.Fatal(err)
// 		}
// 		// viper.Set("LAST_BS", "")
// 	} else if is == 4 {
// 		err = fs.Remove(RemoveFileName(t, backupName, 2))
// 		if err != nil {
// 			// log.Fatal(err)
// 		}
// 		is = 0
// 		// viper.Set("LAST_BS", "")
// 	}

// }

const backupDir = "./backups"
const maxBackups = 5 // مثلاً ۵ عدد آخر رو نگه داریم

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
	fs.MkdirAll(backupDir, 0755)

	backupTime := time.Now()
	filename := fmt.Sprintf("backup-%d-%s.sql", backupType, backupTime.Format("2006-01-02-15-04-05"))
	backupPath := filepath.Join(backupDir, filename)

	// اجرای mysqldump
	cmd := exec.Command("mysqldump", "-u", dbUser, dbName)
	cmd.Env = append(os.Environ(), fmt.Sprintf("MYSQL_PWD=%s", dbPass))

	file, err := fs.Create(backupPath)
	if err != nil {
		return fmt.Errorf("cannot create backup file: %v", err)
	}
	defer file.Close()

	cmd.Stdout = file
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("mysqldump failed: %v", err)
	}

	// فشرده‌سازی فایل
	if err := compressFile(fs, backupPath); err != nil {
		return fmt.Errorf("gzip compress failed: %v", err)
	}

	// حذف فایل .sql پس از فشرده‌سازی
	fs.Remove(backupPath)

	// پاک‌سازی بکاپ‌های قدیمی
	if err := cleanupOldBackups(fs, backupType); err != nil {
		log.Println("⚠️ error in cleanup:", err)
	}

	log.Println("✅ backup created and compressed:", filename+".gz")
	return nil
}

func compressFile(fs afero.Fs, src string) error {
	srcFile, err := fs.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := fs.Create(src + ".gz")
	if err != nil {
		return err
	}
	defer dstFile.Close()

	writer := gzip.NewWriter(dstFile)
	defer writer.Close()

	_, err = io.Copy(writer, srcFile)
	return err
}

func cleanupOldBackups(fs afero.Fs, backupType int) error {
	files, err := afero.ReadDir(fs, backupDir)
	if err != nil {
		return err
	}

	var matched []os.FileInfo
	prefix := fmt.Sprintf("backup-%d-", backupType)

	for _, f := range files {
		if strings.HasPrefix(f.Name(), prefix) && strings.HasSuffix(f.Name(), ".sql.gz") {
			matched = append(matched, f)
		}
	}

	// مرتب‌سازی قدیمی‌ترین به جدیدترین
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].ModTime().Before(matched[j].ModTime())
	})

	// حذف اضافی‌ها
	if len(matched) > maxBackups {
		for _, old := range matched[:len(matched)-maxBackups] {
			fs.Remove(filepath.Join(backupDir, old.Name()))
		}
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
func GetBackupList2(fs afero.Fs, baseURL string) ([]BackupFile, error) {
	var backups []BackupFile

	// فقط پوشه backups رو بخون
	err := afero.Walk(fs, "backups", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasPrefix(info.Name(), "backup-") && strings.HasSuffix(info.Name(), ".sql.gz") {
			// timeStr := strings.TrimPrefix(info.Name(), "backup-")
			// timeStr = strings.TrimSuffix(timeStr, ".sql.gz")

			parts := strings.SplitN(info.Name(), "-", 3)
			if len(parts) < 3 {
				log.Printf("نام فایل بک‌آپ معتبر نیست: %s", info.Name())
				return nil
			}

			timeStr := strings.TrimSuffix(parts[2], ".sql.gz")
			backupTime, err := time.Parse("2006-01-02-15-04-05", timeStr)
			if err != nil {
				log.Printf("خطا در تجزیه تاریخ فایل بک‌آپ %s: %v", info.Name(), err)
				return nil
			}

			// ساخت لینک دانلود
			downloadURL := fmt.Sprintf("%s/backups/%s", strings.TrimRight(baseURL, "/"), info.Name())

			backups = append(backups, BackupFile{
				Name:        info.Name(),
				Size:        info.Size(),
				ModTime:     info.ModTime(),
				BackupTime:  backupTime,
				DownloadURL: downloadURL,
			})
		}
		fmt.Println("sssss2")
		return nil
	})

	if err != nil {
		return nil, err
	}

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

// func PeroudBackup() {
// 	is := 1
// 	// oldbackupName := ""
// 	ticker := time.NewTicker(24 * time.Hour)
// 	fs := afero.NewOsFs()
// 	for range ticker.C {
// 		fmt.Println("\n in parent round number is =", is)

// 		if is == 1 {
// 			TakeBackup2(fs, 1)
// 		} else if is == 2 {
// 			TakeBackup2(fs, 2)
// 		} else if is == 3 {
// 			TakeBackup2(fs, 3)
// 		} else if is == 4 {
// 			TakeBackup2(fs, 4)
// 			is = 0
// 		}
// 		is++
// 		// fmt.Println("number=", is, " name=", oldbackupName)

// 	}
// }
