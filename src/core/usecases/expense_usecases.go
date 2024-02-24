package expense_usecases

import (
	core "github.com/eduardo-paes/cashflow/core/entities"
	"github.com/eduardo-paes/cashflow/core/ports"
)

type UseCase struct {
	Repository core.ExpenseRepository
}

// New returns contract implementation of ExpenseUseCases
func New(repository core.ExpenseRepository) core.ExpenseUseCases {
	return &UseCase{
		Repository: repository,
	}
}

// Create implements core.ExpenseUseCases.
func (u *UseCase) Create(expense *ports.ExpenseInput) (*core.Expense, error) {
	newExpense, err := u.Repository.Create(expense)

	if err != nil {
		return nil, err
	}

	return newExpense, nil
}

// Delete implements core.ExpenseUseCases.
func (u *UseCase) Delete(id int64) (*core.Expense, error) {
	expenseDeleted, err := u.Repository.Delete(id)

	if err != nil {
		return nil, err
	}

	return expenseDeleted, nil
}

// GetOneOrMany implements core.ExpenseUseCases.
func (u *UseCase) GetOneOrMany(skip int, take int, id ...int64) ([]core.Expense, error) {
	expenses, err := u.Repository.GetOneOrMany(skip, take, id...)

	if err != nil {
		return nil, err
	}

	return expenses, nil
}

// Update implements core.ExpenseUseCases.
func (u *UseCase) Update(id int64, expense *ports.ExpenseInput) (*core.Expense, error) {
	expenseDeleted, err := u.Repository.Update(id, expense)

	if err != nil {
		return nil, err
	}

	return expenseDeleted, nil
}
