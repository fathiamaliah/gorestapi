package configs

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	DatabaseUser     string `mapstructure:"DB_USER"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
	DatabaseName     string `mapstructure:"DB_NAME"`
	DatabaseUrl      string `mapstructure:"DB_URL"`
	DatabasePort     string `mapstructure:"DB_PORT"`
}

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}

func loadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
