package handlers

import (
	"net/http"

	"github.com/tamagossi/room-bookings/pkg/config"
	"github.com/tamagossi/room-bookings/pkg/utils"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewHandler(repo *Repository) {
	Repo = repo
}

func NewRepo(repo *config.AppConfig) *Repository {
	return &Repository{App: repo}
}

func (repo *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplateWithCaching(w, "about.page.tmpl")
}

func (repo *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplateWithCaching(w, "home.page.tmpl")
}
