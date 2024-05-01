package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tamagossi/room-booking/pkg/config"
	"github.com/tamagossi/room-booking/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.HomeHandler))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutHandler))

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSruve)
	mux.Use(LoadSession)

	mux.Get("/", handlers.Repo.HomeHandler)
	mux.Get("/about", handlers.Repo.AboutHandler)

	return mux
}
