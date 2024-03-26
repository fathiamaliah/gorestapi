package main

import (
	"gorestapi/configs"
	"gorestapi/controllers/productcontroller"
	"gorestapi/database"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	configs.InitEnvConfigs()

	// Set up database configuration using environment variables
	databaseConfig := database.DBConfig{
		DatabaseUser:     configs.EnvConfigs.DatabaseUser,
		DatabasePassword: configs.EnvConfigs.DatabasePassword,
		DatabaseName:     configs.EnvConfigs.DatabaseName,
		DatabaseUrl:      configs.EnvConfigs.DatabaseUrl,
		DatabasePort:     configs.EnvConfigs.DatabasePort,
	}

	// Establish a connection to the MySQL database
	db, _ := database.NewConnectionToMySQL(&databaseConfig)
	database.InitializeDatabase(db)

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products/:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.Run()

}
