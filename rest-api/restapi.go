package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	todos "github.com/GrzegorzCzaprowski/go-exercises/rest-api/handlers/todos"
	users "github.com/GrzegorzCzaprowski/go-exercises/rest-api/handlers/users"
	"github.com/GrzegorzCzaprowski/go-exercises/rest-api/models"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

var flagDatabaseAddress string

func init() {
	flag.StringVar(&flagDatabaseAddress, "db", "postgres://testuser:testpass@localhost:5555/testdb?sslmode=disable", "database address")
	flag.Parse()
}

func main() {
	db, err := sql.Open("postgres", flagDatabaseAddress)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := httprouter.New()

	todosMod := models.TodosModel{DB: db}
	todoHandler := todos.TodoHandler{M: todosMod}
	router.POST("/api/todos/", todoHandler.Post)
	router.GET("/api/todos/", todoHandler.GetAll)
	router.GET("/api/todos/:id/", todoHandler.Get)
	router.PATCH("/api/todos/:id/", todoHandler.Patch)
	router.DELETE("/api/todos/:id/", todoHandler.Delete)

	userMod := models.UsersModel{DB: db}
	userHandler := users.UserHandler{M: userMod}
	router.POST("/api/user/create/", userHandler.PostUser)
	router.POST("/api/user/log", userHandler.LogUser)

	log.Fatal(http.ListenAndServe(":8000", router))
}
