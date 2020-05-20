package configuration

import (
	"log"

	"github.com/spf13/viper"
)

// Configuration used for the main app
type Configuration struct {
	DbHost      string
	DbUser      string
	DbPass      string
	DbName      string
	GrpcURL     string
	BMSEndpoint string
}

func GetConfig() Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration Configuration
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	return configuration
}
