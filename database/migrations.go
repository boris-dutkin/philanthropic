package database

import (
	"github.com/boris-dutkin/philanthropic/internal/permission"
	"github.com/boris-dutkin/philanthropic/internal/role"
	"github.com/boris-dutkin/philanthropic/internal/user"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&user.User{},
		&role.Role{},
		&permission.Permission{},
	)
}
