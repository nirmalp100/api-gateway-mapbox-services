package mapbox

import (
	"context"
	"iot/integration/mapbox/pb"
)

func (s *Connec) Suggestapigrpc(place string) (*pb.SuggestApiResp, error) {
	return s.Client.Suggestions(context.Background(), &pb.SuggestApiReq{Place: place})
}
