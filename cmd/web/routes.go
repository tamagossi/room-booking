package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/tamagossi/room-bookings/pkg/config"
	"github.com/tamagossi/room-bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.HomeHandler))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.AboutHandler))

	return mux
}
