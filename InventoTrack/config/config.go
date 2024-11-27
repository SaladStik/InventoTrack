package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to load configuration")
	}
}

func Get(key string) string {
	return viper.GetString(key)
}
