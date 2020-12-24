package main

import (
	"github.com/gorilla/mux"
	"github.com/kitche/site/config"
	"github.com/kitche/site/routers"
	"log"
	"net/http"
)


func main() {

	// Initialise the connection pool.
	db, err := config.NewDB("postgres://user:pass@localhost/bookstore")
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	env := &config.Env{DB: db}

	// Create an instance of Env containing the connection pool
	r := mux.NewRouter()
	r.HandleFunc("/", routers.HomeHandler(env))
	//r.HandleFunc("/servers", routers.ServerHandler)
	// r.HandleFunc("/articles", routers.ArticleHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}

