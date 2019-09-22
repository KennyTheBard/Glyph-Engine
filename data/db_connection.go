package data

import (
	"github.com/jinzhu/gorm"
	// this is need in order to be able to use postgres syntax
	_ "github.com/jinzhu/gorm/dialects/postgres"

	model "../model"
)

// DB is the reference to the ORM that comunicates with th data source
var DB *gorm.DB

// Init prepares the link with the data source
func Init(cleanStart bool) {
	var err error
	DB, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=moonshine_square password=postgres sslmode=disable")
	if err != nil {
		panic("failed to connect database")
	}

	if cleanStart {
		DB.DropTableIfExists(&model.StoryModel{})
		DB.DropTableIfExists(&model.ChoiceModel{})
		DB.DropTableIfExists(&model.ItemModel{})
		DB.DropTableIfExists(&model.ItemCost{})
		DB.DropTableIfExists(&model.ItemReward{})
		DB.DropTableIfExists(&model.ItemRequirement{})
	}

	DB.AutoMigrate(&model.StoryModel{})
	DB.AutoMigrate(&model.ChoiceModel{})
	DB.AutoMigrate(&model.ItemModel{})
	DB.AutoMigrate(&model.ItemCost{})
	DB.AutoMigrate(&model.ItemReward{})
	DB.AutoMigrate(&model.ItemRequirement{})
}

func Close() {
	DB.Close()
}
