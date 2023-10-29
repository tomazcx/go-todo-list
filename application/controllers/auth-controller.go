package controllers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/tomazcx/go-todo-list/models"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct{}

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func (ac *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("./templates/auth/register.html")).Execute(w, nil)
}

func (ac *AuthController) HandleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not alloed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Internal server error: Error parsing the form", http.StatusInternalServerError)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirmPassword")

	if password != confirmPassword {
		w.Write([]byte("Passwords do not match"))
		return
	}

	accountModel := models.Account{}
	_, err = accountModel.FindByLogin(login)

	if err == nil {
		w.Write([]byte("Login already registered"))
		return
	}

	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		http.Error(w, "Internal server error: Error hashing the password", http.StatusInternalServerError)
		return
	}

	_, err = accountModel.Create(login, string(passwordBytes))

	if err != nil {
		http.Error(w, "Internal server error: Error creating the account", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/login")
	w.Write([]byte(""))
}

func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("./templates/auth/login.html")).Execute(w, nil)
}

func (ac *AuthController) HandleLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not alloed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Internal server error: Error parsing the form", http.StatusInternalServerError)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	accountModel := models.Account{}
	account, err := accountModel.FindByLogin(login)

	if err != nil {
		w.Write([]byte("Invalid credentials"))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))

	if err != nil {
		w.Write([]byte("Invalid credentials"))
		return
	}

	session, _ := store.Get(r, "todo-user")
	session.Values["auth"] = true

	err = session.Save(r, w)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", "/")
	w.Write([]byte(""))
}

func GetStoreSession() *sessions.CookieStore {
	return store
}
