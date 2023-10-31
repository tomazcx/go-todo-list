package utils

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func GetStoreSession(r *http.Request) *sessions.Session {
	session, _ := store.Get(r, "todo-user")
	return session
}

func SetUserSession(w http.ResponseWriter, r *http.Request, userId uint) error {
	session, _ := store.Get(r, "todo-user")
	session.Values["userId"] = userId
	session.Options.MaxAge = 60 * 60 * 24
	err := session.Save(r, w)
	return err
}

func ExpireSession(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "todo-user")
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	return err
}
