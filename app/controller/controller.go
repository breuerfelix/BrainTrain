package controller

import (
	"html/template"
	"log"
	"net/http"
)

type LayoutPageData struct {
	Filename string
}

func Index(w http.ResponseWriter, r *http.Request) {
	log.Print("hallo world")
	//tmpl := template.Must(template.ParseFiles("templates/layout.tmpl"))
	tmpl, err := template.ParseFiles("templates/layout.tmpl", "templates/register.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	data := LayoutPageData{
		Filename: "profil",
	}
	tmpl.Execute(w, data)
	//fmt.Fprintf(w, "Welcome to my website! yo retard")
}
