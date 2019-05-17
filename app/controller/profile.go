package controller

import "net/http"

type userStatistics struct {
	Cards     int
	Registers int
}

// Profile controller
func Profile(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "profile"

	(*pageData)["stats"] = &userStatistics{
		Cards:     49,
		Registers: 9,
	}
}
