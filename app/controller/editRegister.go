package controller

import "net/http"

// EditRegister controller
func EditRegister(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "edit-register"
}
