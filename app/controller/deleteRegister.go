package controller

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// DeleteRegister controller
func DeleteRegister(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	queryValues := r.URL.Query()

	UserID, _ := session.Values["userID"].(string)

	user := models.NewUser()
	user.ID = UserID
	user.GetByID()

	val, ok := queryValues["register"]

	if !ok {
		panic("missing register in url !")
	}

	registerID := val[0]
	register := models.NewRegister()
	register.ID = registerID
	register.GetByID()

	if register.Owner == user.ID {
		allUser := user.GetAllUser()
		for _, u := range allUser {
			index := -1
			for i, p := range u.Progress {
				if p.Register == register.ID {
					index = i
					break
				}
			}

			if index >= 0 {
				if len(u.Progress) < 2 {
					u.Progress = make([]models.RegisterProgress, 0)
				} else {
					s := u.Progress
					i := index
					s[len(s)-1], s[i] = s[i], s[len(s)-1]
					u.Progress = s[:len(s)-1]
				}

				u.Update()
			}
		}

		register.Delete()
	} else {
		index := -1
		for i, p := range user.Progress {
			if p.Register == register.ID {
				index = i
				break
			}
		}

		if index >= 0 {
			if len(user.Progress) < 2 {
				user.Progress = make([]models.RegisterProgress, 0)
			} else {
				s := user.Progress
				i := index
				s[len(s)-1], s[i] = s[i], s[len(s)-1]
				user.Progress = s[:len(s)-1]
			}

			user.Update()
		}
	}

	http.Redirect(w, r, "/private-registers", http.StatusFound)
}
