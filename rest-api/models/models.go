package models

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	DB *sql.DB
}

type Todo struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

func (server *Server) CreateTodo(w http.ResponseWriter, req *http.Request) {
	todo := &Todo{}
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Println("blad w decodowaniu z jsona", err)
	}
	defer req.Body.Close()

	result, err := server.DB.Exec("INSERT INTO todos(name, description) VALUES($1, $2)", todo.Name, todo.Description)
	if err != nil {
		log.Println("ERROR saving to db - ", err)
	}

	Id64, err := result.LastInsertId()
	id := int(Id64)
	todo = &Todo{ID: id}

	log.Println("POST poszedl")
}

func (server *Server) ReadAllTodos(w http.ResponseWriter, req *http.Request) {
	rows, _ := server.DB.Query("SELECT id, name, description, created_at, updated_at FROM todos")
	var todos []Todo
	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			log.Println("nie tak sie skanuje", err)
		}
		todos = append(todos, todo)
	}
	json.NewEncoder(w).Encode(todos)

}
