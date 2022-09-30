package http

import (
	"iot/integration/mapbox/internal/config"
	"iot/integration/mapbox/internal/lib/mapbox"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Routes(g *gin.Engine) {

	conn := StartClient()

	c := mapbox.NewClientConnection(conn)

	s := &mapbox.Connec{Client: c}
	m := methods{Rep: s}

	g.GET("/forwardgeo", m.ForwardgeoHandler)
	g.GET("/reversegeo", m.ReversegeoHandler)
	g.GET("/suggest", m.SuggestApiHandler)
	g.GET("/retrieve", m.RetrieveApiHandler)

}

func StartClient() *grpc.ClientConn {

	conn, err := grpc.Dial(config.GrpcTarget.Target, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}
	return conn

}
