package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		DB *DB
	}
	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}
)

func New() (*Container, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading env file")
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	if db.Host == "" || db.Port == "" || db.User == "" || db.Password == "" || db.Name == "" {
		return nil, fmt.Errorf("missing required database configuration")
	}

	return &Container{db}, nil
}
