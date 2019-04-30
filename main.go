package main

import (
	"net/http"

	"github.com/breuerfelix/card-learning-web/app/routes"
)

func main() {
	routes.Init()

	server := http.Server{
		Addr: ":8080",
	}

	server.ListenAndServe()
}
