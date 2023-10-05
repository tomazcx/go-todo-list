package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Todo struct {
	ID        int
	Name      string
	Completed bool
}

var todoList []Todo

func getList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./src/templates/list.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, todoList)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	var id int

	if len(todoList) == 0 {
		id = 0
	} else {
		id = todoList[len(todoList)-1].ID
	}

	todo := Todo{id + 1, name, false}

	todoList = append(todoList, todo)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func toggleTodo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(r.FormValue("id"))

	if err != nil {
		http.Error(w, "Invalid ID type", http.StatusUnprocessableEntity)
		return
	}

	for i := range todoList {
		todo := &todoList[i]
		if todo.ID == id {
			todo.Completed = !todo.Completed
			break
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	idFromRequest := r.FormValue("id")

	idToDelete, err := strconv.Atoi(idFromRequest)

	if err != nil {
		http.Error(w, "Invalid ID type", http.StatusUnprocessableEntity)
		return
	}

	var newList []Todo

	for _, todo := range todoList {
		if todo.ID != idToDelete {
			newList = append(newList, todo)
		}
	}

	todoList = newList

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func main() {
	fmt.Println("Server is starting...")

	http.HandleFunc("/", getList)
	http.HandleFunc("/create-todo", createTodo)
	http.HandleFunc("/toggle-todo", toggleTodo)
	http.HandleFunc("/delete-todo", deleteTodo)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	fmt.Println("Server is now running at port " + port + " 🚀")
	http.ListenAndServe(":"+port, nil)
}
