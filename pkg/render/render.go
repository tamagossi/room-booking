package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/tamagossi/room-booking/pkg/config"
	"github.com/tamagossi/room-booking/pkg/models"
)

var tc = make(map[string]*template.Template)
var app *config.AppConfig

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from the template cache")
	}

	buf := new(bytes.Buffer)
	data = AddDefaultData(data)
	_ = t.Execute(buf, data)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser: ", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

func RenderTemplateTest(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		log.Println("creating template and adding to the cache...")

		err = createTemplateCacheTest(t)
		if err != nil {
			log.Println("error parsing template: ", err)
		}
	} else {
		log.Println("using cached template...")
	}

	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
}

func createTemplateCacheTest(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	tc[t] = tmpl

	return nil
}
