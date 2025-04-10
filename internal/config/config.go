package config

import (
	"os"
	"path/filepath"
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

	// Get database directory from environment or use default
	dbDir := os.Getenv("DATABASE_DIR")
	if dbDir == "" {
		dbDir = "./data"
	}

	// Create database directory if it doesn't exist
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		panic(err)
	}

	// For SQLite, we'll use a different database file for each environment
	dbName := "waitlist.db"
	if env == "prod" {
		dbName = "waitlist_prod.db"
	}

	dbPath := filepath.Join(dbDir, dbName)

	return &Config{
		Environment: env,
		Port:        port,
		DatabaseURL: dbPath,
	}
}
