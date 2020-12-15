package routers

import (
	"net/http"
	"panel/config"
)
//var templates map[string]*template.Template

func init() {
	config.LoadTemplates()
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// We can obtain the session token from the requests cookies, which come with every request
	//	c, err := r.Cookie("session_token")
	//	if err != nil {
	//		if err == http.ErrNoCookie {
	// If the cookie is not set, return an unauthorized status
	//			w.WriteHeader(http.StatusUnauthorized)
	//			return
	//		}
	if err := config.Hometpl.Execute(w, nil); err != nil {
		//if err:= templates["index"].Execute(w, "hello"); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}