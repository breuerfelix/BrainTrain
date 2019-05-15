package controller

import "net/http"

// LearnRegister controller
func LearnRegister(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "learn-register"
}
