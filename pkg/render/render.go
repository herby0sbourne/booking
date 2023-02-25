package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"myapp/pkg/config"
	"net/http"
	"path/filepath"
)

var function = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a

}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, html string) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		// get template cache from app config
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// get requested template from  cache
	template, ok := templateCache[html]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = template.Execute(buf, nil)

	//render template
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all teh files named *.page.html or *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html or *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}

		}

		myCache[name] = ts

	}
	return myCache, nil
}
