package data

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	model "../model"
)

var DB *gorm.DB

func Init() {
	DB, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=moonshine_square password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	DB.AutoMigrate(&model.Story{})
}
