package initializers

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() (*gorm.DB) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%s TimeZone=Africa/Nairobi", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}

	sqlDB, err := db.DB()

	if err != nil {
		// control error
		fmt.Printf("Error connecting to database : error=%v", err)
	}

	// sqlDB.Ping()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(2)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(1)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db
}