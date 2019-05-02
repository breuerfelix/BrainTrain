package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type pageData struct {
	Filename            string
	LoggedIn            bool
	Username            string
	NewPublicRegisters  int
	NewPrivateRegisters int
}

type controllerFunction func(*http.Request, *pageData)

// HandleWithContext func
func HandleWithContext(controllerFunc controllerFunction) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := pageData{}

		// fake data
		data.LoggedIn = true
		data.Username = "Max Mustermann"
		data.NewPublicRegisters = 20
		data.NewPrivateRegisters = 0

		controllerFunc(r, &data)

		tmpl, err := template.ParseFiles("templates/layout.tmpl", fmt.Sprintf("templates/%s.tmpl", data.Filename))

		if err != nil {
			log.Println(err)
		}

		tmpl.ExecuteTemplate(w, "layout", data)
	}
}
