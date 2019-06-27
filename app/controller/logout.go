package controller

import (
	"net/http"
)

// Logout current user
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	session.Values["authenticated"] = false
	session.Values["userName"] = nil
	session.Values["userID"] = nil
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)

}
