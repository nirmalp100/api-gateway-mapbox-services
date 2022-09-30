package mapbox

import (
	"context"
	"iot/integration/mapbox/pb"
)

func (s *Connec) Forwardgeogrpc(place string) (*pb.ForwardGeoResp, error) {

	res, err := s.Client.ForwardGeo(context.Background(), &pb.ForwardGeoReq{Place: place})
	return res, err

}
