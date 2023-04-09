package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Heinrich/110_Go_Booking_App/pkg/config"
	"github.com/Heinrich/110_Go_Booking_App/pkg/handlers"
	"github.com/Heinrich/110_Go_Booking_App/pkg/render"
	"github.com/alexedwards/scs/v2"
)

// app is of type config.Appconfig
var app config.AppConfig
const PORT string = ":8080"
var session *scs.SessionManager


// Main Function for Serving my Server
func main()  {

	// Change to true in poduction 
	app.Inproduction = false

	session = scs.New()
	session.Lifetime = 24*time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.Inproduction

	app.Session = session



	tc,err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot Create template Cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplate(&app)


	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting Application on http://localhost"+PORT)
	// _ = http.ListenAndServe(PORT,nil)

	srv := &http.Server{
		Addr:              PORT,
		Handler:           routes(),
	}

	err = srv.ListenAndServe()
	
	log.Fatal(err)


}