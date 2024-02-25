package expenses

import (
	core "github.com/eduardo-paes/cashflow/core/expenses"
	"time"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	db *gorm.DB
}

// NewExpenseRepository returns contract implementation of ExpenseRepository
func NewExpenseRepository(db *gorm.DB) core.ExpenseRepository {
	return &ExpenseRepository{db: db}
}

// Create implements core.ExpenseRepository.
func (r *ExpenseRepository) Create(input *core.ExpenseInput) (*core.Expense, error) {
	dto := Expense{
		Amount:      input.Amount,
		Category:    input.Category,
		Description: input.Description,
		Date:        input.Date,
		Type:        input.Type,
		CreatedAt:   time.Now(),
	}

	if err := r.db.Create(&dto).Error; err != nil {
		return nil, err
	}

	expense := ExpenseConvertToEntity(dto)
	return &expense, nil
}

// Delete implements core.ExpenseRepository.
func (r *ExpenseRepository) Delete(id int64) (*core.Expense, error) {
	var dto Expense

	// Soft delete the expense by updating the DeletedAt field
	if err := r.db.Delete(&dto, id).Error; err != nil {
		return nil, err
	}

	expense := ExpenseConvertToEntity(dto)
	return &expense, nil
}

// GetOneOrMany implements core.ExpenseRepository.
func (r *ExpenseRepository) GetOneOrMany(skip int, take int, id ...int64) ([]core.Expense, error) {
	var expenses []core.Expense
	query := r.db

	if len(id) > 0 {
		query = query.Where("id IN (?)", id)
	}

	query = query.Offset(skip).Limit(take).Where("deleted_at IS NULL").Find(&expenses)

	if query.Error != nil {
		return nil, query.Error
	}

	return expenses, nil
}

// Update implements core.ExpenseRepository.
func (r *ExpenseRepository) Update(id int64, input *core.ExpenseInput) (*core.Expense, error) {
	var dto Expense
	if err := r.db.First(&dto, id).Error; err != nil {
		return nil, err
	}

	dto.Amount = input.Amount
	dto.Category = input.Category
	dto.Date = input.Date
	dto.Description = input.Description
	dto.Type = input.Type

	if err := r.db.Save(dto).Error; err != nil {
		return nil, err
	}

	expense := ExpenseConvertToEntity(dto)

	return &expense, nil
}

// ExpenseConvertToEntity converts a dto to an entity
func ExpenseConvertToEntity(dto Expense) core.Expense {
	var deletedAt *time.Time
	if dto.DeletedAt != (gorm.DeletedAt{}) {
		deletedAt = &dto.DeletedAt.Time
	}

	return core.Expense{
		ID:          dto.ID,
		Amount:      dto.Amount,
		Category:    dto.Category,
		Description: dto.Description,
		Date:        dto.Date,
		Type:        dto.Type,
		CreatedAt:   dto.CreatedAt,
		DeletedAt:   deletedAt,
	}
}

// ExpenseConvertToDto converts an entity to a dto
func ExpenseConvertToDto(entity core.Expense) Expense {
	var deletedAt gorm.DeletedAt
	if entity.DeletedAt != nil {
		deletedAt = gorm.DeletedAt{Time: *entity.DeletedAt}
	}

	return Expense{
		ID:          entity.ID,
		Amount:      entity.Amount,
		Category:    entity.Category,
		Description: entity.Description,
		Date:        entity.Date,
		Type:        entity.Type,
		CreatedAt:   entity.CreatedAt,
		DeletedAt:   deletedAt,
	}
}
