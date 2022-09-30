package mapservice

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"m/internal/models"
	"m/pb"
	"strconv"
)

func (s MapService) MapboxRetrieve(ctx context.Context, Req *pb.RetrieveReq) (*pb.RetrieveResp, error) {
	content, err := ioutil.ReadFile("./internal/models/sampleresponse/retrieve.json")

	if err != nil {
		log.Fatal(err)
	}
	var retree models.RetrieveResp

	err = json.Unmarshal(content, &retree)
	if err != nil {
		log.Fatal(err)
	}

	float1 := strconv.FormatFloat(retree.Features[0].Geometry.Coordinates[0], 'f', 6, 64)
	float2 := strconv.FormatFloat(retree.Features[0].Geometry.Coordinates[1], 'f', 6, 64)
	flot := float1 + "," + float2

	L := flot
	D := retree.Features[0].Properties.HighlightedName
	A := retree.Features[0].Properties.Description

	// r := make([]struct {
	// 	Detail   string
	// 	Addres   string
	// 	Location []float64
	// }, len(retree.Features))

	// for i := 0; i < len(retree.Features); i++ {
	// 	r[i].Detail = retree.Features[i].Properties.HighlightedName
	// 	r[i].Addres = retree.Features[i].Properties.Description
	// 	r[i].Location = []float64{retree.Features[i].Geometry.Coordinates[0], retree.Features[i].Geometry.Coordinates[1]}
	// }

	return &pb.RetrieveResp{Location: L, Details: D, Address: A}, nil
}
