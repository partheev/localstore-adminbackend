package main

import (
	"adminbackend/config"
	"adminbackend/models"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type application config.Application

func main() {
	app := application(config.Application{

		Config: config.AppConfig{
			Port: 4000,
		},
		DB: models.NewDBModel(config.OpenDB()),
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", app.Config.Port),
		Handler: app.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("server not connected!")
	}
	fmt.Println("server listening........")

}
