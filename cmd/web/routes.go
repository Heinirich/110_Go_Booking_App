package main

import (
	"net/http"

	"github.com/Heinrich/110_Go_Booking_App/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes() http.Handler{

	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()
	
	mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	mux.Use(SessionLoad)
	mux.Use(NoSurf)

	mux.Get("/",handlers.Repo.Home)
	mux.Get("/about",handlers.Repo.About)
	
	return mux
}