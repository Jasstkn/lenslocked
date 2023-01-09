package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: replace with gorilla/schema https://github.com/gorilla/schema
	fmt.Fprint(w, "Email: ", r.FormValue("email"))
	fmt.Fprint(w, "Password: ", r.FormValue("password"))
}
