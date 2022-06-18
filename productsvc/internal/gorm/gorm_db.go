package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewClient() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = "127.0.0.1"
	}

	dsn := fmt.Sprintf("root:12345678@tcp(%s:3306)/productsvc?charset=utf8mb4&parseTime=True&loc=Local", host)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
