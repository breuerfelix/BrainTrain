package controller

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// ViewRegister controller
func ViewRegister(r *http.Request, w http.ResponseWriter, data *GeneralData, pageData *PageData) {
	data.Filename = "view-register"
	queryValues := r.URL.Query()
	val, ok := queryValues["register"]

	if !ok {
		panic("missing register in url !")
	}

	registerID := val[0]
	register := models.NewRegister()
	register.ID = registerID
	register.GetByID()

	(*pageData)["amountCards"] = len(register.Cards)
	(*pageData)["cards"] = register.Cards
	(*pageData)["register"] = register
	(*pageData)["card"] = nil
	(*pageData)["progress"] = 0

	if len(register.Cards) > 0 {
		(*pageData)["card"] = register.Cards[0].ID
	}

	// if user wants a specific
	if val, ok := queryValues["card"]; ok {
		(*pageData)["card"] = val[0]
	}
}
