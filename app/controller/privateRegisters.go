package controller

import (
	"fmt"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// PrivateRegisters controller
func PrivateRegisters(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "private-registers"

	// TODO remove, hard coded user id
	// get from jwt via general data
	userID := "c2abd60d207e8b04baae4b2a50000801"

	allEntities, err := models.DB.QueryJSON(`{
		"selector": {
			"type": {
				"$eq": "register"
			},
			"private": {
				"$eq": true
			}
		}
	}`)

	if err != nil {
		panic(err)
	}

	// append number of cards for each register
	for _, register := range allEntities {
		allCards, _ := models.DB.QueryJSON(fmt.Sprintf(`{
			"selector": {
				"type": {
					"$eq": "card"
				},
				"register_id": {
					"$eq": "%s"
				}
			}
		}`, register["_id"]))

		register["amountCards"] = len(allCards)

		// TODO calculate
		register["progress"] = "30%"
	}

	// append to template data
	(*pageData)["registers"] = &allEntities
	(*pageData)["currentUserID"] = userID
}
