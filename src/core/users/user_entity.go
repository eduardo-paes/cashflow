package users

import (
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type UserService interface {
	Login(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	GetOne(context *gin.Context)
}

type UserUseCases interface {
	Login(input *AuthInput) (*AuthOutput, error)
	Create(input *UserInput) (*User, error)
	Update(id int64, input *UserInput) (*User, error)
	Delete(id int64) (*User, error)
	GetOne(id ...int64) (*User, error)
}

type UserRepository interface {
	Login(input *AuthInput) (*User, error)
	Create(input *UserInput) (*User, error)
	Update(id int64, input *UserInput) (*User, error)
	Delete(id int64) (*User, error)
	GetOne(id ...int64) (*User, error)
}

type AuthService interface {
	GenerateToken(userId int64, userName string) (string, error)
	HashPassword(password string) (string, error)
}
