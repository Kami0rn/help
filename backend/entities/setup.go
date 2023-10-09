package entities

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("SoyJuu.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(
		&Customer{},
		// &Delivery{},
		// &Food{},
		// &FoodMenu{},
		// &History{},
		// &Menu{},
		// &Order{},
		// &OrderFood{},
		// &Rider{},
		// &Payment{},
		// &Salary{},
		// &Wallet{},
	)

	db = database
}



