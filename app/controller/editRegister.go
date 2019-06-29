package controller

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// EditRegister controller
func EditRegister(r *http.Request, w http.ResponseWriter, data *GeneralData, pageData *PageData) {
	data.Filename = "edit-register"
	queryValues := r.URL.Query()
	val, ok := queryValues["register"]

	if !ok {
		panic("missing register in url !")
	}

	registerID := val[0]
	register := models.NewRegister()
	register.ID = registerID
	register.GetByID()

	if register.Owner != data.UserID {
		panic("not allowed !! this is not your register!")
	}

	(*pageData)["amountCards"] = len(register.Cards)
	(*pageData)["cards"] = register.Cards
	(*pageData)["register"] = register
	(*pageData)["card"] = nil

	if len(register.Cards) > 0 {
		(*pageData)["card"] = register.Cards[0].ID
	}

	// if user wants a specific
	if val, ok := queryValues["card"]; ok {
		(*pageData)["card"] = val[0]
	}
}
