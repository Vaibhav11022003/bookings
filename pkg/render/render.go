package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/Vaibhav11022003/bookings/pkg/config"
	"github.com/Vaibhav11022003/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmplFileName string, td *models.TemplateData) {
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}
	parsedTemplate, flag := templateCache[tmplFileName]
	if !flag {
		log.Fatalf("%s not present", tmplFileName)
	}
	buffer := new(bytes.Buffer)
	td = AddDefaultData(td)
	err := parsedTemplate.Execute(buffer, td)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	var TemplateCache = make(map[string]*template.Template)
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}
	layouts, err := filepath.Glob("./templates/*.layout.tmpl")
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		pageName := filepath.Base(page)
		parsedTemplate := template.New(pageName)
		parsedTemplate, err := parsedTemplate.ParseFiles(append(layouts, page)...)
		if err != nil {
			return nil, err
		}
		TemplateCache[pageName] = parsedTemplate
	}
	return TemplateCache, nil
}

// cache for templates for fast access
var TemplateCache map[string]*template.Template = make(map[string]*template.Template)

// function that render a template without using cache but directly
func RenderTemplateWIthoutUsingCache(w http.ResponseWriter, tmplFileName string) {
	// can read multiple files and store html + go logic into variable defined as *template.Template {name = name of first file and combines all the files }
	parsedTemplate, err := template.ParseFiles("./templates/"+tmplFileName, "./templates/base.layout.tmpl")
	if err != nil {
		log.Println(err)
		return
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func RenderTemplateWhichAddOneByOne(w http.ResponseWriter, tmplFileName string) {
	_, flag := TemplateCache[tmplFileName]
	// if not present in cache
	if !flag {
		err := addTotemplateCache(tmplFileName)
		if err != nil {
			return
		}
	}
	parsedTemplate := TemplateCache[tmplFileName]
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
func addTotemplateCache(tmplFileName string) error {
	files := []string{
		"./templates/" + tmplFileName, "./templates/base.layout.tmpl",
	}
	parsedTemplate, err := template.ParseFiles(files...)
	if err != nil {
		return err
	}
	TemplateCache[tmplFileName] = parsedTemplate
	return nil
}
