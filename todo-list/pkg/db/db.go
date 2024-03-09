package db

/*
	A good practice when handling shared data -like the items map in our db.go file- is to use mutual exclusion (mutex) to prevent race conditions. In the Go language, the sync package exposes a Mutex type for safe, concurrent operations.
*/

import (
	"errors"
	"sync"
)

// TodoItem represents a single item in the todo list
type TodoItem struct {
	ID          string
	Task        string
	IsCompleted bool
}

// TodoList represents the collection of all items in the todo list
type TodoList struct {
	// We use a mutex to ensure the TodoList is safe to use concurrently
	mu    sync.Mutex
	items map[string]TodoItem // we use a map for 0(1)  lookups
}

// NewTodoList initializes a new todo List
func NewTodoList() *TodoList {
	return &TodoList{
		items: make(map[string]TodoItem),
	}
}

// AddItem adds a new item to the todo list
func (t *TodoList) AddItem(item TodoItem) {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.items[item.ID] = item
}

// GetItem retrieves an item from the todo list by Its ID
func (t *TodoList) GetItem(id string) (TodoItem, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	item, ok := t.items[id]
	if !ok {
		return TodoItem{}, errors.New("item not found")
	}

	return item, nil
}

// DeleteItem removes an item from the todo list by Its ID
func (t *TodoList) DeleteItem(id string) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	if _, ok := t.items[id]; ok {
		delete(t.items, id)
		return nil
	}

	return errors.New("item not found")
}

// UpdateItem updates an existing item in the todo list
func (t *TodoList) UpdateItem(updated TodoItem) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	_, ok := t.items[updated.ID]
	if !ok {
		return errors.New("item not found")
	}

	t.items[updated.ID] = updated
	return nil
}

func (t *TodoList) GetItems() ([]TodoItem, error) {
	t.mu.Lock()
	defer t.mu.Unlock()

	var items []TodoItem
	for _, v := range t.items {
		items = append(items, v)
	}

	if len(items) != 0 {
		return items, nil
	}

	return nil, errors.New("no items found")
}
