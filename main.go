package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := executeTemplate(w, "home")
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error while executing the template", http.StatusInternalServerError)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	err := executeTemplate(w, "contact")
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error while executing the template", http.StatusInternalServerError)
	}
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	err := executeTemplate(w, "faq")
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error while executing the template", http.StatusInternalServerError)
	}
}

func executeTemplate(w http.ResponseWriter, filename string) error {
	// build path for any OS
	tplPath := filepath.Join("templates", filename+".gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		return err
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
