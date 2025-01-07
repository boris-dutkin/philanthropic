package permission

import "gorm.io/gorm"

type Permission struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(50); unique; unique_index;"`
	Description string `json:"description" gorm:"type:varchar(50);"`
}
