package adapter

import (
	"fmt"

	"github.com/choonhong/hotel-data-merge/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	fmt.Println("connect to database successfully")

	return db, db.AutoMigrate(&model.Hotel{})
}
