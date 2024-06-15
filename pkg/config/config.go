package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	GoogleMapsAPIKey string
}

func LoadConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	var cfg Config
	err := viper.Unmarshal(&cfg)
	return cfg, err
}
