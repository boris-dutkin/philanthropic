package user

import (
	"github.com/boris-dutkin/philanthropic/internal/role"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"firstName" gorm:"type:varchar(20)"`
	LastName  string `json:"lastName" gorm:"type:varchar(20)"`
	Email     string `json:"email" gorm:"type:varchar(20); unique; unique_index;"`
	Password  string `json:"-"`
	RoleID    int
	Role      role.Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
