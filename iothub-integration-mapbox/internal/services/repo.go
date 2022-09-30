package services

import "iot/integration/mapbox/pb"

type Repo interface {
	Forwardgeogrpc(string) (*pb.ForwardGeoResp, error)
	Reversegeogrpc(lat, long string) (*pb.ReverseGeoResp, error)
	Suggestapigrpc(place string) (*pb.SuggestApiResp, error)
	Retieveapigrpc(id string) (*pb.RetrieveResp, error)
}

//this interface contains the methods in the library
