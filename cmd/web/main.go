package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/herby0sbourne/booking/internal/config"
	"github.com/herby0sbourne/booking/internal/handlers"
	"github.com/herby0sbourne/booking/internal/models"
	"github.com/herby0sbourne/booking/internal/render"

	"github.com/alexedwards/scs/v2"
)

const PORT = ":3001"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	//
	gob.Register(models.Reservation{})

	// change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

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
