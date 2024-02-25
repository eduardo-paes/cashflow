package users

import (
	"github.com/eduardo-paes/cashflow/core/users"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository returns contract implementation of UserRepository
func NewUserRepository(db *gorm.DB) users.UserRepository {
	return &UserRepository{db: db}
}

// Login implements core.UserRepository.
func (r *UserRepository) Login(input *users.AuthInput) (*users.User, error) {
	var dto User
	if err := r.db.Where("email = ? AND password = ?", input.Email, input.Password).First(&dto).Error; err != nil {
		return nil, err
	}

	user := dto.UserConvertToEntity()
	return &user, nil
}

// Create implements core.UserRepository.
func (r *UserRepository) Create(input *users.UserInput) (*users.User, error) {
	dto := User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		UpdateAt:  time.Now(),
		CreatedAt: time.Now(),
	}

	if err := r.db.Create(&dto).Error; err != nil {
		return nil, err
	}

	user := dto.UserConvertToEntity()
	return &user, nil
}

// Delete implements core.UserRepository.
func (r *UserRepository) Delete(id int64) (*users.User, error) {
	var dto User

	// Soft delete the user by updating the DeletedAt field
	if err := r.db.Delete(&dto, id).Error; err != nil {
		return nil, err
	}

	user := dto.UserConvertToEntity()
	return &user, nil
}

// GetOne implements core.UserRepository.
func (r *UserRepository) GetOne(id ...int64) (*users.User, error) {
	var dto User
	if err := r.db.First(&dto, id).Error; err != nil {
		return nil, err
	}

	user := dto.UserConvertToEntity()
	return &user, nil
}

// Update implements core.UserRepository.
func (r *UserRepository) Update(id int64, input *users.UserInput) (*users.User, error) {
	var dto User
	if err := r.db.First(&dto, id).Error; err != nil {
		return nil, err
	}

	dto.Name = input.Name
	dto.Email = input.Email
	dto.Password = input.Password
	dto.UpdatedAt = time.Now()

	if err := r.db.Save(dto).Error; err != nil {
		return nil, err
	}

	user := dto.UserConvertToEntity()

	return &user, nil
}
