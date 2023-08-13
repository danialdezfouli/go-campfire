package utils

import (
	"github.com/go-faker/faker/v4"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateRandomSubdomain() string {
	var sample struct {
		Subdomain string `faker:"username"`
	}
	faker.FakeData(&sample)
	return sample.Subdomain
}
