package http

import (
	"fmt"
	"iot/integration/mapbox/internal/config"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	config.InitializeConfig()

	r := gin.New()
	Routes(r)
	fmt.Println(config.Sevconfig.Port)

	r.Run(config.Sevconfig.Port)
}
