package config

import (
	"os"
)

type Config struct {
	Environment string
	Port        string
	DatabaseURL string
}

func LoadConfig() *Config {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "uat" // Default to UAT if not specified
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// For SQLite, we'll use a different database file for each environment
	dbPath := "./waitlist.db"
	if env == "prod" {
		dbPath = "./waitlist_prod.db"
	}

	return &Config{
		Environment: env,
		Port:        port,
		DatabaseURL: dbPath,
	}
}
