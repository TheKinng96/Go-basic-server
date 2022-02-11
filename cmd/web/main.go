package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/TheKinng96/Go-basic-server/pkg/config"
	"github.com/TheKinng96/Go-basic-server/pkg/controllers"
	"github.com/TheKinng96/Go-basic-server/pkg/render"
)

const portNumber = ":8000"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := controllers.NewRepo(&app)
	controllers.NewControllers(repo)
	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)
}
