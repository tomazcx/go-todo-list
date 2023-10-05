package services

import "github.com/tomazcx/go-todo-list/repositories"

func CreateTodoService(title string, repository *repositories.TodoRepository) {
	repository.Create(title)
}
