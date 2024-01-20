package ioc

import (
	core "github.com/eduardo-paes/cashflow/core/entities"
	expense_usecases "github.com/eduardo-paes/cashflow/core/usecases"
	expense_repository "github.com/eduardo-paes/cashflow/infra/data/repositories"
	expense_controller "github.com/eduardo-paes/cashflow/infra/http/controllers"
	"gorm.io/gorm"
)

// ConfigExpenseDI return a ExpenseService abstraction with dependency injection configuration
func ConfigExpenseDI(db *gorm.DB) core.ExpenseService {
	expenseRepository := expense_repository.New(db)
	expenseUseCase := expense_usecases.New(expenseRepository)
	expenseService := expense_controller.New(expenseUseCase)

	return expenseService
}