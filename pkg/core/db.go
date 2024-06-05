package core

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func GetDBInstance() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		panic("Failed to load environment variables, make sure you have .env file!")
	}
	
	once.Do(func() {
		dsn := os.Getenv("MYSQL_DB_URL")
		fmt.Println("mysql url: ", dsn)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("[ERROR] Failed to connect to database ->  %w", err))
		}

		fmt.Println("[INFO] Successfully connected to database")
		DB = db
	})
	return DB
}
