package controller

import (
	"fmt"
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
	bla, _ := user.GetAll()

	fmt.Println(bla)

	(*pageData)["stats"] = &statistics{
		Users:          32,
		CardsTotal:     124,
		RegistersTotal: 22,
	}
}
