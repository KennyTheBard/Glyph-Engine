package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type MasterModel struct {
	ID     uint
	Slaves []SlaveModel `gorm:"one2many:MasterRefer;association_foreignkey:ID"`
}

type SlaveModel struct {
	ID          uint
	MasterRefer uint
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=moonshine_square password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.DropTableIfExists(&MasterModel{})
	db.AutoMigrate(&MasterModel{})

	db.DropTableIfExists(&SlaveModel{})
	db.AutoMigrate(&SlaveModel{})
	defer db.Close()

	var master MasterModel
	db.Save(&master)

	slave := SlaveModel{MasterRefer: master.ID}
	db.Save(&slave)

	{
		var slaves []SlaveModel
		slaves = append(slaves, slave)
		master.Slaves = slaves
		db.Save(&master)
	}

	var slaves []SlaveModel
	db.Model(&master).Related(&slaves, "MasterRefer")
	fmt.Println(master)
	fmt.Println(slaves)
}
