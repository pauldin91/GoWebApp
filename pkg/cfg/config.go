package cfg

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCache  map[string]*template.Template
	Production     bool
	Session *scs.SessionManager
}
