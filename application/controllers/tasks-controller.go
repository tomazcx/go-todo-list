package controllers

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"

	"github.com/tomazcx/go-todo-list/models"
)

func TodoListHome(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.Index()
	if err != nil {
		http.Error(w, "Internal server error: Error fetching the tasks from the database.", http.StatusInternalServerError)
		return
	}
	template.Must(template.ParseFiles("./templates/index.html", "./templates/partials/task.html")).Execute(w, tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Internal server error: Error parsing the form", http.StatusInternalServerError)
		return
	}

	taskName := r.FormValue("name")
	task, err := models.CreateTask(taskName)

	if err != nil {
		http.Error(w, "Internal server error: Error creating the task", http.StatusInternalServerError)
		return
	}

	template.Must(template.ParseFiles("./templates/partials/task.html")).Execute(w, task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idFromQuery := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idFromQuery)

	if err != nil {
		http.Error(w, "Error: invalid ID", http.StatusUnprocessableEntity)
	}

	err = models.DeleteTask(uint(id))

	if err != nil {
		http.Error(w, "Internal server error: Error deleting the task", http.StatusInternalServerError)
		return
	}

	tasks, err := models.Index()
	if err != nil {
		http.Error(w, "Internal server error: Error fetching the tasks from the database.", http.StatusInternalServerError)
		return
	}

	taskTmpl := template.Must(template.ParseFiles("./templates/partials/task.html"))

	var buf bytes.Buffer
	for _, task := range tasks {
		if err = taskTmpl.Execute(&buf, task); err != nil {
			http.Error(w, "Internal server error: Error rendering the task template.", http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(buf.String()))
}
