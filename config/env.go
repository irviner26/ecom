package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost    string
	Port          string
	DBUser        string
	DBPassword    string
	DBAddress     string
	DBName        string
	JWTExpiration int64
	JWTSecret     string
}

var Global Config = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:    getEnv("PUBLIC_HOST", "localhost"),
		Port:          getEnv("PORT", "8080"),
		DBUser:        getEnv("DB_USER", "postgres"),
		DBPassword:    getEnv("DB_PASSWORD", "postgres"),
		DBName:        getEnv("DB_NAME", "postgres"),
		DBAddress:     getEnv("DB_ADDRESS", "localhost:5432"),
		JWTExpiration: getEnvAsInt("JWT_EXPIRATION", 3600*24*7),
		JWTSecret:     getEnv("JWT_SECRET", "secret-not-set!"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
