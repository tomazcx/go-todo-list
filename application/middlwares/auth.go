package middlewares

import (
	"fmt"
	"net/http"

	"github.com/tomazcx/go-todo-list/application/controllers"
)

func UsesAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store := controllers.GetStoreSession()
		session, _ := store.Get(r, "todo-user")

		if authenticaded, ok := session.Values["auth"].(bool); !authenticaded && !ok {
			fmt.Println(authenticaded, ok)
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next(w, r)
	}
}
