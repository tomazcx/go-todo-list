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
	http.HandleFunc("/createTask", taskController.CreateTask)
	http.HandleFunc("/toggleCompletedTask", taskController.ToggleCompleted)
	http.HandleFunc("/deleteTask", taskController.DeleteTask)

	//AUTH ROUTES
	http.HandleFunc("/register", authController.Register)
	http.HandleFunc("/login", authController.Login)

	http.HandleFunc("/auth/createAccount", authController.HandleRegister)
	http.HandleFunc("/auth/login", authController.HandleLogin)
	http.HandleFunc("/auth/logout", authController.HandleLogOut)
}
