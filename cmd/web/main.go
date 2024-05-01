package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/tamagossi/room-booking/pkg/config"
	"github.com/tamagossi/room-booking/pkg/handlers"
	"github.com/tamagossi/room-booking/pkg/render"
)

var port = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour // The session will last for 24 hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tmpl, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tmpl
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.HomeHandler)
	// http.HandleFunc("/about", handlers.Repo.AboutHandler)
	// http.ListenAndServe(":8080", nil)

	fmt.Println("Server running on port: " + port)
	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
