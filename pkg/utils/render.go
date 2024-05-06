package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var templateCache = make(map[string]*template.Template)

/*
	Used in 3.30
*/
func CreateTemplateCache(templateName string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", templateName),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	templateCache[templateName] = tmpl
	return nil
}

func CreateTemplateCacheEnhanced() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	/*
		Get all the files named *.page.tmpl in ./templates
	*/
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		filename := filepath.Base(page)
		template, err := template.New(filename).ParseFiles(page)
		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return cache, err
		}

		if len(layouts) > 0 {
			template, err = template.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}

		cache[filename] = template
	}

	return cache, nil
}

/*
	Used before 3.30
*/
func RenderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

/*
	Used in 3.30
*/
func RenderTemplateWithCaching(w http.ResponseWriter, templateName string) {
	var tmpl *template.Template
	var err error

	_, isCacheExist := templateCache[templateName]
	if !isCacheExist {
		err = CreateTemplateCache(templateName)
		if err != nil {
			log.Println(err)
		}
	} else {
		log.Println("using cache template")
	}

	tmpl = templateCache[templateName]
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func RenderTemplateWithCachingEnhanced(w http.ResponseWriter, templateName string) {
	/*
		Get template cache from app config
	*/

	templateCache, err := CreateTemplateCacheEnhanced()
	if err != nil {
		log.Fatal(err)
	}

	template, ok := templateCache[templateName]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	err = template.Execute(buf, nil)
	if !ok {
		log.Fatal(err)
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}
