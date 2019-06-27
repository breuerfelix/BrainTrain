package controller

import (
	"fmt"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// LearnRegister controller
func LearnRegister(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "learn-register"
	queryValues := r.URL.Query()
	val, ok := queryValues["register"]

	if !ok {
		panic("missing register in url !")
	}

	registerID := val[0]
	register := models.NewRegister()
	register.ID = registerID
	register.GetByID()

	/* if register.Owner != data.UserID {
		panic("not allowed !! this is not your register!")
	} */

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

	cardsZero := 0
	cardsOne := 0
	cardsTwo := 0
	cardsThree := 0
	cardsFour := 0

	for _, card := range allCards {
		switch card["progress"] {
		case 0.0:
			cardsZero++
		case 1.0:
			cardsOne++
		case 2.0:
			cardsTwo++
		case 3.0:
			cardsThree++
		case 4.0:
			cardsFour++
		}
	}

	(*pageData)["cardsZero"] = cardsZero
	(*pageData)["cardsOne"] = cardsOne
	(*pageData)["cardsTwo"] = cardsTwo
	(*pageData)["cardsThree"] = cardsThree
	(*pageData)["cardsFour"] = cardsFour

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
