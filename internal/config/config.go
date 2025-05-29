package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Config struct {
	Port        string `env:"PORT" default:"8080"`
	DatabaseDSN string `env:"DATABASE_DSN"`
	JWTSecret   string `env:"JWT_SECRET" default:"your-secret-key-here"`
	DBHost      string `env:"DB_HOST" default:"localhost"`
	DBPort      string `env:"DB_PORT" default:"5432"`
	DBUser      string `env:"DB_USER" default:"postgres"`
	DBPassword  string `env:"DB_PASSWORD" default:"postgres"`
	DBName      string `env:"DB_NAME" default:"payment_service"`
	LogLevel    string `env:"LOG_LEVEL"`
	DB          *sql.DB
}

func LoadConfig() (*Config, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	// Create database connection string
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Connect to database
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &Config{
		Port:       os.Getenv("PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
		DB:         db,
	}, nil
}
