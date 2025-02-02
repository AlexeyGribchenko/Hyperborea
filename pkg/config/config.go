package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Config struct {
	Port      string `mapstructure:"PORT"`
	DBHost    string `mapstructure:"DB_HOST"`
	DBUser    string `mapstructure:"DB_USER"`
	DBPass    string `mapstructure:"DB_PASS"`
	DBName    string `mapstructure:"DB_NAME"`
	DBPort    string `mapstructure:"DB_PORT"`
	SecretKey string `mapstructure:"SECRET_KEY"`
}

type handler struct {
	DB *gorm.DB
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&c)

	if err != nil {
		return
	}

	return
}
