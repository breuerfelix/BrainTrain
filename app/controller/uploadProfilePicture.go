package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/models"
)

// UploadProfilePicture saves pic to DB
func UploadProfilePicture(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userName, _ := session.Values["userName"].(string)

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic("error reading body")
	}

	// fmt.Println(string(body))

	user := models.NewUser()
	user.Get("name", userName)

	user.Picture = string(body)

	user.Update()

	http.Redirect(w, r, "/profile", http.StatusFound)
	return
}
