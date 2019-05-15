package controller

import "net/http"

// Signup controller
func Signup(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "signup"
}
