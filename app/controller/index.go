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

type userResp struct {
	Users []struct {
		Email    string
		Password string
		Name     string
	}
}

// Index controller
func Index(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "index"

	user := &models.User{}
	user.Type = "user"
	allUser, _ := user.GetAll()

	card := &models.Card{}
	card.Type = "card"
	allCards, _ := card.GetAll()

	register := &models.Register{}
	register.Type = "register"
	allRegisters, _ := register.GetAll()

	(*pageData)["stats"] = &statistics{
		Users:          len(allUser),
		CardsTotal:     len(allCards),
		RegistersTotal: len(allRegisters),
	}
}
