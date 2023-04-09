package main

import (
	"net/http"
	"github.com/justinas/nosurf"
)

// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(
// 		func(w http.ResponseWriter, r *http.Request) {
// 			fmt.Println("Hit the Page")
// 			next.ServeHTTP(w, r)
// 		})
// }

// Nosurf adds CSRF protection to all request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.Inproduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad Loads and Saves Session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}