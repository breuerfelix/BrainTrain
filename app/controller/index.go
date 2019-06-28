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
func Index(r *http.Request, w http.ResponseWriter, data *GeneralData, pageData *PageData) {
	data.Filename = "index"

	user := models.NewUser()
	allUser := user.GetAllUser()

	register := models.NewRegister()
	allRegisters := register.GetAllRegister()

	countCards := 0
	for _, register := range allRegisters {
		countCards += len(register.Cards)
	}

	(*pageData)["stats"] = &statistics{
		Users:          len(allUser),
		CardsTotal:     countCards,
		RegistersTotal: len(allRegisters),
	}
}
