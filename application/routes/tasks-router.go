package routes

import (
	"net/http"

	"github.com/tomazcx/go-todo-list/application/controllers"
)

func TaskRoutes() {
	http.HandleFunc("/", controllers.TodoListHome)
	http.HandleFunc("/createTask", controllers.CreateTask)
	http.HandleFunc("/deleteTask", controllers.DeleteTask)
}
