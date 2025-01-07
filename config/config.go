package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type Server struct {
	Port int
}

type Config struct {
	DB     DB
	Server Server
}

func New() *Config {
	directory, _ := os.Getwd()
	var config Config

	viper.AutomaticEnv()
	env := viper.GetString("env")

	viper.AddConfigPath(".")
	viper.SetConfigName(env)
	viper.AddConfigPath(fmt.Sprintf("%s/config", directory))
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to read the config file.")
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic("Failed to parse the config.")
	}

	return &config
}
