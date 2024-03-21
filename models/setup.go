package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	database, err := gorm.Open(mysql.Open("root:Suppasit@tcp(localhost:3306)/gorestapi"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Product{})

	DB = database

}
