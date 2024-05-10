package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

/*
 * adds CSRF protetction to all POST requests
 */
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   app.InProduction,
	})

	return csrfHandler
}

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}
