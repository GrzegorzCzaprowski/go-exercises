package models

import (
	"database/sql"

	"github.com/lib/pq"
)

type Todo struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type User struct {
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

type Model struct {
	DB *sql.DB
}

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

func (model Model) LogUser(user User) error {
	row := model.DB.QueryRow("SELECT id, email, password, created_at FROM users WHERE email=$1 and password=$2", user.Email, user.Password)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return err
	}
	return err
}

func (model Model) CreateUser(user User) error {
	_, err := model.DB.Exec("INSERT INTO users(email, password) VALUES($1, $2)", user.Email, user.Password)
	if err, ok := err.(*pq.Error); ok {
		if err.Code == "23505" {
			return err
		}
	}
	return err
}

func (model Model) CreateTodo(todo Todo) error {
	_, err := model.DB.Exec("INSERT INTO todos(name, description) VALUES($1, $2)", todo.Name, todo.Description)
	return err
}

func (model Model) ReadAllTodos() ([]Todo, error) {
	var todos []Todo
	rows, err := model.DB.Query("SELECT id, name, description, created_at, updated_at FROM todos")
	if err != nil {
		return todos, err
	}

	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}
	return todos, rows.Err()
}

func (model Model) ReadById(id int) (Todo, error) {
	todo := Todo{}
	row := model.DB.QueryRow("SELECT id, name, description, created_at, updated_at FROM todos WHERE id=$1", id)

	err := row.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
	if err != nil {
		return todo, err
	}
	return todo, err
}

func (model Model) UpdateById(todo Todo, id int) error {
	_, err := model.DB.Exec("UPDATE todos SET name=$1, description=$2, updated_at = now() WHERE id=$3", todo.Name, todo.Description, id)
	if err != nil {
		return err
	}
	return err
}

func (model Model) RemoveById(id int) error {
	_, err := model.DB.Exec("DELETE FROM todos WHERE id=$1", id)
	if err != nil {
		return err
	}
	return err
}
