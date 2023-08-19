package token

import (
	"campfire/internal/domain"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenType int

const (
	AccessToken JwtTokenType = iota + 1
	RefreshToken
)

type Claims struct {
	User         domain.UserId         `json:"user"`
	Organization domain.OrganizationId `json:"organization"`
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

func Validate(tokenString string, secretKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		return secretKey, nil
	})

	if !token.Valid {
		return nil, errors.New("token is expired")
	}

	if err != nil {
		return nil, err
	}

	claims, OK := token.Claims.(*Claims)
	if !OK {
		return nil, errors.New("unable to parse claims")
	}

	if claims.User == 0 {
		return nil, errors.New("no user property in claims")
	}

	return claims, nil
}

func Parse(tokenString string, secretKey []byte) (*Claims, error) {
	token, err := ExtractBearerToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, err := Validate(token, secretKey)
	if err != nil {
		return nil, err
	}

	return claims, nil
}
