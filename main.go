package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/csrf"

	"github.com/Jasstkn/lenslocked/migrations"
	"github.com/Jasstkn/lenslocked/models"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/Jasstkn/lenslocked/controllers"
	"github.com/Jasstkn/lenslocked/templates"
	"github.com/Jasstkn/lenslocked/views"
)

func main() {
	// Bootstrap the database
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// run migrations
	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	// Setup services
	userService := models.UserService{DB: db}
	sessionService := models.SessionService{DB: db}

	// Setup middleware
	umw := controllers.UserMiddleware{
		SessionService: &sessionService,
	}

	csrfKey := os.Getenv("CSRF_KEY")
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: make configurable
		csrf.Secure(false))

	// Setup controllers
	layoutTpl := "tailwind.gohtml"
	usersC := controllers.Users{
		UserService:    &userService,
		SessionService: &sessionService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(
		templates.FS,
		"signup.gohtml", layoutTpl,
	))
	usersC.Templates.SignIn = views.Must(views.ParseFS(
		templates.FS,
		"signin.gohtml", layoutTpl,
	))

	// Setup router and routes
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(csrfMw)
	r.Use(umw.SetUser)

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", layoutTpl)),
	))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", layoutTpl)),
	))

	r.Get("/signup", usersC.New)
	r.Get("/signin", usersC.SignIn)
	r.Get("/users/me", usersC.CurrentUser)
	r.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", layoutTpl)),
	))

	r.Post("/users", usersC.Create)
	r.Post("/signin", usersC.ProcessSignIn)
	r.Post("/signout", usersC.ProcessSignOut)

	r.Route("/users/me", func(r chi.Router) {
		r.Use(umw.RequireUser)
		r.Get("/", usersC.CurrentUser)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})

	// Start the server
	fmt.Println("Starting the server on http://localhost:3000")

	// csrfMiddleware runs first, then SetUser
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
