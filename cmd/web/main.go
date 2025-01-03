package main

import (
	"encoding/gob"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/pauldin91/GoWebApp/internal/config"
	"github.com/pauldin91/GoWebApp/internal/driver"
	"github.com/pauldin91/GoWebApp/internal/handlers"
	"github.com/pauldin91/GoWebApp/internal/helpers"
	"github.com/pauldin91/GoWebApp/internal/models"
	"github.com/pauldin91/GoWebApp/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application function
func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	app.UseCache = false
	defer db.SQL.Close()

	defer close(app.MailChan)

	fmt.Println("Starting mail listener")

	listenForMails()
	/**/

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(map[string]int{})

	inProduction := flag.Bool("production", true, "Application environment")
	useCache := flag.Bool("cache", true, "Use template cache")
	dbhost := flag.String("dbhost", "", "Database host")
	dbport := flag.String("dbport", "", "Database port")
	dbname := flag.String("dbname", "", "Database name")
	dbuser := flag.String("dbuser", "", "Database user")
	dbpass := flag.String("dbpass", "", "Database password")

	flag.Parse()

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan
	// change this to true when in production
	app.InProduction = *inProduction
	app.UseCache = *useCache

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", *dbhost, *dbport, *dbname, *dbuser, *dbpass)

	db, err := driver.ConnectSQL(connectionString)

	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
