package repositories

import (
	"time"

	"gorm.io/gorm"

	core "github.com/eduardo-paes/cashflow/core/entities"
	"github.com/eduardo-paes/cashflow/core/ports"
	"github.com/eduardo-paes/cashflow/infra/data/dtos"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository returns contract implementation of UserRepository
func NewUserRepository(db *gorm.DB) core.UserRepository {
	return &UserRepository{db: db}
}

// Login implements core.UserRepository.
func (r *UserRepository) Login(input *ports.AuthInput) (*core.User, error) {
	var dto dtos.User
	if err := r.db.Where("email = ? AND password = ?", input.Email, input.Password).First(&dto).Error; err != nil {
		return nil, err
	}

	user := UserConvertToEntity(dto)
	return &user, nil
}

// Create implements core.UserRepository.
func (r *UserRepository) Create(input *ports.UserInput) (*core.User, error) {
	dto := dtos.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: time.Now(),
	}

	if err := r.db.Create(&dto).Error; err != nil {
		return nil, err
	}

	user := UserConvertToEntity(dto)
	return &user, nil
}

// Delete implements core.UserRepository.
func (r *UserRepository) Delete(id int64) (*core.User, error) {
	var dto dtos.User

	// Soft delete the user by updating the DeletedAt field
	if err := r.db.Delete(&dto, id).Error; err != nil {
		return nil, err
	}

	user := UserConvertToEntity(dto)
	return &user, nil
}

// GetOne implements core.UserRepository.
func (r *UserRepository) GetOne(id ...int64) (*core.User, error) {
	var dto dtos.User
	if err := r.db.First(&dto, id).Error; err != nil {
		return nil, err
	}

	user := UserConvertToEntity(dto)
	return &user, nil
}

// Update implements core.UserRepository.
func (r *UserRepository) Update(id int64, input *ports.UserInput) (*core.User, error) {
	var dto dtos.User
	if err := r.db.First(&dto, id).Error; err != nil {
		return nil, err
	}

	dto.Name = input.Name
	dto.Email = input.Email
	dto.Password = input.Password

	if err := r.db.Save(dto).Error; err != nil {
		return nil, err
	}

	user := UserConvertToEntity(dto)

	return &user, nil
}

func UserConvertToEntity(dto dtos.User) core.User {
	var deletedAt *time.Time
	if dto.DeletedAt != (gorm.DeletedAt{}) {
		deletedAt = &dto.DeletedAt.Time
	}

	return core.User{
		ID:        dto.ID,
		Name:      dto.Name,
		Email:     dto.Email,
		Password:  dto.Password,
		DeletedAt: deletedAt,
	}
}

func UserConvertToDto(entity core.User) dtos.User {
	var deletedAt gorm.DeletedAt
	if entity.DeletedAt != nil {
		deletedAt = gorm.DeletedAt{Time: *entity.DeletedAt}
	}

	return dtos.User{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		DeletedAt: deletedAt,
	}
}
