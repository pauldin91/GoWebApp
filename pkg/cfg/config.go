package cfg

import "text/template"

type AppConfig struct {
	TemplateCache map[string]*template.Template
}
