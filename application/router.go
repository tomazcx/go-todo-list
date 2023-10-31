package application

import (
	"net/http"

	"github.com/tomazcx/go-todo-list/application/controllers"
	middlewares "github.com/tomazcx/go-todo-list/application/middlwares"
)

func Router() {

	taskController := controllers.TaskController{}
	authController := controllers.AuthController{}

	//TASK ROUTES
	http.HandleFunc("/", middlewares.UsesAuth(taskController.Index))
	http.HandleFunc("/createTask", middlewares.UsesAuth(taskController.CreateTask))
	http.HandleFunc("/toggleCompletedTask", middlewares.UsesAuth(taskController.ToggleCompleted))
	http.HandleFunc("/deleteTask", middlewares.UsesAuth(taskController.DeleteTask))

	//AUTH ROUTES
	http.HandleFunc("/register", authController.Register)
	http.HandleFunc("/login", authController.Login)

	http.HandleFunc("/auth/createAccount", authController.HandleRegister)
	http.HandleFunc("/auth/login", authController.HandleLogin)
	http.HandleFunc("/auth/logout", authController.HandleLogOut)
}
