package models

import (
	"errors"

	"github.com/extimsu/JsonRestApi/db"
	"github.com/extimsu/JsonRestApi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return errors.New("error preparing statement " + err.Error())
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return errors.New("error executing statement " + err.Error())
	}

	u.ID, err = result.LastInsertId()
	if err != nil {
		return errors.New("error with result ID " + err.Error())
	}

	return nil
}

func GetAllUsers() ([]User, error) {
	var (
		users []User
		user  User
	)
	query := `SELECT * FROM users`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, errors.New("cannot query users " + err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Password); err != nil {
			return nil, errors.New("cannot Scan user " + err.Error())
		}

		users = append(users, user)
	}
	return users, nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	if err := row.Scan(&u.ID, &retrivedPassword); err != nil {
		return errors.New("cannot find the user " + err.Error())
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)
	if !passwordIsValid {
		return errors.New("password is invalid! ")
	}

	return nil
}
