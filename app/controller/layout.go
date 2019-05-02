package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type pageData struct {
	Filename string
}

type controllerFunction func(*http.Request, *pageData)

// HandleWithContext func
func HandleWithContext(controllerFunc controllerFunction) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := pageData{}
		controllerFunc(r, &data)

		tmpl, err := template.ParseFiles("templates/layout.tmpl", fmt.Sprintf("templates/%s.tmpl", data.Filename))

		if err != nil {
			log.Fatal(err)
		}

		tmpl.ExecuteTemplate(w, "layout", data)
	}
}
