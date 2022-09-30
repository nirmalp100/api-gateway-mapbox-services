package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	MAPBOX_ACCESS_TOKEN string `default:"mapbox:key"`
}

type ServerConfiguration struct {
	ServicePort string `default:"4040"`
}

var Congig Configuration

var Serconfig ServerConfiguration

func IntializeConfig() {
	err := envconfig.Process("", &Congig)
	if err != nil {
		log.Fatal(err)
	}

	err = envconfig.Process("", &Serconfig)
	if err != nil {
		log.Fatal(err)
	}
}
