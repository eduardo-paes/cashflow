package users

import (
	services "github.com/eduardo-paes/cashflow/core/users"
	"github.com/eduardo-paes/cashflow/infra/security"
	"gorm.io/gorm"
)

// ConfigUserDI return a UserService abstraction with dependency injection configuration
func ConfigUserDI(db *gorm.DB, jwtKey string, salt string) services.UserService {
	userRepository := NewUserRepository(db)
	authService := security.NewAuthServices(jwtKey, salt)
	userUseCase := services.NewUserService(userRepository, authService)
	userService := NewUserController(userUseCase)

	return userService
}
