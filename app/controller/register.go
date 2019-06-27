package controller

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// Register
func Register(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	userName := r.FormValue("username")
	userEmail := r.FormValue("userEmail")
	password := r.FormValue("password")
	passwordAgain := r.FormValue("passwordAgain")
	agreement := r.FormValue("agreement")

	user := models.NewUser()
	user.Get("name", userName)

	if user.ID != nil {
		session.Values["errorUserExistsAlready"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	user.Name = userName
	user.Email = userEmail
	user.Password = password

	if user.Password == password {
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
