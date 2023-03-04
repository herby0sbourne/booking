package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/herby0sbourne/booking/internal/config"
	"github.com/herby0sbourne/booking/internal/forms"
	"github.com/herby0sbourne/booking/internal/models"
	"github.com/herby0sbourne/booking/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remoteIp", remoteIp)

	render.RenderTemplate(w, r, "home.page.html", &models.TemplateData{})
}

// About is the home page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	remoteIp := m.App.Session.GetString(r.Context(), "remoteIp")

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"
	stringMap["ip"] = remoteIp

	render.RenderTemplate(w, r, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Generals renders generals quarters page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.page.html", &models.TemplateData{})
}

// Majors renders majors suite page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.page.html", &models.TemplateData{})
}

// Reservation render maker reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}

// PostReservation handle make reservation form submit
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Panicln(err)
		return
	}

	reservation := models.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Phone:     r.Form.Get("phone"),
		Email:     r.Form.Get("email"),
	}

	form := forms.New(r.PostForm)

	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3, r)

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

}

// Availability render the search availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.page.html", &models.TemplateData{})
}

// PostAvailability render the search availability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.FormValue("start")
	end := r.FormValue("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

// AvailabilityJSON handle request for availability and send json request
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Availability",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.page.html", &models.TemplateData{})
}
