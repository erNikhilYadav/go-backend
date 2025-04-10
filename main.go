package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nikhilyadav/go-backend/internal/config"
	"github.com/nikhilyadav/go-backend/internal/response"
)

type WaitlistEntry struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

var db *sql.DB
var cfg *config.Config

func initDB() {
	var err error
	cfg = config.LoadConfig()

	log.Printf("Starting server in %s environment", cfg.Environment)

	db, err = sql.Open("sqlite3", cfg.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// Create waitlist table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS waitlist (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT UNIQUE NOT NULL,
		name TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}

func addToWaitlist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("only POST method is allowed"))
		return
	}

	var entry WaitlistEntry
	if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
		response.ErrorResponse(w, http.StatusBadRequest, "Invalid request body", err)
		return
	}

	// Basic validation
	if entry.Email == "" || entry.Name == "" {
		response.ErrorResponse(w, http.StatusBadRequest, "Validation failed", errors.New("email and name are required"))
		return
	}

	// Insert into database
	stmt, err := db.Prepare("INSERT INTO waitlist (email, name) VALUES (?, ?)")
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Database error", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(entry.Email, entry.Name)
	if err != nil {
		response.ErrorResponse(w, http.StatusConflict, "Email already exists", err)
		return
	}

	response.SuccessResponse(w, http.StatusCreated, "Successfully added to waitlist", entry)
}

func getWaitlist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.ErrorResponse(w, http.StatusMethodNotAllowed, "Method not allowed", errors.New("only GET method is allowed"))
		return
	}

	rows, err := db.Query("SELECT id, email, name, created_at FROM waitlist ORDER BY created_at DESC")
	if err != nil {
		response.ErrorResponse(w, http.StatusInternalServerError, "Database error", err)
		return
	}
	defer rows.Close()

	var entries []WaitlistEntry
	for rows.Next() {
		var entry WaitlistEntry
		if err := rows.Scan(&entry.ID, &entry.Email, &entry.Name, &entry.CreatedAt); err != nil {
			response.ErrorResponse(w, http.StatusInternalServerError, "Database error", err)
			return
		}
		entries = append(entries, entry)
	}

	response.SuccessResponse(w, http.StatusOK, "Successfully retrieved waitlist", entries)
}

func main() {
	initDB()
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/api/waitlist", addToWaitlist)
	mux.HandleFunc("/api/waitlist/list", getWaitlist)

	log.Printf("Server starting on :%s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, mux))
}
