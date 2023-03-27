package utils

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

// type Config struct {
// 	Port int    `yaml:"port"`
// 	Gin  string `yaml:"gin_mode"`
// }

func LoadEnv(name string) {
	_, filename, _, _ := runtime.Caller(0)
	envPath := filepath.Join(filepath.Dir(filename), "../../", name)

	err := godotenv.Load(envPath)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
