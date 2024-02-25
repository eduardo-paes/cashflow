package expenses

import (
	"net/http"
	"time"
)

type Expense struct {
	ID          int64      `json:"id"`
	Category    string     `json:"category"`
	Type        int        `json:"type"`
	Description string     `json:"description"`
	Amount      float64    `json:"amount"`
	Date        time.Time  `json:"date"`
	CreatedAt   time.Time  `json:"createdAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

type ExpenseService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Update(response http.ResponseWriter, request *http.Request)
	Delete(response http.ResponseWriter, request *http.Request)
	GetOneOrMany(response http.ResponseWriter, request *http.Request)
}

type ExpenseUseCases interface {
	Create(input *ExpenseInput) (*Expense, error)
	Update(id int64, input *ExpenseInput) (*Expense, error)
	Delete(id int64) (*Expense, error)
	GetOneOrMany(skip int, take int, id ...int64) ([]Expense, error)
}

type ExpenseRepository interface {
	Create(input *ExpenseInput) (*Expense, error)
	Update(id int64, input *ExpenseInput) (*Expense, error)
	Delete(id int64) (*Expense, error)
	GetOneOrMany(skip int, take int, id ...int64) ([]Expense, error)
}
