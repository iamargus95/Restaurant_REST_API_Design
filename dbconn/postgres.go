package dbconn

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/iamargus95/restaurant_rest_api_design/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//*DB is connected db object
var DB *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("HOST"), os.Getenv("DBPORT"), os.Getenv("DBUSER"),
		os.Getenv("DBNAME"), os.Getenv("PASSWORD"))

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to DB, ERROR: %v", err)
		os.Exit(100)
	}

	// Shows RAW SQL Queries in stdout
	conn.Logger.LogMode(4)

	DB = conn
	sqlDB, _ := DB.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(10 * time.Minute)

	// Creates tables via tableName method
	DB.Debug().AutoMigrate(&models.Users{}, &models.Restaurant{}, &models.Menu_Item{})
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}
