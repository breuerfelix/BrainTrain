package controller

import (
	"net/http"
)

// Index controller
func Index(r *http.Request, data *pageData) {
	data.Filename = "index"
}

// PublicRegisters controller
func PublicRegisters(r *http.Request, data *pageData) {
	data.Filename = "public-registers"
}

func Signup(r *http.Request, data *pageData) {
	data.Filename = "signup"
}
