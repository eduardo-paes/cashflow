package expenses

type ExpenseServices struct {
	Repository ExpenseRepository
}

// New returns contract implementation of ExpenseUseCases
func NewExpenseService(repository ExpenseRepository) ExpenseUseCases {
	return &ExpenseServices{
		Repository: repository,
	}
}

// Create implements core.ExpenseUseCases.
func (u *ExpenseServices) Create(expense *ExpenseInput) (*Expense, error) {
	newExpense, err := u.Repository.Create(expense)

	if err != nil {
		return nil, err
	}

	return newExpense, nil
}

// Delete implements core.ExpenseUseCases.
func (u *ExpenseServices) Delete(id int64) (*Expense, error) {
	expenseDeleted, err := u.Repository.Delete(id)

	if err != nil {
		return nil, err
	}

	return expenseDeleted, nil
}

// GetOneOrMany implements core.ExpenseUseCases.
func (u *ExpenseServices) GetOneOrMany(skip int, take int, id ...int64) ([]Expense, error) {
	expenses, err := u.Repository.GetOneOrMany(skip, take, id...)

	if err != nil {
		return nil, err
	}

	return expenses, nil
}

// Update implements core.ExpenseUseCases.
func (u *ExpenseServices) Update(id int64, expense *ExpenseInput) (*Expense, error) {
	expenseDeleted, err := u.Repository.Update(id, expense)

	if err != nil {
		return nil, err
	}

	return expenseDeleted, nil
}
