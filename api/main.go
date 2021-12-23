package main

import (
	"adminbackend/models"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {

	app := &application{

		config: appConfig{
			port: 4000,
		},
		DB: models.NewDBModel(openDB()),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", app.config.port),
		Handler: app.Routes(),
	}
	
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal("server not connected!")
	}
	fmt.Println("server listening........")

}
