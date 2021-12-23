package config

import (
	"adminbackend/models"
	"log"
)

type Application struct {
	DB     models.DB
	Config AppConfig
	Log    *log.Logger
}

type AppConfig struct {
	Port int
}
