package controller

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// NewRegister controller
func NewRegister(r *http.Request, w http.ResponseWriter, data *GeneralData, pageData *PageData) {
	data.Filename = "new-register"
	queryValues := r.URL.Query()
	val, ok := queryValues["register"]

	if ok {
		registerID := val[0]
		register := models.NewRegister()
		register.ID = registerID
		register.GetByID()

		(*pageData)["register"] = register
	}
}
