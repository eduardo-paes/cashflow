package expenses

import (
	core "github.com/eduardo-paes/cashflow/core/expenses"
	"gorm.io/gorm"
)

// ConfigExpenseDI return a ExpenseService abstraction with dependency injection configuration
func ConfigExpenseDI(db *gorm.DB) core.ExpenseService {
	expenseRepository := NewExpenseRepository(db)
	expenseUseCase := core.NewExpenseService(expenseRepository)
	expenseService := NewExpenseController(expenseUseCase)

	return expenseService
}
