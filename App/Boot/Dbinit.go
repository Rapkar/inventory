package Boot

import (
	"fmt"
	"inventory/App/Utility"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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
	fmt.Println("linl", dsn)
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
			TakeBackup(fs, 1)
		} else if is == 2 {
			TakeBackup(fs, 2)
		} else if is == 3 {
			TakeBackup(fs, 3)
		} else if is == 4 {
			TakeBackup(fs, 4)
			is = 0
		}
		is++
		// fmt.Println("number=", is, " name=", oldbackupName)

	}
}
