package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"libraryproject/config"
)

var Db *sql.DB

func InitDB(envs *config.Config) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		envs.DBHost, envs.DBPort, envs.DBUser, envs.DBPassword, envs.DBName, envs.DBSSLMode)
	fmt.Println("Trying to connect with db")

	var err error
	Db, err = sql.Open("postgres", dsn)
	if err != nil {
		fmt.Println("Error opening DB:", err)
		return err
	}
	if err := Db.Ping(); err != nil {
		fmt.Println("Error pinging DB:", err)
		Db.Close()
		return err
	}
	fmt.Println("Connected to DB successfully!")

	// Автоматическая миграция
	if err := runMigrations(); err != nil {
		fmt.Println("Migration failed:", err)
		return err
	}

	return nil
}

func runMigrations() error {
	query := `
	CREATE TABLE IF NOT EXISTS authors (
		id SERIAL PRIMARY KEY,
		full_name VARCHAR(255) NOT NULL,
		specialization VARCHAR(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		genre VARCHAR(100),
		isbn_code INTEGER,
		author_id INTEGER REFERENCES authors(id) ON DELETE CASCADE
	);

	CREATE TABLE IF NOT EXISTS readers (
		id SERIAL PRIMARY KEY,
		full_name VARCHAR(255) NOT NULL
	);

	CREATE TABLE IF NOT EXISTS book_requests (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		genre VARCHAR(100),
		isbn_code INTEGER,
		author_name VARCHAR(255),
		author_id INTEGER REFERENCES authors(id),
		status VARCHAR(50) DEFAULT 'pending'
	);

	CREATE TABLE IF NOT EXISTS author_users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		author_id INTEGER NOT NULL REFERENCES authors(id) ON DELETE CASCADE
	);
	`

	_, err := Db.Exec(query)
	if err != nil {
		fmt.Println("Migration error:", err)
	}
	return err
}
