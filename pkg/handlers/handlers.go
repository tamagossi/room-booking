package handlers

import (
	"net/http"

	"github.com/tamagossi/room-bookings/pkg/config"
	"github.com/tamagossi/room-bookings/pkg/models"
	"github.com/tamagossi/room-bookings/pkg/utils"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewHandler(repo *Repository) {
	Repo = repo
}

func NewRepo(app *config.AppConfig) *Repository {
	return &Repository{App: app}
}

func (repo *Repository) AboutHandler(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hallo, this is test"

	utils.RenderTemplateWithCachingEnhanced(w, "about.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	utils.RenderTemplateWithCachingEnhanced(w, "home.page.tmpl", &models.TemplateData{})
}
