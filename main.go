package main

import (
	"fmt"
	"net/http"

	"github.com/Jasstkn/lenslocked/controllers"
	"github.com/Jasstkn/lenslocked/templates"
	"github.com/Jasstkn/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	layoutTpl := "tailwind.gohtml"

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", layoutTpl)),
	))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", layoutTpl)),
	))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", layoutTpl,
	))
	r.Get("/signup", usersC.New)

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", layoutTpl)),
	))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on http://localhost:3000")

	http.ListenAndServe(":3000", r)
}
