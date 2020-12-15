package routers

import (
	"net/http"
	"panel/config"
)

func init() {
	config.LoadTemplates()
}

func ServerHandler(w http.ResponseWriter, r *http.Request) {
	//	c, err := r.Cookie("session_token")
	//	if err != nil {
	//		if err == http.ErrNoCookie {
	// If the cookie is not set, return an unauthorized status
	//			w.WriteHeader(http.StatusUnauthorized)
	//			return
	//		}
	if err := config.Servertpl.Execute(w, nil); err != nil {
		//if err:= templates["index"].Execute(w, "hello"); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
func EditServerHandler(w http.ResponseWriter, r *http.Request) {

}
