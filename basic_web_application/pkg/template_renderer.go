package pkg

import (
	"fmt"
	"html/template"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := getOrCreateCachedTemplate(tmpl)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func getOrCreateCachedTemplate(tmpl string) (*template.Template, error) {
	t, inCache := templateCache[tmpl]
	if inCache {
		return t, nil
	}

	files := []string{
		"./templates/" + tmpl,
		"./templates/base.layout.tmpl",
	}

	parsedTemplate, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	templateCache[tmpl] = parsedTemplate
	return parsedTemplate, nil
}
