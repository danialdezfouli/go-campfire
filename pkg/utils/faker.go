package utils

import (
	"github.com/go-faker/faker/v4"
)

func GenerateRandomSubdomain() string {
	var sample struct {
		Subdomain string `faker:"username"`
	}
	faker.FakeData(&sample)
	return sample.Subdomain
}
