package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	model "../model"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=moonshine_square password=postgres sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(&model.Story{})
}
