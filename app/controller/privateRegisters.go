package controller

import "net/http"

// PrivateRegisters controller
func PrivateRegisters(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "private-registers"
}
