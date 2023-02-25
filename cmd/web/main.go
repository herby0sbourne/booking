package main

import (
	"fmt"
	"log"
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const PORT = ":3001"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("server running on http://localhost%s", PORT))
	// _ = http.ListenAndServe("localhost:3001", nil)
	// _ = http.ListenAndServe(PORT, nil)

	server := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = server.ListenAndServe()
	log.Fatal(err)

}
