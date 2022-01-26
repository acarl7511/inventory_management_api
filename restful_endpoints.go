package Inventory

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

func Create_item(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "stocks.db")
	if err != nil {
		log.Fatal("Couldn't conect to database...")
	}

	var new Stock
	byteArray, e := ioutil.ReadAll(r.Body)
	if e != nil {
		log.Fatal("Error parsing response")
	}
	_ = json.Unmarshal(byteArray, &new)

	db.Create(&new)

	db.Save(&new)
}

func Get_item(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "stocks.db")
	if err != nil {
		log.Fatal("Couldn't conect to database...")
	}
	var list []Stock

	db.Model(&list[0]).Find(&list)
	byteArray, e := ioutil.ReadAll(r.Body)
	if e != nil {
		log.Fatal("Error parsing response")
	}
	search := string(byteArray)
	for _, i := range list {
		if search == i.Categories {
			e = json.NewEncoder(w).Encode(i)
			if e != nil {
				log.Fatal("Error Encoding Response")
			}
		}
	}
}

func Update_item(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "stocks.db")
	if err != nil {
		log.Fatal("Couldn't conect to database...")
	}

	var (
		list []Stock
		new  Stock
	)
	byteArray, e := ioutil.ReadAll(r.Body)
	if e != nil {
		log.Fatal("Error parsing response")
	}
	_ = json.Unmarshal(byteArray, &new)

	db.Model(&list).Find(&list)

	for _, i := range list {
		if i.Categories == new.Categories {
			db.Model(&new).Update(&new)
			db.Save(&new)
		}
	}
}

func Delete_item(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "stocks.db")
	if err != nil {
		log.Fatal("Couldn't conect to database...")
	}

	var (
		list []Stock
		new  Stock
	)
	byteArray, e := ioutil.ReadAll(r.Body)
	if e != nil {
		log.Fatal("Error parsing response")
	}
	_ = json.Unmarshal(byteArray, &new)

	db.Model(&list).Find(&list)

	for _, i := range list {
		if i.Categories == new.Categories {
			db.Model(&new).Delete(&new)
			db.Save(&new)
		}
	}
}
