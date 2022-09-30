package geoservice

import (
	"context"
	"m/pb"
)

type Service struct {
	repo Repo
	pb.UnimplementedGeoCodeServer
}

func NewService(repoSvc Repo) *Service {
	return &Service{repo: repoSvc}
}
func (s *Service) ReverseGeo(ctx context.Context, Req *pb.ReverseGeoReq) (*pb.ReverseGeoResp, error) {

	return s.repo.MapboxReverseGeo(ctx, Req)

}

func (s *Service) ForwardGeo(ctx context.Context, Req *pb.ForwardGeoReq) (*pb.ForwardGeoResp, error) {
	return s.repo.MapboxForwardGeo(ctx, Req)

}

func (s *Service) Retrieve(ctx context.Context, Req *pb.RetrieveReq) (*pb.RetrieveResp, error) {
	return s.repo.MapboxRetrieve(ctx, Req)

}

func (s *Service) Suggestions(ctx context.Context, Req *pb.SuggestApiReq) (*pb.SuggestApiResp, error) {
	return s.repo.MapboxSuggestions(ctx, Req)

}
