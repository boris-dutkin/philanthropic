package role

import (
	"github.com/boris-dutkin/philanthropic/internal/permission"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	Name        string                  `json:"name" gorm:"type:varchar(50); unique; unique_index;"`
	Description string                  `json:"description" gorm:"type:varchar(50);"`
	Permissions []permission.Permission `gorm:"many2many:role_permissions;"`
}
