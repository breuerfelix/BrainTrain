package controller

import (
	"net/http"
)

type userStatistics struct {
	Cards     int
	Registers int
}

// Profile controller
func Profile(r *http.Request, w http.ResponseWriter, data *GeneralData, pageData *PageData) {
	data.Filename = "profile"
	(*pageData)["hasPicture"] = false

	if data.User.Picture != "" {
		(*pageData)["hasPicture"] = true
	}

	(*pageData)["stats"] = &userStatistics{
		Cards:     49,
		Registers: 9,
	}
}
