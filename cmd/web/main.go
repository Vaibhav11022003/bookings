package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Vaibhav11022003/bookings/pkg/config"
	"github.com/Vaibhav11022003/bookings/pkg/handlers"
	"github.com/Vaibhav11022003/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app = config.AppConfig{}
var session *scs.SessionManager // this one are saved in db or some type of file and diff for diff users

func main() {
	//create appConfig and set them

	app.InProd = false
	app.UseCache = false
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	// if false whenever browser quits or window quits get lost
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProd
	app.Session = session
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln(err, "Cannot create TemplateCache")
	}
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	fmt.Println("Server is listening on port :", portNumber)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
