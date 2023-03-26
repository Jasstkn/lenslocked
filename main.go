package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/csrf"

	"github.com/Jasstkn/lenslocked/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Jasstkn/lenslocked/controllers"
	"github.com/Jasstkn/lenslocked/templates"
	"github.com/Jasstkn/lenslocked/views"
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

	cfg := models.DefaultPostgresConfig()
	db, err := sql.Open("pgx", cfg.String())
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	userService := models.UserService{DB: db}
	if err != nil {
		panic(err)
	}

	usersC := controllers.Users{
		UserService: &userService, // TODO: configure db connection
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", layoutTpl,
	))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", layoutTpl,
	))

	r.Get("/signup", usersC.New)
	r.Get("/signin", usersC.SignIn)
	r.Get("/users/me", usersC.CurrentUser)

	r.Post("/users", usersC.Create)
	r.Post("/signin", usersC.ProcessSignIn)

	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", layoutTpl)),
	))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on http://localhost:3000")

	csrfKey := os.Getenv("CSRF_KEY")
	csrfMiddleware := csrf.Protect(
		[]byte(csrfKey),
		// TODO: make configurable
		csrf.Secure(false))
	err = http.ListenAndServe(":3000", csrfMiddleware(r))
	if err != nil {
		panic(err)
	}
}
