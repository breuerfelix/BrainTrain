package routes

import (
	"net/http"

	"github.com/breuerfelix/card-learning-web/app/controller"
)

func Init() {
	http.HandleFunc("/", controller.Index)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
