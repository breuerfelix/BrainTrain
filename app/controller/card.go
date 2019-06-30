package controller

import (
	"fmt"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
	"github.com/rs/xid"
)

// NewCard create a new card for a given register
func NewCard(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	queryValues := r.URL.Query()

	UserID, _ := session.Values["userID"].(string)

	user := models.NewUser()
	user.ID = UserID
	user.GetByID()

	val, ok := queryValues["register"]

	if !ok {
		panic("missing register in url !")
	}

	registerID := val[0]
	register := models.NewRegister()
	register.ID = registerID
	register.GetByID()

	id := xid.New()
	newCard := models.Card{ID: id.String()}
	register.Cards = append(register.Cards, newCard)
	register.Update()

	url := fmt.Sprintf("/edit-register?register=%s&card=%s", register.ID, id)
	http.Redirect(w, r, url, http.StatusFound)
}

// EditCard edits an existing card
func EditCard(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	queryValues := r.URL.Query()

	UserID, _ := session.Values["userID"].(string)

	user := models.NewUser()
	user.ID = UserID
	user.GetByID()

	val, ok := queryValues["register"]

	if !ok {
		panic("missing register in url !")
	}

	cardval, cardok := queryValues["card"]

	if !cardok {
		panic("missing card in url !")
	}

	cardid := cardval[0]

	registerID := val[0]
	register := models.NewRegister()
	register.ID = registerID
	register.GetByID()

	title := r.FormValue("title")
	question := r.FormValue("question")
	answer := r.FormValue("answer")

	index := 0
	for i, card := range register.Cards {
		if card.ID == cardid {
			index = i
			break
		}
	}

	oldCard := register.Cards[index]
	newCard := models.Card{
		ID:       oldCard.ID,
		Title:    title,
		Answer:   answer,
		Question: question,
	}
	register.Cards[index] = newCard
	register.Update()

	url := fmt.Sprintf("/edit-register?register=%s&card=%s", register.ID, cardid)
	http.Redirect(w, r, url, http.StatusFound)
}

// DeleteCard deletes a card
func DeleteCard(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	queryValues := r.URL.Query()

	UserID, _ := session.Values["userID"].(string)

	user := models.NewUser()
	user.ID = UserID
	user.GetByID()

	val, ok := queryValues["register"]

	if !ok {
		panic("missing register in url !")
	}

	registerID := val[0]
	register := models.NewRegister()
	register.ID = registerID
	register.GetByID()

	cardval, cardok := queryValues["card"]

	if !cardok {
		panic("missing card in url !")
	}

	s := register.Cards
	i := 0

	for index, card := range register.Cards {
		if card.ID == cardval[0] {
			i = index
			break
		}
	}

	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	register.Cards = s[:len(s)-1]
	register.Update()

	url := fmt.Sprintf("/edit-register?register=%s", register.ID)
	http.Redirect(w, r, url, http.StatusFound)
}
