package handlers

import (
	"net/http"

	"github.com/Heinrich/110_Go_Booking_App/pkg/config"
	"github.com/Heinrich/110_Go_Booking_App/pkg/models"
	"github.com/Heinrich/110_Go_Booking_App/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is a repository type
type Repository struct {
	App *config.AppConfig

}
// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository)  {
	Repo = r
}

// Home returns Home page
func(m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr

	m.App.Session.Put(r.Context(),"remote_ip",remoteIp)
	

	
	render.RenderTemplate(w,"home.page.tmpl", &models.TemplateData{})
}

// About returns About Page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "About Page"

	remoteIp := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_ip"] = remoteIp
	
	render.RenderTemplate(w,"about.page.tmpl",&models.TemplateData{
		StringMap: stringMap,
	})
}

