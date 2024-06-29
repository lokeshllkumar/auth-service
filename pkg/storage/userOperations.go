package storage

import (
	"database/sql"
	"fmt"

	"github.com/lokeshllkumar/auth-service/pkg/utils"
)

func (db *DB) AddUser(username string, password string) error {
	addCmd := `INSERT INTO users (username, password) VALUES (?, ?)`
	encPswd, err := utils.EncryptPassword(password)
	if err != nil {
		fmt.Printf("Error encrypting password\n")
		return err
	}
	_, err = db.Conn.Exec(addCmd, username, encPswd)
	if err != nil {
		fmt.Printf("Error adding entry to database\n")
		return err
	}

	return nil
}

func (db *DB) GetUser(username string) (User, error) {
	var user User
	getCmd := `SELECT username, password FROM users WHERE username = ?`
	rows, err := db.Conn.Query(getCmd, username)
	if err != nil {
		fmt.Printf("Error retrieving user information\n")
		panic(err)
	}
	err = rows.Scan(&user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, err
		}
		fmt.Printf("Error parsing retrieved data\n")
		panic(err)
	}

	return user, nil
}

func (db *DB) UpdateUser(username string, newPassword string) error {
	encPswd, err := utils.EncryptPassword(newPassword)
	if err != nil {
		fmt.Printf("Error encypting password\n")
		return err
	}

	updateCmd := `UPDATE users SET password = ? WHERE username = ?`
	_, err = db.Conn.Exec(updateCmd, encPswd, username)
	return err
}