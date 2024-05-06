package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	InfoLog       *log.Logger
	TemplateCache map[string]*template.Template
	UseCache      bool
}
