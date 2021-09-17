package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	// DB, err := gorm.Open("mysql", "${DB_USER}:${DB_PASS}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?charset=utf8&parseTime=True&loc=Local")
	DB, err := gorm.Open("mysql", "root:QWErty@123@tcp(localhost:3306)/todo_schema?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	defer DB.Close()

	fmt.Println("database connected")

	return DB
}
