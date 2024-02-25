package security

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	core "github.com/eduardo-paes/cashflow/core/users"
	"time"
)

type AuthServices struct {
	secretKey string
	salt      string
}

// CustomClaims represents the claims you want to include in the JWT
type CustomClaims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"username"`
	jwt.StandardClaims
}

// NewAuthServices returns a new instance of AuthServices.
func NewAuthServices(jwtKey string, salt string) core.AuthService {
	return &AuthServices{
		secretKey: jwtKey,
		salt:      salt,
	}
}

// GenerateToken implements core.AuthService.
func (u *AuthServices) GenerateToken(userId int64, userName string) (string, error) {
	// Define the claims
	claims := CustomClaims{
		UserID:   userId,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Set expiration time
			IssuedAt:  time.Now().Unix(),                     // Set issuance time
		},
	}

	// Create a new token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	secretKey := []byte(u.secretKey)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// HashPassword implements core.AuthService.
func (u *AuthServices) HashPassword(password string) (string, error) {
	// Concatenate the password and salt
	combined := password + u.salt

	// Hash the combined string using SHA-256
	hash := sha256.New()
	_, err := hash.Write([]byte(combined))
	if err != nil {
		return "", err
	}

	// Get the hexadecimal representation of the hash
	hashedPassword := hex.EncodeToString(hash.Sum(nil))

	return hashedPassword, nil
}
