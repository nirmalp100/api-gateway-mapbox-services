package mapbox

import (
	"context"
	"iot/integration/mapbox/pb"
)

func (s *Connec) Retieveapigrpc(id string) (*pb.RetrieveResp, error) {
	return s.Client.Retrieve(context.Background(), &pb.RetrieveReq{Id: id})
}
