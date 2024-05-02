package utils

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templateCache = make(map[string]*template.Template)

func createTemplateCache(templateName string) error {
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

func RenderTemplate(w http.ResponseWriter, templateName string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+templateName, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

func RenderTemplateWithCaching(w http.ResponseWriter, templateName string) {
	var tmpl *template.Template
	var err error

	_, isCacheExist := templateCache[templateName]
	if !isCacheExist {
		err = createTemplateCache(templateName)
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
