package DB

import (
	"log"

	"DB_GORM/models"
	"DB_GORM/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	LoadEnv()

	dsn := GetDSN()

	utils.InitLogger()
	utils.InfoLog.Println("Connecting to DB with DSN:", dsn)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(&models.User{}, &models.Job{}, &models.Recruiter{}, &models.Application{}, &models.Interview{})
	if err != nil {
		utils.ErrorLog.Printf("AutoMigrate failed: %v", err)
	}

	utils.InfoLog.Println("Database connected successfully")
}

