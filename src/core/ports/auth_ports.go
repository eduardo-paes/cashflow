package ports

import (
	"encoding/json"
	"io"
)

type AuthInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthOutput struct {
	Token    string `json:"token"`
	UserName string `json:"userName"`
	UserId   int64  `json:"userId"`
}

// Converts JSON body request to a AuthInput struct
func FromJSONAuthInput(body io.Reader) (*AuthInput, error) {

	// Create a AuthInput struct
	authInput := AuthInput{}

	// Decode JSON body into a AuthInput struct
	if err := json.NewDecoder(body).Decode(&authInput); err != nil {
		return nil, err
	}

	// Return the AuthInput struct
	return &authInput, nil
}
