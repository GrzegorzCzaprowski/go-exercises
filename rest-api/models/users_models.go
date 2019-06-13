package models

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type UsersModel struct {
	DB *sql.DB
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (model UsersModel) LogUser(user User) (User, error) {
	password := user.Password
	row := model.DB.QueryRow("SELECT id, email, password, created_at FROM users WHERE email=$1", user.Email)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	if CheckPasswordHash(password, user.Password) {
		return user, err
	}
	return user, errors.New("incorect password")
}

func (model UsersModel) CreateUser(user User) (User, error) {
	var err error
	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return user, err
	}

	_, err = model.DB.Exec("INSERT INTO users(email, password) VALUES($1, $2)", user.Email, user.Password)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return user, err
		}
	}
	row := model.DB.QueryRow("SELECT id, email, password, created_at FROM users WHERE email=$1", user.Email)
	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	return user, err
}
