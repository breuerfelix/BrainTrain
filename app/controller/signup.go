package controller

import "net/http"

// Signup controller
func Signup(r *http.Request, w http.ResponseWriter, data *GeneralData, pageData *PageData) {
	data.Filename = "signup"
}
