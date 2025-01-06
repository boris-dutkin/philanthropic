package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name        string       `json:"name" gorm:"type:varchar(50); unique; unique_index;"`
	Description string       `json:"description" gorm:"type:varchar(50);"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
