package controller

import (
	"net/http"
)

// Register controller
func Register(r *http.Request, data *pageData) {
	data.Filename = "register"
}
