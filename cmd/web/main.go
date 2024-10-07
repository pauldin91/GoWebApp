package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pauldin91/GoWebApp/pkg/cfg"
	"github.com/pauldin91/GoWebApp/pkg/handlers"
	"github.com/pauldin91/GoWebApp/pkg/render"
)

const portNumber = ":8080"

func main() {

	var app cfg.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.Initialize(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Println("listening on http://localhost" + portNumber)
	srv.ListenAndServe()
}
