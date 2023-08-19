package token

import (
	"errors"

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
	if err != nil {
		return "", err
	}
	return stringToken, err
}

func Validate(tokenString string, secretKey []byte) error {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("unknown error - token is not valid")
	}

	return nil
}
