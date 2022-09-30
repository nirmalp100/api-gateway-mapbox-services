package main

import (
	"context"
	"fmt"
	"log"
	"m/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//transport security not added
	conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	client := pb.NewGeoCodeClient(conn)

	resp, err := client.ReverseGeo(context.Background(), &pb.ReverseGeoReq{Latitude: "12.983154473350895", Longitude: "77.59549735301529"})
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.ForwardGeo(context.Background(), &pb.ForwardGeoReq{Place: "Los Angeles"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.Result)
	fmt.Println(res.Latitude)
}
