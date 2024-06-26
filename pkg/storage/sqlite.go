package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	Conn *sql.DB
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewStorage() (*DB, error) {
	db, err := sql.Open("sqlite3", "./database/auth_service_data.db")
	if err != nil {
		fmt.Printf("Error connecting to database\n")
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Printf("Error pinging database connection object\n")
		return nil, err
	}

	dbObj := DB{Conn: db}
	dbObj.CreateTable()
	return &dbObj, nil
}

func (db *DB) CreateTable() error {
	creationCmd := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		entry_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		role TEXT DEFAULT "user" NOT NULL
	);`

	_, err := db.Conn.Exec(creationCmd)
	return err
}

func (db *DB) CloseDB() error {
	return db.Conn.Close()
}