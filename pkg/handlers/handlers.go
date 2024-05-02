package handlers

import (
	"net/http"

	"github.com/tamagossi/room-bookings/pkg/utils"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplateWithCaching(w, "about.page.tmpl")
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplateWithCaching(w, "home.page.tmpl")
}
