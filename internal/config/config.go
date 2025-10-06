package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	Database string
	Port     string
	SmtpHost string
	SmtpPort string
	EmailFrom string
	EmailPassword string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system environment variables.")
	}

	return &Config{
		MongoURI: getEnv("MONGO_URI", "mongodb://localhost:27017"),
		Database: getEnv("MONGO_DB", "notification_db"),
		Port:    getEnv("PORT","8081"),
		SmtpHost: getEnv("SmtpHost",""),
		SmtpPort: getEnv("SmtpPort",""),
		EmailFrom: getEnv("EmailFrom",""),
		EmailPassword: getEnv("EmailPassword",""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}