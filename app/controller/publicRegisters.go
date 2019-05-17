package controller

import (
	"fmt"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// PublicRegisters controller
func PublicRegisters(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "public-registers"

	allEntities, err := models.DB.QueryJSON(`{
		"selector": {
			"type": {
				"$eq": "register"
			},
			"private": {
				"$eq": false
			}
		}
	}`)

	if err != nil {
		panic(err)
	}

	// TODO make list from categorie class and filter out not used ones
	categories := [...]string{
		"Naturwissenschaften",
		"Sprachen",
		"Gesellschaft",
		"Wirtschaft",
		"Geisteswissenschaften",
	}

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
	}

	(*pageData)["registers"] = &allEntities
	(*pageData)["categories"] = &categories
}
