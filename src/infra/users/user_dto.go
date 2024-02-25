package users

import (
	core "github.com/eduardo-paes/cashflow/core/users"
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

// UserConvertToEntity converts a User entity to a User DTO
func (dto User) UserConvertToEntity() core.User {
	var deletedAt *time.Time
	if dto.DeletedAt != (gorm.DeletedAt{}) {
		deletedAt = &dto.DeletedAt.Time
	}

	var createAt *time.Time
	if dto.CreatedAt != (time.Time{}) {
		createAt = &dto.CreatedAt
	}

	var updateAt *time.Time
	if dto.UpdateAt != (time.Time{}) {
		updateAt = &dto.UpdateAt
	}

	return core.User{
		ID:        dto.ID,
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		CreatedAt: createAt,
		UpdatedAt: updateAt,
		DeletedAt: deletedAt,
	}
}

// UserConvertToDto converts a User DTO to a User entity
func UserConvertToDto(entity core.User) User {
	var deletedAt gorm.DeletedAt
	if entity.DeletedAt != nil {
		deletedAt = gorm.DeletedAt{Time: *entity.DeletedAt}
	}

	return User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		DeletedAt: deletedAt,
	}
}
