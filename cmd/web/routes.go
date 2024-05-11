package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tamagossi/room-bookings/pkg/config"
	"github.com/tamagossi/room-bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	/*
	 * Commented 4.39 to replace pat to use chi
	   mux := pat.New()

	   mux.Get("/", http.HandlerFunc(handlers.Repo.HomeHandler))
	   mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutHandler))
	*/

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(LoadSession)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
