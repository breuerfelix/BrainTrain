package controller

import (
	"fmt"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// PrivateRegisters controller
func PrivateRegisters(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "private-registers"

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
	}`, data.UserID))

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
}
