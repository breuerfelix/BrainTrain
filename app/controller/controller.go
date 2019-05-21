package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// Categorie struct
type Categorie struct {
	Name          string
	SubCategories []string
}

// GeneralData struct
type GeneralData struct {
	models.User
	UserID              string
	Filename            string
	LoggedIn            bool
	NewPublicRegisters  int
	NewPrivateRegisters int
	ShowAnswer          bool
	Categories          []Categorie
}

// PageData for templates
type PageData map[string]interface{}

type controllerFunction func(*http.Request, *GeneralData, *PageData)

// HandleWithContext func
func HandleWithContext(controllerFunc controllerFunction) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := GeneralData{}
		pageData := PageData{
			"general": &data,
		}

		initGeneralData(&data)
		controllerFunc(r, &data, &pageData)

		tmpl, err := template.ParseFiles("templates/layout.tmpl", fmt.Sprintf("templates/%s.tmpl", data.Filename))

		if err != nil {
			log.Println(err)
		}

		tmpl.ExecuteTemplate(w, "layout", pageData)
	}
}

func initGeneralData(data *GeneralData) {
	// types
	// user, card, register

	// fake data
	data.LoggedIn = true
	data.NewPublicRegisters = 20
	data.NewPrivateRegisters = 0
	data.ShowAnswer = false

	// fake user data
	data.User.Name = "Max Mustermann"
	data.User.Date = "24.12.2018"
	data.UserID = "c2abd60d207e8b04baae4b2a50000801"

	// init categories
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
}
