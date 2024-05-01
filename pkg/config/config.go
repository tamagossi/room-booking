package config

import (
	"log"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
	TemplateCache map[string]*template.Template
	UseCache      bool
}
