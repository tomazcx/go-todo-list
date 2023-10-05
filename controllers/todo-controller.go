package controllers

import (
	"html/template"
	"net/http"
)

func TodoPageController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	template, err := template.ParseFiles("../templates/list.html")

	if err != nil {
		http.Error(w, "Error parsing the template", http.StatusInternalServerError)
		return
	}

	template.Execute(w, nil)
}
