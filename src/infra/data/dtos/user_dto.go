package dtos

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID        int64          `gorm:"primary_key" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Email     string         `gorm:"column:email" json:"email"`
	Password  string         `gorm:"column:password" json:"password"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"createdAt"`
	UpdateAt  time.Time      `gorm:"column:updated_at" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}

// TableName specifies the table name for the User entity
func (User) TableName() string {
	return "users"
}
