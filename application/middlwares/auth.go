package middlewares

import (
	"net/http"

	"github.com/tomazcx/go-todo-list/application/utils"
)

func UsesAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := utils.GetStoreSession(r)

		if _, ok := session.Values["userId"]; !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		next(w, r)
	}
}
