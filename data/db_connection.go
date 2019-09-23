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

	migrate(&model.StoryModel{}, cleanStart)
	migrate(&model.ChoiceModel{}, cleanStart)
	migrate(&model.ItemModel{}, cleanStart)
	migrate(&model.ItemCost{}, cleanStart)
	migrate(&model.ItemReward{}, cleanStart)
	migrate(&model.ItemRequirement{}, cleanStart)
}

func migrate(dbmodel interface{}, cleanStart bool) {
	if cleanStart {
		DB.DropTableIfExists(&dbmodel)
	}
	DB.AutoMigrate(&dbmodel)
}

func Close() {
	DB.Close()
}
