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
	// build path for any OS
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error while parsing the template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("execute teamplate: %v", err)
		http.Error(w, "There was an error while executing the template", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact page</h1><p>To get in touch, email me at <a href=\"mailto:jasstkn.051@gmail.com\">jasstkn.051@gmail.com</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprint(w, `<h1> This is article #`+id+`</h1>`)
	fmt.Fprint(w, `
	<h1>FAQ Page</h1>
	<ul>
	<li>
		<b>Is there a free version?</b><br>
		Yes! We offer a free trial for 30 days on any paid plans.
	</li>
	<li>
		<b>What are your support hours?</b><br>
		We have support staff answering emails 24/7, though response
		times may be a bit slower on weekends.
	</li>
	<li>
		<b>How do I contact support?</b><br>
		Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
	</li>
	</ul>
	`)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq/{id}", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
