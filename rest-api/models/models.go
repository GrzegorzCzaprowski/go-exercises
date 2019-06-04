package models

import (
	"database/sql"
)

type Todo struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

type Model struct {
	db *sql.DB
}

func New(db *sql.DB) Model {
	return Model{db: db}
}

func (model Model) CreateTodo(todo Todo) error {
	result, err := model.db.Exec("INSERT INTO todos(name, description) VALUES($1, $2)", todo.Name, todo.Description)
	if err != nil {
		return err
	}

	Id64, err := result.LastInsertId()
	id := int(Id64)
	todo = Todo{ID: id}
	return err
}

func (model Model) ReadAllTodos() ([]Todo, error) {
	var todos []Todo
	rows, err := model.db.Query("SELECT id, name, description, created_at, updated_at FROM todos")
	if err != nil {
		return todos, err
	}

	for rows.Next() {
		todo := &Todo{}
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			return todos, err
		}
		todos = append(todos, *todo)
	}
	return todos, err
}
