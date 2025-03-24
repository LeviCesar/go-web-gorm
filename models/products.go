package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	Id          uint `gorm:"primaryKey;default:auto_random()"`
	Name        string
	Description string
	Price       float64
	Quantity    uint16
}

func Connect() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Product{})

	return db
}

func (product *Product) Create() bool {
	db := Connect()

	result := db.Create(&product)
	return result.RowsAffected == 1
}

func (product *Product) Read() []Product {
	db := Connect()

	var productList []Product
	db.Find(&productList)

	return productList
}

// func (product *Product) Update() {
// 	db := Connect()

// 	var model Product
// 	result := db.Model(&model).Updates(product)
// }
