package database

import "github.com/boris-dutkin/philanthropic/models"

func Migrate() {
	DB.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Permission{},
	)
}
