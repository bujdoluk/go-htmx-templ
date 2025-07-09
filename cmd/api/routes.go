package main

import (
	"expvar"
	"fmt"
	"html"
	"net/http"
	"regexp"

	"github.com/bujdoluk/go-htmx-templ/views"
	"github.com/julienschmidt/httprouter"
)

func (app *application) helloHandler(w http.ResponseWriter, r *http.Request) {
	views.Hello().Render(r.Context(), w)
}

func (app *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		views.Login().Render(r.Context(), w)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@example\.com$`, email)
	if !matched {
		fmt.Fprint(w, `<p style="color:red;">❌ Email must be from @example.com</p>`)
		return
	}

	passValid, _ := regexp.MatchString(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`, password)
	if !passValid {
		fmt.Fprint(w, `<p style="color:red;">❌ Password must be at least 8 characters long and include a letter and a number</p>`)
		return
	}

	if email == "user@example.com" && password == "password123" {
		fmt.Fprint(w, `<p style="color:green;">✅ Login successful!</p>`)
	} else {
		fmt.Fprint(w, `<p style="color:red;">❌ Invalid credentials.</p>`)
	}
}

func (app *application) signUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		views.SignUp().Render(r.Context(), w)
		return
	}

	// POST: handle form submission
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	matched, _ := regexp.MatchString(`^[a-zA-Z0-9._%+\-]+@example\.com$`, email)
	if !matched {
		fmt.Fprint(w, `<p style="color:red;">❌ Email must be from @example.com</p>`)
		return
	}

	passValid, _ := regexp.MatchString(`^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`, password)
	if !passValid {
		fmt.Fprint(w, `<p style="color:red;">❌ Password must be at least 8 characters long and include a letter and a number</p>`)
		return
	}

	// Fake sign-up logic: Just respond for now
	fmt.Fprintf(w, `<p style="color:green;">✅ Account created successfully for %s</p>`, html.EscapeString(email))
}

func (app *application) redirectToLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)

	router.HandlerFunc(http.MethodGet, "/", app.redirectToLoginHandler)

	router.HandlerFunc(http.MethodGet, "/signup", app.signUpHandler)
	router.HandlerFunc(http.MethodPost, "/signup", app.signUpHandler)

	router.HandlerFunc(http.MethodGet, "/login", app.loginHandler)
	router.HandlerFunc(http.MethodPost, "/login", app.loginHandler)

	router.HandlerFunc(http.MethodGet, "/hello", app.helloHandler)

	router.ServeFiles("/static/*filepath", http.Dir("static"))

	router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	return router
}
