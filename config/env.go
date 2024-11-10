package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost      string
	Port            string
	DBUser          string
	DBPassword      string
	DBAddress       string
	DBName          string
	JwtSecret       string
	JwtExpInSeconds int64
}

var Envs Config = initConfig()

func initConfig() Config {
	godotenv.Load()

	return Config{
		PublicHost:      getEnv("PUBLIC_HOST", "localhost"),
		Port:            ":" + getEnv("PORT", "8080"),
		DBUser:          getEnv("DB_USER", "mysql"),
		DBPassword:      getEnv("DB_PASSWORD", ""),
		DBAddress:       fmt.Sprintf("%s:%s", getEnv("DB_HOST", "localhost"), getEnv("DB_PORT", "3306")),
		DBName:          getEnv("DB_NAME", "postgres"),
		JwtSecret:       getEnv("JWT_SECRET", "top-secret"),
		JwtExpInSeconds: getEnvAsInt("JWT_EXP", 3600*24*7), // 7 days
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
		n, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}
		return n
	}

	return fallback
}
