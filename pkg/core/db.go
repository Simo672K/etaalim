package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var DB *gorm.DB
var once sync.Once

func GetDBInstance() *gorm.DB {
	once.Do(func() {
		dsn := "root:123456@tcp(localhost:3306)/etaalim?charset=utf8&parseTime=True"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(fmt.Errorf("[ERROR] Failed to connect to database ->  %w", err))
		}

		fmt.Println("[INFO] Successfully connected to database")
		DB = db
	})
	return DB
}
