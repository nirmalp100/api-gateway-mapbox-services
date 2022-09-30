package server

import (
	"fmt"
	"log"
	"m/pb"
	"net"

	"m/internal/config"
	"m/internal/services/geoservice"

	mapSvc "m/internal/lib/mapservice"

	"google.golang.org/grpc"
)

func RunServer() {
	config.IntializeConfig()
	fmt.Println(config.Serconfig.ServicePort)
	err := StartServer(config.Serconfig.ServicePort)
	if err != nil {
		log.Fatal(err)
	}

}

func StartServer(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	mapSvc := mapSvc.NewMapService()
	//svc := &geoservice.Service{}
	svc := geoservice.NewService(mapSvc)

	pb.RegisterGeoCodeServer(server, svc)
	err = server.Serve(listener)
	if err != nil {
		return err
	}
	return nil

}
