package repositories

import (
	"github.com/tomazcx/go-todo-list/entities"
)

type TodoRepository struct {
	data []entities.Todo
}

func (r *TodoRepository) Find(id int) *entities.Todo {
	for _, todo := range r.data {
		if todo.ID == id {
			return &todo
		}
	}

	return nil
}

func (r *TodoRepository) Exists(id int) bool {
	found := false
	for _, todo := range r.data {
		if todo.ID == id {
			found = true
		}
	}

	return found
}

func (r *TodoRepository) GetAll() []entities.Todo {
	return r.data
}

func (r *TodoRepository) Create(title string) *entities.Todo {
	lastId := r.data[len(r.data)-1].ID
	newTodo := entities.Todo{ID: lastId + 1, Title: title, Completed: true}
	r.data = append(r.data, newTodo)
	return &newTodo
}

func (r *TodoRepository) Toggle(id int) {
	for _, todo := range r.data {
		crrTodo := &todo
		if crrTodo.ID == id {
			crrTodo.Completed = !crrTodo.Completed
		}
	}
}

func (r *TodoRepository) Delete(id int) {
	var newData []entities.Todo

	for _, todo := range r.data {
		if todo.ID != id {
			newData = append(newData, todo)
		}
	}

	r.data = newData
}
