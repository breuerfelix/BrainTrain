package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// User struct
type User struct {
	NumCards     int
	NumRegisters int
	MemberSince  string
	Username     string
}

type pageData struct {
	Filename            string
	LoggedIn            bool
	NewPublicRegisters  int
	NewPrivateRegisters int
	NumUsers            int
	NumCardsTotal       int
	NumRegistersTotal   int
	User
}

type controllerFunction func(*http.Request, *pageData)

// HandleWithContext func
func HandleWithContext(controllerFunc controllerFunction) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := pageData{}

		// fake data
		data.LoggedIn = true
		data.NewPublicRegisters = 20
		data.NewPrivateRegisters = 0
		data.NumUsers = 32
		data.NumCardsTotal = 124
		data.NumRegistersTotal = 22

		// fake user data
		data.User.Username = "Max Mustermann"
		data.User.MemberSince = "24.12.2018"
		data.User.NumCards = 24
		data.User.NumRegisters = 7

		controllerFunc(r, &data)

		tmpl, err := template.ParseFiles("templates/layout.tmpl", fmt.Sprintf("templates/%s.tmpl", data.Filename))

		if err != nil {
			log.Println(err)
		}

		tmpl.ExecuteTemplate(w, "layout", data)
	}
}
