package models

import (
	"github.com/go-rest-api/db"
	"github.com/go-rest-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// var users = []User{}

func (u *User) Save() (User, error) {
	stmt, err := db.DB.Prepare(`
		INSERT INTO users (email, password) VALUES (?, ?)
		`)
	if err != nil {
		return User{}, err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return User{}, err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return User{}, err
	}

	u.ID, err = result.LastInsertId()
	if err != nil {
		return User{}, err
	}

	return GetUserByEmail(u.Email)
}

func GetUserByEmail(email string) (User, error) {
	row := db.DB.QueryRow("SELECT id , email, password FROM users WHERE email = ?", email)

	var user User
	err := row.Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func VerifyCredentials(email, password string) (User, error) {
	user, err := GetUserByEmail(email)
	if err != nil {
		return User{}, err
	}

	err = utils.VerifyPassword(password, user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
