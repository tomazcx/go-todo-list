package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/tomazcx/go-todo-list/application/utils"
	"github.com/tomazcx/go-todo-list/models"
)

type TaskController struct{}

func (tc *TaskController) Index(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		template.Must(template.ParseFiles("./templates/not-found.html")).Execute(w, nil)
		return
	}
	session := utils.GetStoreSession(r)
	userId, ok := session.Values["userId"].(uint)

	fmt.Println(userId)

	if !ok {
		http.Error(w, "Not alloed", http.StatusForbidden)
		return
	}

	taskModel := models.Task{}
	tasks, err := taskModel.Index(userId)

	if !ok {
		http.Error(w, "Not alloed", http.StatusForbidden)
		return
	}

	if err != nil {
		http.Error(w, "Internal server error: Error fetching the tasks from the database.", http.StatusInternalServerError)
		return
	}
	template.Must(template.ParseFiles("./templates/index.html", "./templates/partials/task.html")).Execute(w, tasks)
}

func (tc *TaskController) CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Internal server error: Error parsing the form", http.StatusInternalServerError)
		return
	}
	taskModel := models.Task{}
	session := utils.GetStoreSession(r)
	userId, ok := session.Values["userId"].(uint)

	if !ok {
		http.Error(w, "Not alloed", http.StatusForbidden)
		return
	}

	taskName := r.FormValue("name")
	task, err := taskModel.Create(taskName, userId)

	if err != nil {
		http.Error(w, "Internal server error: Error creating the task", http.StatusInternalServerError)
		return
	}

	template.Must(template.ParseFiles("./templates/partials/task.html")).Execute(w, task)
}

func (tc *TaskController) ToggleCompleted(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPatch {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idFromQuery := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromQuery)

	if err != nil {
		http.Error(w, "Error: invalid ID", http.StatusUnprocessableEntity)
	}

	taskModel := models.Task{}
	task, err := taskModel.ToggleCompleted(uint(id))

	if err != nil {
		http.Error(w, "Internal server error: Error updating the task", http.StatusInternalServerError)
		return
	}

	template.Must(template.ParseFiles("./templates/partials/task.html")).Execute(w, task)
}

func (tc *TaskController) DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idFromQuery := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromQuery)

	if err != nil {
		http.Error(w, "Error: invalid ID", http.StatusUnprocessableEntity)
	}

	taskModel := models.Task{}
	err = taskModel.Delete(uint(id))

	if err != nil {
		http.Error(w, "Internal server error: Error deleting the task", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(""))
}
