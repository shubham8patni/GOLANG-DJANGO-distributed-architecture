package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB1 *gorm.DB

func Initialize() error {
	//dsn := "root:mysql@tcp(127.0.0.1:3309)/public_catalog?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open("mysql", "root:mysql@tcp(127.0.0.1:3309)/public_catalog?charset=utf8mb4&parseTime=True&loc=Local")
	db1, err := gorm.Open(mysql.Open("root:mysql@tcp(127.0.0.1:3309)/public_catalog?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	//db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB1 = db1
	return nil
}

func close() error {
	sqlDB, err := DB1.DB()
	if err != nil {
		// Handle error
	}
	sqlDB.Close()
	return nil
}
