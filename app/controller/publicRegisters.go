package controller

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// PublicRegisters controller
func PublicRegisters(r *http.Request, w http.ResponseWriter, data *GeneralData, pageData *PageData) {
	data.Filename = "public-registers"

	tempRegisters := models.NewRegister().GetAllRegister()
	allRegister := make([]models.Register, 0)

	for _, reg := range tempRegisters {
		if !reg.Private {
			allRegister = append(allRegister, reg)
		}
	}

	// filter out categories which are not used
	categories := make([]string, 0)
	for _, cat := range data.Categories {
		found := false
		for _, reg := range allRegister {
			if reg.Category == cat.Name {
				found = true
				break
			}
		}

		if found {
			categories = append(categories, cat.Name)
		}
	}

	// append number of cards for each register
	for _, register := range allRegister {
		register.Misc["amountCards"] = len(register.Cards)
	}

	// append to template data
	(*pageData)["registers"] = &allRegister
	(*pageData)["categories"] = &categories
}
