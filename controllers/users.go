package controllers

import (
	"fmt"
	"net/http"

	"github.com/Jasstkn/lenslocked/models"
)

type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, r, data)
}

func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, r, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// TODO: replace with gorilla/schema https://github.com/gorilla/schema
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.UserService.Create(email, password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %v", user.Email)
}

func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.UserService.Authenticate(data.Email, data.Password)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "email",
		Value:    user.Email,
		Path:     "/",  // limit cookie to this path
		HttpOnly: true, // don't allow cookie to be accessible via JavaScript
	}
	http.SetCookie(w, &cookie)
	fmt.Fprintf(w, "User authenticated: %s", user.Email)
}

func (u Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	emailCookie, err := r.Cookie("email")
	if err != nil {
		http.Redirect(w, r, "http://localhost:3000/signin", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, "Email cookie: %s\n", emailCookie.Value)
	fmt.Fprintf(w, "Headers: %+v\n", r.Header)
}
