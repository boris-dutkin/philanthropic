package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	env := viper.Get("env")
	host := viper.Get(fmt.Sprintf("%s.db.host", env))
	port := viper.Get(fmt.Sprintf("%s.db.port", env))
	user := viper.Get(fmt.Sprintf("%s.db.user", env))
	password := viper.Get(fmt.Sprintf("%s.db.password", env))
	database := viper.Get(fmt.Sprintf("%s.db.name", env))

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user,
		password,
		database,
		host,
		port,
	)

	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})

	if err != nil {
		panic("Failed to connect to the database.")
	}

	DB = connection
}
