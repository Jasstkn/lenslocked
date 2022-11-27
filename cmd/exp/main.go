package main

import (
	"html/template"
	"os"
)

func main() {
	type User struct {
		Name string
	}

	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := User{Name: "John Smith"}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
