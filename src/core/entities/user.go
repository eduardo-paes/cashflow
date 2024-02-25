package core

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	"github.com/eduardo-paes/cashflow/core/ports"
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
	Login(response http.ResponseWriter, request *http.Request)
	Create(response http.ResponseWriter, request *http.Request)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	GetOne(context *gin.Context)
}

type UserUseCases interface {
	Login(input *ports.AuthInput) (*ports.AuthOutput, error)
	Create(input *ports.UserInput) (*User, error)
	Update(id int64, input *ports.UserInput) (*User, error)
	Delete(id int64) (*User, error)
	GetOne(id ...int64) (*User, error)
}

type UserRepository interface {
	Login(input *ports.AuthInput) (*User, error)
	Create(input *ports.UserInput) (*User, error)
	Update(id int64, input *ports.UserInput) (*User, error)
	Delete(id int64) (*User, error)
	GetOne(id ...int64) (*User, error)
}

type AuthService interface {
	GenerateToken(userId int64, userName string) (string, error)
	HashPassword(password string) (string, error)
}
