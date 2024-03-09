package main

import (
	"net/http"
	"todo-list/pkg/db"
	"todo-list/pkg/handlers"
)

func main() {
	todoList := db.NewTodoList()

	taskHandler := handlers.NewHandler(todoList)

	http.HandleFunc("/tasks/", taskHandler.TasksEndpoint)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
