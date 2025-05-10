package token

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	key         string
	identityKey string
	expiration  time.Duration
}

var (
	config = Config{
		key:         "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
		identityKey: "identityKey",
		expiration:  2 * time.Hour,
	}

	once sync.Once
)

// Init sets the packag-level configuration config, which is used for token issuance and parsing in the subsequent parts of this package.
func Init(key string, identityKey string, expiration time.Duration) {
	once.Do(func() {
		if key != "" {
			config.key = key
		}
		if identityKey != "" {
			config.identityKey = identityKey
		}
		if expiration != 0 {
			config.expiration = expiration
		}
	})
}

// Parse uses the specified key key to parse the token. If the parsing is successful, it returns the token context; otherwise, it reports an error.
func Parse(tokenString string, key string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(key), nil
	})

	if err != nil {
		return "", err
	}

	var identityKey string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if key, exists := claims[config.identityKey]; exists {
			if identity, valid := key.(string); valid {
				identityKey = identity
			}
		}
	}

	if identityKey == "" {
		return "", jwt.ErrSignatureInvalid
	}

	return identityKey, nil
}

// ParseRequest retrieves the token from the request header and passes it to the Parse function to parse the token.
func ParseRequest(c *gin.Context) (string, error) {
	header := c.Request.Header.Get("Authorization")

	if len(header) == 0 {
		return "", errors.New("the length of the `Authorization` header is zero")
	}

	var token string
	fmt.Sscanf(header, "Bearer %s", &token)

	return Parse(token, config.key)
}

// Sign uses jwtSecret to issue tokens, and the claims of the token will store the incoming subject.
func Sign(identityKey string) (string, time.Time, error) {
	expireAt := time.Now().Add(config.expiration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,
		"exp":              expireAt.Unix(),
		"iat":              time.Now().Unix(),
		"nbf":              time.Now().Unix(),
	})

	if config.key == "" {
		return "", time.Time{}, jwt.ErrInvalidKey
	}

	tokenString, err := token.SignedString([]byte(config.key))
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expireAt, nil
}
