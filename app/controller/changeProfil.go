package controller

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// ChangeProfile updates profile details
func ChangeProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userName, _ := session.Values["userName"].(string)

	newEmail := r.FormValue("newEmail")
	password := r.FormValue("password")
	newPassword := r.FormValue("newPassword")
	newPasswordAgain := r.FormValue("newPasswordAgain")
	
	user := models.NewUser()
	user.Get("name", userName)

	if user.Password != password {
		session.Values["errorWrongPassword"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	if newEmail != "" {
		user.Email = newEmail
	}

	if newPassword != "" && newPassword == newPasswordAgain {
		user.Password = newPassword
	}

	if err:= user.Update(); err != nil {
		session.Values["errorDuringSave"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	session.Values["successUpdateProfile"] = true
	session.Save(r, w)
	http.Redirect(w, r, "/profile", http.StatusFound)
	return
}