package routes

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/controller"
)

// Init function
func Init() {
	// template routes
	http.HandleFunc("/", controller.HandleWithContext(controller.Index, false))
	http.HandleFunc("/public-registers", controller.HandleWithContext(controller.PublicRegisters, false))
	http.HandleFunc("/private-registers", controller.HandleWithContext(controller.PrivateRegisters, true))
	http.HandleFunc("/signup", controller.HandleWithContext(controller.Signup, false))
	http.HandleFunc("/new-register", controller.HandleWithContext(controller.NewRegister, true))
	http.HandleFunc("/edit-register", controller.HandleWithContext(controller.EditRegister, true))
	http.HandleFunc("/view-register", controller.HandleWithContext(controller.ViewRegister, false))
	http.HandleFunc("/learn-register", controller.HandleWithContext(controller.LearnRegister, true))
	http.HandleFunc("/profile", controller.HandleWithContext(controller.Profile, true))

	// functional routes
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/create-register", controller.CreateRegister)
	http.HandleFunc("/update-register", controller.UpdateRegister)
	http.HandleFunc("/new-card", controller.NewCard)
	http.HandleFunc("/edit-card", controller.EditCard)
	http.HandleFunc("/delete-card", controller.DeleteCard)

	// file server
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
