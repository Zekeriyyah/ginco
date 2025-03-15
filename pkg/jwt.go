package pkg

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret_key = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func GeneratJWT(userID uint) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret_key)
}

// validate if token is valid and retrieve the claims
func ValidateJWT(tokenStr string) (*Claims, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	// return err if token is expired or not valid
	if err != nil {
		return nil, err
	}

	// extract claim to return it
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil

}
