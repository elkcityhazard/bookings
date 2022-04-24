package handlers

import (
	"net/http"

	"github.com/elkcityhazard/bookings/pkg/config"
	"github.com/elkcityhazard/bookings/pkg/models"
	"github.com/elkcityhazard/bookings/pkg/render"
)

//	Repository is the repository type

type Repository struct {
	App *config.AppConfig
}

//	Repo the repository used by the handlers

var Repo *Repository

//	NewRepo creates a new repository

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//	NewHandlers sets the repository for the handlers

func NewHandlers(r *Repository) {
	Repo = r
}

// Home is a handler function and it handles a response writer and a request

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

//	About is the about page handler it handles a response writer and request

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//	perform some business logic

	stringMap := make(map[string]string)

	stringMap["test"] = "Hello Again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	stringMap["remote_ip"] = remoteIP

	//	send data the to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
