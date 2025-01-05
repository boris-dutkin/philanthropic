package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func LoadEnv() {
	directory, _ := os.Getwd()

	viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("%s/config", directory))
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read the config file.")
	}
}
