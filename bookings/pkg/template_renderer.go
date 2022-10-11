package pkg

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

const templateDir = "./ui/templates/"

type TemplateData map[string]interface{}

type TemplateRenderer struct {
	appConfig AppConfig
}

func NewTemplateRenderer(appConfig AppConfig) TemplateRenderer {
	return TemplateRenderer{appConfig: appConfig}
}

var templateCache = make(map[string]*template.Template)

func (tr TemplateRenderer) RenderTemplate(w http.ResponseWriter, tmpl string, data *TemplateData) {
	t, err := tr.getOrCreateCachedTemplate(tmpl)
	if err != nil {
		fmt.Println(err)
		return
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, data)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (tr TemplateRenderer) getOrCreateCachedTemplate(tmpl string) (*template.Template, error) {
	t, inCache := templateCache[tmpl]
	if inCache && tr.appConfig.IsUseCache {
		fmt.Println("Template in cache:", tmpl)
		return t, nil
	}

	layouts, err := filepath.Glob(templateDir + "*.layout.tmpl")

	files := []string{
		templateDir + tmpl,
	}

	files = append(files, layouts...)

	parsedTemplate, err := template.ParseFiles(files...)
	if err != nil {
		return nil, err
	}

	templateCache[tmpl] = parsedTemplate
	fmt.Println("Template generated:", tmpl)
	return parsedTemplate, nil
}
