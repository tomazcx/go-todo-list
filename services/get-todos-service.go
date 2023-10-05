package services

import (
	"github.com/tomazcx/go-todo-list/entities"
	"github.com/tomazcx/go-todo-list/repositories"
)

func GetTodosService(repository repositories.TodoRepository) []entities.Todo {
	return repository.GetAll()
}
