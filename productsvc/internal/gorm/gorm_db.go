package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	dsn := "root:12345678@tcp(127.0.0.1:3306)/productsvc?charset=utf8mb4&parseTime=True&loc=Local"

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
