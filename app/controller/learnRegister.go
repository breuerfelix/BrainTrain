package controller

import (
	"math/rand"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// LearnRegister controller
func LearnRegister(r *http.Request, w http.ResponseWriter, data *GeneralData, pageData *PageData) {
	data.Filename = "learn-register"
	// get register
	queryValues := r.URL.Query()
	val, ok := queryValues["register"]

	if !ok {
		panic("missing register in url !")
	}

	registerID := val[0]
	register := models.NewRegister()
	register.ID = registerID
	register.GetByID()

	if register.Private && register.Owner != data.UserID {
		http.Redirect(w, r, "/public-registers", http.StatusFound)
	}

	found := false
	for _, prog := range data.User.Progress {
		if prog.Register == register.ID {
			found = true
		}
	}

	if !found {
		// save to DB
		newProgress := models.RegisterProgress{}
		newProgress.Register = register.ID
		data.User.Progress = append(data.User.Progress, newProgress)
		data.User.Update()
	}

	allCards := register.Cards

	(*pageData)["register"] = register
	(*pageData)["amountCards"] = len(allCards)

	if len(allCards) > 0 {
		(*pageData)["card"] = allCards[rand.Intn(len(allCards))]
	}
}
