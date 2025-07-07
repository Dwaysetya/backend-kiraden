package database

import (
	"backend/config"
	"backend/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var DB*gorm.DB

func InitDB(){
	dbUser := config.GetEnv("DB_USER", "postgres")
	dbPass := config.GetEnv("DB_PASS", "1")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "5434")
	dbName := config.GetEnv("DB_NAME", "mydb")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbName, dbPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	fmt.Println("Database connected succesfully")

	err = DB.AutoMigrate(&models.User{}, (&models.Products{}))
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}


	fmt.Println("Database migrated successfully!")


}