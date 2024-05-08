package views

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/Jasstkn/lenslocked/context"
	"github.com/Jasstkn/lenslocked/models"
	"github.com/gorilla/csrf"
)

type public interface {
	Public() string
}

type Template struct {
	htmlTpl *template.Template
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0])

	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField not implemeneted") //nolint:goerr113
			},
			"currentUser": func() (template.HTML, error) {
				return "", fmt.Errorf("currentUser not implemeneted") //nolint:goerr113
			},
			"alerts": func() []string {
				return nil
			},
		},
	)

	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsefs template: %w", err)
	}

	return Template{htmlTpl: tpl}, nil
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}, errs ...error) {
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
	alerts := alertsMsgs(errs...)
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
			"currentUser": func() *models.User {
				return context.User(r.Context())
			},
			"alerts": func() []string {
				return alerts
			},
		})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var buffer bytes.Buffer
	err = tpl.Execute(&buffer, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buffer) //nolint:errcheck
}

func alertsMsgs(errs ...error) []string {
	var alerts []string
	for _, e := range errs {
		var pubErr public
		if errors.As(e, &pubErr) {
			alerts = append(alerts, pubErr.Public())
		} else {
			fmt.Println("error: ", e)
			alerts = append(alerts, "Something went wrong.")
		}
	}
	return alerts
}
