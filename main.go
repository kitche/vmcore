package main

import (
 "github.com/gorilla/mux"
 "log"
 "net/http"
 "panel/config"
 "panel/routers"
)

func init(){
 config.LoadTemplates()
}

func main() {

 err := config.InitDB("postgres://user:pass@localhost/database")
 if err != nil {
  log.Fatal(err)
 }

 r := mux.NewRouter()
 r.HandleFunc("/", routers.HomeHandler)
 r.HandleFunc("/servers", routers.ServerHandler)
// r.HandleFunc("/articles", routers.ArticleHandler)
 http.Handle("/", r)
 http.ListenAndServe(":8080", r)
}

