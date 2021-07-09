package model

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const dns = "deepanshu:12345@tcp(127.0.0.1:3306)/localDb?charset=utf8mb4&parseTime=True&loc=Local"

func ConnectionStablish() (*gorm.DB, error) {
	fmt.Printf("\n-----Connecting DATABASE----------")
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Connection Error", err)
	}

	return db, err
}
func Migration() {
	fmt.Print("\n---Migration Works--------")
	db, _ := ConnectionStablish()
	db.AutoMigrate(&Category{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Users{})
	db.AutoMigrate(&Cart{})
	db.AutoMigrate(&ProductPurchaseStock{})
	db.AutoMigrate(&Orders{})
	db.AutoMigrate(&Shipping{})
	db.AutoMigrate(&Address{})

}
