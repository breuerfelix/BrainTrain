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

	// needs to query "user hase used but is public somehow" too
	allPrivateRegisters, err := models.DB.QueryJSON(fmt.Sprintf(`{
		"selector": {
			"type": {
				"$eq": "register"
			},
			"user": {
				"$eq": "%s"
			}
		}
	}`, data.Name))

	print(len(allPrivateRegisters))

	if err != nil {
		panic(err)
	}

	for _, register := range allPrivateRegisters {
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
	(*pageData)["registers"] = &allPrivateRegisters
	(*pageData)["currentUserID"] = userID
}
