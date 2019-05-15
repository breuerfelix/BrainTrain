package controller

import "net/http"

// Index controller
func Index(r *http.Request, data *GeneralData, pageData *PageData) {
	data.Filename = "index"
}
