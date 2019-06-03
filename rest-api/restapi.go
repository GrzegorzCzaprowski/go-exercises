//https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b
//https://golang.org/doc/effective_go.html

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Todo struct {
	Id          int    `json:"id,omitempty`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Created_at  string `json:"created_at,omitempty"`
	Updated_at  string `json:"updated_at,omitempty"`
}

func (server *Server) CreateTodo(w http.ResponseWriter, req *http.Request) {
	todo := &Todo{}
	err := json.NewDecoder(req.Body).Decode(&todo)
	if err != nil {
		log.Println("blad w decodowaniu z jsona", err)
	}
	defer req.Body.Close()

	result, err := server.db.Exec("INSERT INTO todos(name, description) VALUES($1, $2)", todo.Name, todo.Description)
	if err != nil {
		log.Println("ERROR saving to db - ", err)
	}

	Id64, err := result.LastInsertId()
	id := int(Id64)
	todo = &Todo{Id: id}

	log.Println("POST poszedl")
}

func (server *Server) ReadAllTodos(w http.ResponseWriter, req *http.Request) {
	rows, _ := server.db.Query("SELECT id, name, description, created_at, updated_at FROM todos")
	var todos []Todo
	for rows.Next() {
		todo := Todo{}
		err := rows.Scan(&todo.Id, &todo.Name, &todo.Description, &todo.Created_at, &todo.Updated_at)
		if err != nil {
			log.Println("nie tak sie skanuje", err)
		}
		todos = append(todos, todo)
	}
	json.NewEncoder(w).Encode(todos)

}

type Server struct {
	db *sql.DB
}

func main() {

	db, err := sql.Open("postgres", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := &Server{db: db}

	router := mux.NewRouter()

	router.HandleFunc("/api/todos/", server.CreateTodo).Methods("POST")
	router.HandleFunc("/api/todos/", server.ReadAllTodos).Methods("GET")
	// router.HandleFunc("/api/todos/:id/", ReadByID).Methods("GET")
	// router.HandleFunc("/api/todos/:id/", UpdateTodo).Methods("PATCH")
	// router.HandleFunc("/api/todos/:id/", RemoveTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
