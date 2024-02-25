package ioc

import (
	core "github.com/eduardo-paes/cashflow/core/entities"
	services "github.com/eduardo-paes/cashflow/core/services"
	repositories "github.com/eduardo-paes/cashflow/infra/data/repositories"
	controllers "github.com/eduardo-paes/cashflow/infra/http/controllers"
	"gorm.io/gorm"
)

// ConfigExpenseDI return a ExpenseService abstraction with dependency injection configuration
func ConfigExpenseDI(db *gorm.DB) core.ExpenseService {
	expenseRepository := repositories.NewExpenseRepository(db)
	expenseUseCase := services.NewExpenseService(expenseRepository)
	expenseService := controllers.NewExpenseController(expenseUseCase)

	return expenseService
}
