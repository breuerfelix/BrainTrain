package controller

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
	"github.com/gorilla/sessions"
)

// Categorie struct
type Categorie struct {
	Name          string
	SubCategories []string
}

// GeneralData struct
type GeneralData struct {
	UserID              string
	Username            string
	User                models.User
	Filename            string
	LoggedIn            bool
	NewPublicRegisters  int
	NewPrivateRegisters int
	ShowAnswer          bool
	Categories          []Categorie
	WrongPassword       bool
}

// PageData for templates
type PageData map[string]interface{}

type controllerFunction func(*http.Request, *GeneralData, *PageData)

var key []byte
var store *sessions.CookieStore

func init() {
	key = make([]byte, 32)
	rand.Read(key)
	store = sessions.NewCookieStore(key)
}

// HandleWithContext func
func HandleWithContext(controllerFunc controllerFunction, authRequired bool) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		data := GeneralData{}

		// Check if user is authenticated
		if authRequired {
			if auth, ok := session.Values["authenticated"].(bool); ok && auth {
				fmt.Println("logged in")
				data.LoggedIn = true
				data.UserID, _ = session.Values["userID"].(string)
				data.Username, _ = session.Values["userName"].(string)

				user := models.NewUser()
				user.ID = data.UserID
				user.GetByID()
				data.User = *user

			} else {
				fmt.Println("not logged in")
				http.Redirect(w, r, "/", http.StatusFound)
				return
			}
		}

		if wrong, ok := session.Values["wrongPassword"].(bool); ok && wrong {
			fmt.Println("wrong password")
			data.WrongPassword = true
			session.Values["wrongPassword"] = false
			session.Save(r, w)
		}

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
