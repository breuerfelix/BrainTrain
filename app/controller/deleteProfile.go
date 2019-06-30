package controller

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userName, _ := session.Values["userName"].(string)

	user := models.NewUser()
	user.Get("name", userName)

	if err := user.Delete(); err != nil {
		session.Values["errorDeleteProfil"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/profile", http.StatusFound)
		return
	}

	session.Values["successDeleteProfile"] = true

	session.Values["authenticated"] = false
	session.Values["userName"] = ""
	session.Values["userID"] = ""

	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
	return
}