package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Vaibhav11022003/bookings/pkg/config"
	"github.com/Vaibhav11022003/bookings/pkg/models"
	"github.com/Vaibhav11022003/bookings/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// creates a new Repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// set the Repo  for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// handler functions
func (repo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	ipAddr := r.RemoteAddr
	repo.App.Session.Put(r.Context(), "remote_ip", ipAddr)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{StringMap: map[string]string{"test": "This is Dyamic data for Home Page"}})
}
func (repo *Repository) About(w http.ResponseWriter, r *http.Request) {
	ipAddr := repo.App.Session.GetString(r.Context(), "remote_ip")
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{StringMap: map[string]string{
		"test":   "This is Dyamic data for About Page",
		"ipAddr": ipAddr,
	}})
}

// func which uses a helper function
func (repo *Repository) Divide(w http.ResponseWriter, r *http.Request) {
	ans, err := divide(10, -1)
	if err != nil {
		fmt.Println(err)
	}
	size, err := fmt.Fprintln(w, "This is Divide Page and 10 divided by -1 is :", ans)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes written : %d/n", size)
}

// it is a helper function which is of type http.Handler acutally an internal type as this type has a listen function which we will understand later
func (repo *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	size, err := fmt.Fprintln(w, "This is Contact Page")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Number of bytes written : %d/n", size)
}

// helper functions
func divide(x, y float32) (float32, error) {
	if y <= 0 {
		return 0, errors.New("cannot divide by a number <= 0")
	}
	return x / y, nil
}
