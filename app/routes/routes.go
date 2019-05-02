package routes

import (
	"net/http"

	"github.com/breuerfelix/BrainTrain/app/controller"
)

// Init function
func Init() {
	http.HandleFunc("/", controller.HandleWithContext(controller.Register))

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
