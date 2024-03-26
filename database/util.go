package database

import (
	"gorestapi/models"

	"github.com/jinzhu/gorm"
)

func InitializeDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.Product{})
	DB = db
}
