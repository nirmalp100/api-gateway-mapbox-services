package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type HttpServerConfiguration struct {
	Port string `default:":3000"`
}

type GrpcTargetString struct {
	Target string `default:":4040"`
}

var Sevconfig HttpServerConfiguration
var GrpcTarget GrpcTargetString

func InitializeConfig() {
	err := envconfig.Process("", &Sevconfig)

	if err != nil {
		log.Fatal(err)
	}

	err = envconfig.Process("", &GrpcTarget)

	if err != nil {
		log.Fatal(err)
	}
}
