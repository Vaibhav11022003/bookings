package main

import (
	"net/http"

	"github.com/Vaibhav11022003/bookings/pkg/config"
	"github.com/Vaibhav11022003/bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Use(middleware.Recoverer)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/divide", handlers.Repo.Divide)
	mux.Get("/contact", handlers.Repo.Contact)
	return mux
}
