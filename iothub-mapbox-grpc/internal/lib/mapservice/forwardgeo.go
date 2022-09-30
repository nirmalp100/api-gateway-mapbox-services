package mapservice

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"m/internal/config"
	"m/internal/models"
	"m/pb"
	"net/http"
	"strings"
)

func (s MapService) MapboxForwardGeo(ctx context.Context, Req *pb.ForwardGeoReq) (*pb.ForwardGeoResp, error) {

	pl := Req.GetPlace()
	b := strings.Split(pl, " ")
	var basic string
	if len(b) == 1 {
		basic = b[0]
	} else if len(b) == 2 {
		basic = b[0] + "%20" + b[1]
	} else if len(b) == 3 {
		basic = b[0] + "%20" + b[1] + "%20" + b[2]
	} else {
		return nil, errors.New("invalid request")
	}

	YOUR_MAPBOX_ACCESS_TOKEN := config.Congig.MAPBOX_ACCESS_TOKEN
	url := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%s.json?access_token=%v", basic, YOUR_MAPBOX_ACCESS_TOKEN)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var rsp models.ForwardGeoAddress

	err = json.Unmarshal(body, &rsp)
	if err != nil {
		log.Fatal(err)
	}

	var data []float64

	for i := 0; i < len(rsp.Features); i++ {
		for j := 0; j < len(rsp.Features[i].Center); j++ {
			data = append(data, rsp.Features[i].Center...)
		}
	}

	lat := data[0]
	long := data[1]

	return &pb.ForwardGeoResp{Latitude: lat, Longitude: long}, nil

}
