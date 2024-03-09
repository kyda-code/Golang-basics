package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"todo-list/pkg/db"
)

type Handler struct {
	TodoList *db.TodoList
}

func NewHandler(list *db.TodoList) *Handler {
	return &Handler{TodoList: list}
}
func (h *Handler) TasksEndpoint(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "")
	id := strings.Trim(path, "/tasks/")

	switch {
	case path == "/tasks/":
		switch r.Method {
		case http.MethodPost:
			h.handleCreateTask(w, r)
		case http.MethodGet:
			h.handleGetTasks(w)
		case http.MethodPut:
			h.handleUpdateTask(w, r)
		default:
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		}
	case id != "":
		switch r.Method {
		case http.MethodDelete:
			h.handleDeleteTask(w, r, id)
		case http.MethodGet:
			h.handleGetTaskById(w, id)
		default:
			http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		}
	}
}

func (h *Handler) handleGetTasks(w http.ResponseWriter) {
	tasks, err := h.TodoList.GetItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(tasks)
	w.Header().Set("Content-Type", "application/json")
	_, err2 := w.Write(data)
	if err2 != nil {
		return
	}
}

func (h *Handler) handleGetTaskById(w http.ResponseWriter, id string) {
	tasks, err := h.TodoList.GetItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, _ := json.Marshal(tasks)
	w.Header().Set("Content-Type", "application/json")
	_, err2 := w.Write(data)
	if err2 != nil {
		return
	}
}

func (h *Handler) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	var todo db.TodoItem
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.TodoList.AddItem(todo)
	_, err = w.Write([]byte("Item added successfully"))
	if err != nil {
		return
	}
}

func (h *Handler) handleUpdateTask(w http.ResponseWriter, r *http.Request) {
	var todo db.TodoItem
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.TodoList.UpdateItem(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte("Item updated successfully"))
	if err != nil {
		return
	}
}

func (h *Handler) handleDeleteTask(w http.ResponseWriter, r *http.Request, i string) {
	id := i
	if id == "" {
		http.Error(w, "Invalid request, id required", http.StatusBadRequest)
		return
	}

	err := h.TodoList.DeleteItem(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte("Item deleted successfully"))
	if err != nil {
		return
	}
}
