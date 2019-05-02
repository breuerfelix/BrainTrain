package routes

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/controller"
)

// Init function
func Init() {
	http.HandleFunc("/", controller.HandleWithContext(controller.Index))
	http.HandleFunc("/public-registers", controller.HandleWithContext(controller.PublicRegisters))
	http.HandleFunc("/private-registers", controller.HandleWithContext(controller.PrivateRegisters))
	http.HandleFunc("/signup", controller.HandleWithContext(controller.Signup))

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
