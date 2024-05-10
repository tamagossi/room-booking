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
	remoteIP := repo.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := make(map[string]string)
	stringMap["test"] = "Hallo, this is test"
	stringMap["remote_ip"] = remoteIP

	utils.RenderTemplateWithCachingEnhanced(w, "about.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) HomeHandler(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	utils.RenderTemplateWithCachingEnhanced(w, "home.page.tmpl", &models.TemplateData{})
}
