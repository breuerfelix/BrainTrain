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

// PrivateRegisters controller
func PrivateRegisters(r *http.Request, data *pageData) {
	data.Filename = "private-registers"
}

// Signup controller
func Signup(r *http.Request, data *pageData) {
	data.Filename = "signup"
}

// NewRegister controller
func NewRegister(r *http.Request, data *pageData) {
	data.Filename = "new-register"
}

// EditRegister controller
func EditRegister(r *http.Request, data *pageData) {
	data.Filename = "edit-register"
}

// ViewRegister controller
func ViewRegister(r *http.Request, data *pageData) {
	data.Filename = "view-register"
}

// Profile controller
func Profile(r *http.Request, data *pageData) {
	data.Filename = "profile"
}
