package controller

import (
	"fmt"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// UpdateRegister updated an existing register
func UpdateRegister(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	queryValues := r.URL.Query()
	val, ok := queryValues["register"]

	if !ok {
		panic("missing register in url!")
	}

	registerID := val[0]
	register := models.NewRegister()
	register.ID = registerID
	register.GetByID()

	UserID, _ := session.Values["userID"].(string)

	user := models.NewUser()
	user.ID = UserID
	user.GetByID()

	title := r.FormValue("title")
	category := r.FormValue("category")
	description := r.FormValue("description")
	tempprivate := r.FormValue("private")
	private := false
	if tempprivate == "on" {
		private = true
	}

	if private != register.Private {
		allUser := user.GetAllUser()
		for _, u := range allUser {
			// skip your own user
			if u.ID == user.ID {
				continue
			}

			index := -1
			for i, p := range u.Progress {
				if p.Register == register.ID {
					index = i
					break
				}
			}

			if index >= 0 {
				s := u.Progress
				i := index
				s[len(s)-1], s[i] = s[i], s[len(s)-1]
				u.Progress = s[:len(s)-1]
				u.Update()
			}
		}
	}

	register.Description = description
	register.Private = private
	register.Title = title
	register.SubCategory = category

	for _, cat := range categries {
		for _, subcat := range cat.SubCategories {
			if subcat == category {
				register.Category = cat.Name
				break
			}
		}
	}

	register.Update()

	url := fmt.Sprintf("/edit-register?register=%s", register.ID)
	http.Redirect(w, r, url, http.StatusFound)
}
