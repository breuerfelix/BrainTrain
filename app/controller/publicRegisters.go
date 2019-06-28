package controller

import (
	"fmt"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// PublicRegisters controller
func PublicRegisters(r *http.Request, w http.ResponseWriter, data *GeneralData, pageData *PageData) {
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

	// filter out categories which are not used
	categories := make([]string, 0)
	for _, cat := range data.Categories {
		found := false
		for _, reg := range allEntities {
			if reg["category"] == cat.Name {
				found = true
				break
			}
		}

		if found {
			categories = append(categories, cat.Name)
		}
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
	}

	// append to template data
	(*pageData)["registers"] = &allEntities
	(*pageData)["categories"] = &categories
}
