package Inventory

import (
	"log"

	"github.com/jinzhu/gorm"
)

func Init_migration() {
	db, err := gorm.Open("sqlite3", "inventory.db")
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}
	defer db.Close()

	db.AutoMigrate(&Stock{}, &categoryContents{})
}
