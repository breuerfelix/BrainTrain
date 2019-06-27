package controller

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// Login a user
func Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	userName := r.FormValue("username")
	password := r.FormValue("password")

	user := models.NewUser()
	user.Get("name", userName)

	if user.ID != "" && user.Password == password {
		session.Values["authenticated"] = true
		session.Values["userName"] = user.Name
		session.Values["userID"] = user.ID
		session.Save(r, w)
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	session.Values["wrongPassword"] = true
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}
