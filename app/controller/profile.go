package controller

import "net/http"

// Profile controller
func Profile(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "profile"
}
