package config

import "html/template"

var baseTemplate = "templates/layout/_base.html"
//var templates map[string]*template.Template
var Hometpl *template.Template
var Servertpl *template.Template

func LoadTemplates() {
	Hometpl = template.Must(template.ParseFiles(baseTemplate, "templates/home/index.html"))
	Servertpl = template.Must(template.ParseFiles(baseTemplate, "templates/home/server.html"))

}