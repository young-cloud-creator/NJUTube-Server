package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func InitDB() error {
	// connect to the database
	// dsn -> [user]:[password]@[protocol]([address])/[database]
	const dsn = "root:12345678@tcp(127.0.0.1:3306)/tube?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	database = db

	return err
}
