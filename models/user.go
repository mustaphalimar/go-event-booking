package models

import "github.com/mustaphalimar/event-booking/db"

type User struct {
	ID       int64
	email    string `binding:"required"`
	password string `binding:"required"`
}

func (user User) Save() error {
	query := "INSERT INTO users(email,password) VALUES (?,?)"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.email, user.password)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	user.ID = userId

	return err
}
