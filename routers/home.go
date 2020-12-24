package routers

import (
	"github.com/kitche/site/models"
	"log"
	"net/http"
	"github.com/kitche/site/config"
)


func HomeHandler(env *config.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bks, err := models.AllServers(env.DB)

		if err = config.Hometpl.Execute(w, bks); err != nil {
			//if err:= templates["index"].Execute(w, "hello"); err != nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
		}
		}
	}