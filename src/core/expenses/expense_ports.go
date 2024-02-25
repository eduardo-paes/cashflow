package expenses

import (
	"encoding/json"
	"io"
	"time"
)

type ExpenseInput struct {
	Category    string    `json:"category"`
	Type        int       `json:"type"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        time.Time `json:"date"`
}

// Converts JSON body request to a CreateExpense struct
func FromJSONCreateExpense(body io.Reader) (*ExpenseInput, error) {

	// Create a CreateExpense struct
	createExpense := ExpenseInput{}

	// Decode JSON body into a CreateExpense struct
	if err := json.NewDecoder(body).Decode(&createExpense); err != nil {
		return nil, err
	}

	// Return the CreateExpense struct
	return &createExpense, nil
}
