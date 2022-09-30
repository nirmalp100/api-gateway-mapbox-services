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
	"strconv"
)

func IsValidCoordinates(Req *pb.ReverseGeoReq) (string, string, bool) {
	lat := Req.GetLatitude()
	long := Req.GetLongitude()
	latitude, err1 := strconv.ParseFloat(lat, 64)
	longitude, err2 := strconv.ParseFloat(long, 64)
	if err1 != nil || err2 != nil || latitude < -90 || latitude > 90 || longitude < -180 || longitude > 180 {
		return "", "", false
	}

	return lat, long, true
}

func (s MapService) MapboxReverseGeo(ctx context.Context, Req *pb.ReverseGeoReq) (*pb.ReverseGeoResp, error) {
	lat1, long1, valid := IsValidCoordinates(Req)
	if !valid {
		return nil, errors.New("invalid coordinates")
	}

	YOUR_MAPBOX_ACCESS_TOKEN := config.Congig.MAPBOX_ACCESS_TOKEN
	url := fmt.Sprintf("https://api.mapbox.com/geocoding/v5/mapbox.places/%v,%v.json?access_token=%v", long1, lat1, YOUR_MAPBOX_ACCESS_TOKEN)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var respo models.ReverseGeoAddress

	json.Unmarshal(body, &respo)

	var data []string
	for i := 0; i < len(respo.Features); i++ {
		for j := 0; j < len(respo.Features[i].Context); j++ {
			data = append(data, respo.Features[i].Context[j].Text)

		}
	}

	return &pb.ReverseGeoResp{Result: &pb.Result{Locality: data[0], Place: data[1], Dsitrict: data[2], State: data[3], Country: data[4]}}, nil
}
