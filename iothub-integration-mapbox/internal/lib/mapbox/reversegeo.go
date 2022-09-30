package mapbox

import (
	"context"
	"iot/integration/mapbox/pb"
)

func (s *Connec) Reversegeogrpc(lat, long string) (*pb.ReverseGeoResp, error) {
	resp, err := s.Client.ReverseGeo(context.Background(), &pb.ReverseGeoReq{Latitude: lat, Longitude: long})
	return resp, err
}
