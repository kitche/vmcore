package config

import (
	"html/template"
	"database/sql"

	_ "github.com/lib/pq"
	)

var baseTemplate *template.Template
//var templates map[string]*template.Template
var Hometpl *template.Template
var Servertpl *template.Template

var db *sql.DB

func LoadTemplates() {
	baseTemplate = "templates/layout/_base.html"
	Hometpl = template.Must(template.ParseFiles(baseTemplate, "templates/home/index.html"))
	Servertpl = template.Must(template.ParseFiles(baseTemplate, "templates/home/server.html"))

}

	// InitDB sets up setting up the connection pool global variable.
	func InitDB(dataSourceName string) error {
		var err error

		db, err = sql.Open("postgres", dataSourceName)
		if err != nil {
		return err
	}

		return db.Ping()
	}
