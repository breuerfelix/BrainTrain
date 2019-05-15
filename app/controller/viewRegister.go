package controller

import "net/http"

// ViewRegister controller
func ViewRegister(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "view-register"
}
