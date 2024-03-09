package controllers

import (
	"fmt"
	"net/http"
)

const (
	CookieSession = "session"
)

// newCookie is a helper function to create new cookie
// it takes name and value as arguments
// return pointer to the http.Cookie
func newCookie(name, value string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
}

// setCookie is a function that creates and sets cookie
// it takes http.ResponseWriter, and name and value for cookie
func setCookie(w http.ResponseWriter, name, value string) {
	cookie := newCookie(name, value)
	fmt.Printf("new cookie: %s", value)
	http.SetCookie(w, cookie)
}

// readCookie is a function that return cookie value
// it takes *http.Request and name of the cookie
// returns cookie's value and error
func readCookie(r *http.Request, name string) (string, error) {
	c, err := r.Cookie(name)
	if err != nil {
		return "", fmt.Errorf("failed to read cookie %s: %w", name, err)
	}

	return c.Value, nil
}

// deleteCookie is a function that deletes cookie
// it takes http.ResponseWriter and name of the cookie
func deleteCookie(w http.ResponseWriter, name string) {
	cookie := newCookie(name, "")
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}
