package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing write to console middleware start")
		next.ServeHTTP(w, r)
		fmt.Println("Executing write to console middleware end")
	})
}

/**
* Add CSRF to all request
 */
func NoSruve(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   app.InProduction,
	})

	return csrfHandler
}

func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
