package authn

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// IToken defines methods to implement a generic token
type IToken interface {
	GetToken() string

	GetTokenType() string

	// Get token expiration timestamp
	GetExpireAt() int64

	EncodeToJSON() ([]byte, error)
}

// Authenticator defines methods used for tekon processing
type AuthenTicator interface {
	// Sign is used to generate a tekon
	Sign(ctx context.Context, userID string) (IToken, error)

	// Destroy is used to destroy a taken
	Destroy(ctx context.Context, accessToken string) error

	// ParseClaims parse the token and return the claims
	ParseClaims(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error)

	// Release used to release the requested resource
	Release() error
}

// Encrypt encrypts the plain text with bcrypt
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
