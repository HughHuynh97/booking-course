package main

import (
	"github.com/HughHuynh97/booking-course/pkg/config"
	"github.com/HughHuynh97/booking-course/pkg/handlers"
	"github.com/HughHuynh97/booking-course/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = true
	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", repo.Home)
	http.HandleFunc("/about", repo.About)

	_ = http.ListenAndServe(portNumber, nil)
}
