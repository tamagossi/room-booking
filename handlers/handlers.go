package handlers

import (
	"net/http"

	"github.com/tamagossi/room-bookings/utils"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "about.page.tmpl")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplate(w, "home.page.tmpl")
}
