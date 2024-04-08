package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("secrets")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
