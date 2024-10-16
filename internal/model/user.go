package model

import (
	"AI-Recruitment-backend/pkg/common"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username string        `gorm:"type:varchar(255);not null" json:"username"`
	Email    string        `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password string        `gorm:"type:varchar(255);not null" json:"password"`
	Role     common.Role   `gorm:"type:int;not null" json:"role"`
	Age      int           `gorm:"type:int;not null" json:"age"`
	Degree   common.Degree `gorm:"type:varchar(255);not null" json:"degree"`
}
