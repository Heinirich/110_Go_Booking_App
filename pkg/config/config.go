package config

import (
	"html/template"

	"github.com/alexedwards/scs/v2"
)

// Appconfig holds the application configurations such as TemplateCache
type AppConfig struct {
	UseCache bool
	TemplateCache map[string] *template.Template
	Inproduction bool
	Session *scs.SessionManager
}