package main

import (
	"adminbackend/models"
	"log"
)

type application struct {
	DB     models.DB
	config appConfig
	log    *log.Logger
}

type appConfig struct {
	port int
}
