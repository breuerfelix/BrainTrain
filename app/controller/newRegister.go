package controller

import "net/http"

// NewRegister controller
func NewRegister(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "new-register"
}
