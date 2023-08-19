package config

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

// type ConfigShape struct {
// 	AccessTokenSecret string
// }
// var Config ConfigShape

func GetAccessTokenSecret() []byte {
	return []byte(os.Getenv("AccessTokenSecret"))
}

func LoadEnv(name string) {
	_, filename, _, _ := runtime.Caller(0)
	envPath := filepath.Join(filepath.Dir(filename), "../../", name)

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
