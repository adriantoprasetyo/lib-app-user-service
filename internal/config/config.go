package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServicePort string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPass      string
	DBName      string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}
	return Config{
		ServicePort: os.Getenv("SERVICE_PORT"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBUser:      os.Getenv("DB_USER"),
		DBPass:      os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
	}
}

func (c Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", c.DBHost, c.DBUser, c.DBPass, c.DBName, c.DBPort)
}
