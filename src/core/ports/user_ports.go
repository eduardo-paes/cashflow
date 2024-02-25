package ports

import (
	"encoding/json"
	"io"
)

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Converts JSON body request to a UserInput struct
func FromJSONCreateUser(body io.Reader) (*UserInput, error) {

	// Create a UserInput struct
	createUser := UserInput{}

	// Decode JSON body into a UserInput struct
	if err := json.NewDecoder(body).Decode(&createUser); err != nil {
		return nil, err
	}

	// Return the UserInput struct
	return &createUser, nil
}
