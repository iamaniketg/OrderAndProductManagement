package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	dsn := "root:aniket@12345G@tcp(127.0.0.1:3306)/goDatabase?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate database models
	AutoMigrate()
}

// AutoMigrate performs auto-migration of database models
func AutoMigrate() {
	db.AutoMigrate(&User{}, &Product{}, &Order{}, &Customer{}, &Inventory{})
}
