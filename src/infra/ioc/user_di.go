package ioc

import (
	core "github.com/eduardo-paes/cashflow/core/entities"
	services "github.com/eduardo-paes/cashflow/core/services"
	repositories "github.com/eduardo-paes/cashflow/infra/data/repositories"
	controllers "github.com/eduardo-paes/cashflow/infra/http/controllers"
	"github.com/eduardo-paes/cashflow/infra/security"
	"gorm.io/gorm"
)

// ConfigUserDI return a UserService abstraction with dependency injection configuration
func ConfigUserDI(db *gorm.DB, jwtKey string) core.UserService {
	userRepository := repositories.NewUserRepository(db)
	authService := security.NewAuthServices(jwtKey)
	userUseCase := services.NewUserService(userRepository, authService)
	userService := controllers.NewUserController(userUseCase)

	return userService
}
