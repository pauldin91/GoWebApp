package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pauldin91/GoWebApp/pkg/cfg"
	"github.com/pauldin91/GoWebApp/pkg/handlers"
	"github.com/pauldin91/GoWebApp/pkg/render"
)

const portNumber = ":8080"

var app cfg.AppConfig
var session *scs.SessionManager

func main() {

	tc, err := render.CreateTemplateCache()

	session = scs.New()
	session.Lifetime = time.Hour * 24
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Production
	app.Session = session

	
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
