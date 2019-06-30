package controller

import (
	"fmt"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// CreateRegister creates a new register
func CreateRegister(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

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

	register := models.NewRegister()
	register.Owner = user.ID
	register.Description = description
	register.Private = private
	register.Title = title
	register.SubCategory = category
	register.Cards = make([]models.Card, 0)

	for _, cat := range categries {
		for _, subcat := range cat.SubCategories {
			if subcat == category {
				register.Category = cat.Name
				break
			}
		}
	}

	id := register.Insert()

	newProgress := models.RegisterProgress{Register: id}
	user.Progress = append(user.Progress, newProgress)
	user.Update()
	url := fmt.Sprintf("/edit-register?register=%s", id)
	http.Redirect(w, r, url, http.StatusFound)
}
