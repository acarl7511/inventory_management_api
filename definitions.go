package Inventory

import (
	"github.com/jinzhu/gorm"
)

type Stock struct {
	gorm.Model
	Categories string           `json:"Categories"`
	Content    categoryContents `gorm:"many2many:stock_categoryContents"`
}

type categoryContents struct {
	Contents map[string]string `json:"contents"`
}
