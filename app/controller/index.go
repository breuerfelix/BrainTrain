package controller

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

type statistics struct {
	Users          int
	CardsTotal     int
	RegistersTotal int
}

// Index controller
func Index(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "index"

	// TODO maybe use a sum aggregation ? 
	user := models.NewUser()
	allUser, _ := user.GetAll()

	card := models.NewCard()
	allCards, _ := card.GetAll()

	register := models.NewRegister()
	allRegisters, _ := register.GetAll()

	(*pageData)["stats"] = &statistics{
		Users:          len(allUser),
		CardsTotal:     len(allCards),
		RegistersTotal: len(allRegisters),
	}
}
