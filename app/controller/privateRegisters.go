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

	// user has to learn his own registers at creation time
	register := models.NewRegister()
	allRegisters := register.GetAllRegister()

	userRegisters := make([]models.Register, 0)

	for _, register := range allRegisters {
		found := false
		for _, userReg := range data.User.Progress {
			if userReg.Register == register.ID {
				found = true
				break
			}
		}

		if found {
			register.Misc["CardCount"] = len(register.Cards)
			userRegisters = append(userRegisters, register)
		}
	}

	fmt.Println(userRegisters)

	// append to template data
	(*pageData)["registers"] = &userRegisters
}
