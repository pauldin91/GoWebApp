package handlers

import (
	"net/http"

	"github.com/pauldin91/GoWebApp/pkg/cfg"
	"github.com/pauldin91/GoWebApp/pkg/models"
	"github.com/pauldin91/GoWebApp/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *cfg.AppConfig
}

func NewRepo(a *cfg.AppConfig) *Repository {

	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r

}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{}
	stringMap["test"] = "Hello again"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	
	stringMap["remote_ip"] = remoteIp
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: stringMap})
}
