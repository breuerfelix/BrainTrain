package controller

import "net/http"

// PublicRegisters controller
func PublicRegisters(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "public-registers"
}
