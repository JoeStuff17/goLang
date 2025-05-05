package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"main.go/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

// var DBSql DbInstance
var DBSql *gorm.DB

func ConnectToMySql() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	// SQL Connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", user, password, host, port, name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to SQL database! \n", err.Error())
	} else {
		log.Println("Connected to Mysql database successfully!")
	}
	db.Logger = logger.Default.LogMode(logger.Info)

	if os.Getenv("DB_MIGRATE") == "true" {
		fmt.Println("Running Migrations")
		err := db.AutoMigrate(
			&models.Users{},
			&models.Admins{},
			&models.Churches{},
			&models.ChurchUser{},
			&models.ChurchFamily{},
		)
		if err != nil {
			fmt.Println("DB Migrations error, aborting...")
			fmt.Println(err.Error())
			panic("Without Database, I am nothing. Bye. 🥺")
		}
	} else {
		fmt.Println("Migrations disabled, skipping...")
	}

	DBSql = db
}
