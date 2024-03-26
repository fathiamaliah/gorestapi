package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBConfig struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DatabaseUrl      string
	DatabasePort     string
}

// ConnectMySQL- Connect to MySQL database
var DB *gorm.DB

func NewConnectionToMySQL(config *DBConfig) (*gorm.DB, error) {
	databaseURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseUser,
		config.DatabasePassword,
		config.DatabaseUrl,
		config.DatabasePort,
		config.DatabaseName)
	db, err := gorm.Open("mysql", databaseURI)
	if err != nil {
		fmt.Printf("Error connecting to database: %v\nConfig: %+v\n", err, config)
		return nil, err
	}

	fmt.Println("Successfully connected to database:", config.DatabaseName)

	return db, nil

}

// CloseDB - closes the database connection.
func CloseDB(db *gorm.DB) {
	if db != nil {
		db.Close()
		fmt.Printf("Database connection closed.")
	}
}
