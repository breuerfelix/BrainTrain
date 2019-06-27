package controller

import (
	"fmt"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// Register new user
func Register(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	userName := r.FormValue("userName")
	userEmail := r.FormValue("userEmail")
	password := r.FormValue("password")
	passwordAgain := r.FormValue("passwordAgain")
	// agreement := r.FormValue("agreement") // is ok if "on"

	user := models.NewUser()
	user.Get("name", userName)

	fmt.Println(user)

	if userName == "" {
		session.Values["errorNameNotSet"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	if userEmail == "" {
		session.Values["errorEmailNotSet"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	if password != passwordAgain {
		session.Values["errorPasswordNotSame"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	if user.ID != "" {
		session.Values["errorUserExistsAlready"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	user.Name = userName
	user.Email = userEmail
	user.Password = password

	if err := user.Insert(); err != nil {
		session.Values["errorDuringSave"] = true
		session.Save(r, w)
		http.Redirect(w, r, "/signup", http.StatusFound)
		return
	}

	session.Values["authenticated"] = true
	session.Values["userName"] = user.Name
	session.Values["userID"] = user.ID
	session.Save(r, w)
	http.Redirect(w, r, "/profile", http.StatusFound)
	return
}
