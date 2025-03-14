package models

import (
	"errors"
	"fmt"

	"github.com/mustaphalimar/event-booking/db"
	"golang.org/x/crypto/bcrypt"

	"github.com/mustaphalimar/event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users(email,password) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println("Failed to prepare the query")
		return err
	}
	defer stmt.Close()

	// hashing the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	user.ID = userId

	return err
}

func (user User) ValidateCredentials() error {
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return errors.New("Invalid credentials.")
	}

	isPasswordMatch := ComparePasswordHash(user.Password, retrievedPassword)

	if !isPasswordMatch {
		return errors.New("Invalid credentials.")
	}

	return nil
}

func ComparePasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
