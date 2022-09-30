package mapbox

import (
	"iot/integration/mapbox/pb"

	"google.golang.org/grpc"
)

type Connec struct {
	Client pb.GeoCodeClient
}

func NewClientConnection(C *grpc.ClientConn) pb.GeoCodeClient {
	return pb.NewGeoCodeClient(C)

}
