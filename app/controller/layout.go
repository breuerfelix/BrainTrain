package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Categorie struct {
	Name          string
	SubCategories []string
}

// User struct
type User struct {
	NumCards     int
	NumRegisters int
	MemberSince  string
	Username     string
}

type pageData struct {
	User
	Filename            string
	LoggedIn            bool
	NewPublicRegisters  int
	NewPrivateRegisters int
	ShowAnswer          bool
	NumUsers            int
	NumCardsTotal       int
	NumRegistersTotal   int
	Categories          []Categorie
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
		data.ShowAnswer = false
		data.NumUsers = 32
		data.NumCardsTotal = 124
		data.NumRegistersTotal = 22

		// fake user data
		data.User.Username = "Max Mustermann"
		data.User.MemberSince = "24.12.2018"
		data.User.NumCards = 24
		data.User.NumRegisters = 7

		// nature categorie fake
		nature := Categorie{}
		nature.Name = "Naturwissenschaften"
		nature.SubCategories = []string{
			"Biologie",
			"Chemie",
			"Elektrotechnik",
			"Informatik",
			"Mathematik",
			"Medizin",
			"Naturkunde",
			"Physik",
			"Sonstiges",
		}

		// languages categorie fake
		languages := Categorie{}
		languages.Name = "Sprachen"
		languages.SubCategories = []string{
			"Chinesisch",
			"Deutsch",
			"Englisch",
			"Fransösisch",
			"Griechisch",
			"Italienisch",
			"Latein",
			"Russisch",
			"Sonstiges",
		}

		// languages categorie fake
		society := Categorie{}
		society.Name = "Gesellschaft"
		society.SubCategories = []string{
			"Ethik",
			"Geschichte",
			"Literatur",
			"Musik",
			"Politik",
			"Recht",
			"Soziales",
			"Sport",
			"Verkehrskunde",
			"Sonstiges",
		}

		economy := Categorie{}
		economy.Name = "Wirtschaft"
		economy.SubCategories = []string{
			"BWL",
			"Finanzen",
			"Landwirtschaft",
			"Marketing",
			"VWL",
			"Sonstiges",
		}

		mindstuff := Categorie{}
		mindstuff.Name = "Geisteswissenschaften"
		mindstuff.SubCategories = []string{
			"Kriminalogie",
			"Philosophie",
			"Psychologie",
			"Pädagogik",
			"Theologie",
			"Sonstiges",
		}

		data.Categories = []Categorie{
			nature,
			languages,
			society,
			economy,
			mindstuff,
		}

		controllerFunc(r, &data)

		tmpl, err := template.ParseFiles("templates/layout.tmpl", fmt.Sprintf("templates/%s.tmpl", data.Filename))

		if err != nil {
			log.Println(err)
		}

		tmpl.ExecuteTemplate(w, "layout", data)
	}
}
