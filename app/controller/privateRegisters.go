package controller

import (
	"fmt"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// PrivateRegisters controller
func PrivateRegisters(r *http.Request, w http.ResponseWriter, data *GeneralData, pageData *PageData) {
	data.Filename = "private-registers"

	fmt.Println(data.UserID)

	// needs to query "user hase used but is public somehow" too
	register := models.NewRegister()
	allRegisters := register.GetAllRegister()

	userRegisters := make([]models.Register, 0)

	for _, register := range allRegisters {
		if register.Owner == data.UserID {
			register.Misc["CardCount"] = len(register.Cards)
			userRegisters = append(userRegisters, register)
		}
	}

	fmt.Println(userRegisters)

	// append to template data
	(*pageData)["registers"] = &userRegisters
}
