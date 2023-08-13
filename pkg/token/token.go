package token

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenType int

const (
	AccessToken JwtTokenType = iota + 1
	RefreshToken
)

type Claims struct {
	UserID         any    `json:"userId"`
	OrganizationId any    `json:"organizationId"`
	Email          string `json:"email"`
	jwt.RegisteredClaims
}

func Generate(claims *Claims, signingKey []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(signingKey)
	return stringToken, err
}

func Validate(token string, secretKey []byte) (bool, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	if !t.Valid {
		return false, nil
	}

	return true, nil
}
