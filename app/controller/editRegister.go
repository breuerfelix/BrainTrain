package controller

import (
	"fmt"
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

	if register.Owner == data.UserID {
		panic("not allowed !! this is not your register!")
	}

	allCards, _ := models.DB.QueryJSON(fmt.Sprintf(`{
		"selector": {
			"type": {
				"$eq": "card"
			},
			"register_id": {
				"$eq": "%s"
			}
		}
	}`, register.ID))

	(*pageData)["amountCards"] = len(allCards)
	(*pageData)["cards"] = allCards
	(*pageData)["register"] = register
	(*pageData)["card"] = nil
	// TODO calculate
	(*pageData)["progress"] = 30

	if len(allCards) > 0 {
		(*pageData)["card"] = allCards[0]
	}

	// if user wants a specific
	if val, ok := queryValues["card"]; ok {
		(*pageData)["card"] = val
	}
}
