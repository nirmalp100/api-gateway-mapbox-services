package geoservice

import (
	"context"
	"m/pb"
)

type Repo interface {
	MapboxReverseGeo(ctx context.Context, Req *pb.ReverseGeoReq) (*pb.ReverseGeoResp, error)
	MapboxForwardGeo(ctx context.Context, Req *pb.ForwardGeoReq) (*pb.ForwardGeoResp, error)
	MapboxSuggestions(ctx context.Context, Req *pb.SuggestApiReq) (*pb.SuggestApiResp, error)
	MapboxRetrieve(ctx context.Context, Req *pb.RetrieveReq) (*pb.RetrieveResp, error)
}
