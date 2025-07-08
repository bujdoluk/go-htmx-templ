package main

import (
	"net/http"

	"github.com/bujdoluk/go-htmx-templ/views"
	"github.com/julienschmidt/httprouter"
)

func (app *application) mainHandler(w http.ResponseWriter, r *http.Request) {
	views.Index().Render(r.Context(), w)
}

func (app *application) helloHandler(w http.ResponseWriter, r *http.Request) {
	views.Hello().Render(r.Context(), w)
}

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.mainHandler)
	router.HandlerFunc(http.MethodGet, "/hello", app.helloHandler)

	router.ServeFiles("/static/*filepath", http.Dir("static"))

	return router
}
