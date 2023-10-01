package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Todo struct {
	ID        int
	Name      string
	Completed bool
}

type TodoList struct {
	Data []Todo
}

var todoList TodoList

func getList(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/list.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, todoList.Data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	var id int

	if len(todoList.Data) == 0 {
		id = 0
	} else {
		id = todoList.Data[len(todoList.Data)-1].ID
	}

	todo := Todo{id + 1, name, false}

	todoList.Data = append(todoList.Data, todo)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
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
	}

	todoList.Data = append(todoList.Data[:int(idToDelete)-1], todoList.Data[idToDelete:]...)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func main() {
	fmt.Println("Server is starting...")

	http.HandleFunc("/", getList)
	http.HandleFunc("/create-todo", createTodo)
	http.HandleFunc("/delete-todo", deleteTodo)

	fmt.Println("Server is now running at port 8000 🚀")
	http.ListenAndServe(":8000", nil)
}
