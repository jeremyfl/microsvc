package gorm

import (
	"fmt"
	"gitlab.com/jeremylo/microsvc/lib"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewClient() (*gorm.DB, error) {
	host := lib.Config["DB_HOST"]
	if host == "" {
		host = "127.0.0.1"
	}

	dsn := fmt.Sprintf("root:12345678@tcp(%s:3306)/ordersvc?charset=utf8mb4&parseTime=True&loc=Local", host)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
