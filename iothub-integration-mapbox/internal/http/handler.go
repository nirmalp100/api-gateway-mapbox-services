package http

import (
	"fmt"
	"iot/integration/mapbox/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type parms struct {
	Place string `json:"place"`
}

type methods struct {
	Rep services.Repo
}

func (s *methods) ForwardgeoHandler(g *gin.Context) {
	var object parms
	err := g.BindJSON(&object)
	fmt.Println(err)
	if err != nil {
		g.JSON(http.StatusBadRequest, nil)
	}

	fmt.Println(object.Place)
	resp, err := s.Rep.Forwardgeogrpc(object.Place)
	//resp := &pb.ForwardGeoResp{Latitude: 2.2, Longitude: 3.3}
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)

	}

	var Coord struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	Coord.Latitude = resp.Latitude
	Coord.Longitude = resp.Longitude

	g.JSON(http.StatusOK, Coord)

}

///////////////////////////////////////////////////////////////

type coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func (s *methods) ReversegeoHandler(g *gin.Context) {
	var Object coordinates

	err := g.BindJSON(&Object)
	if err != nil {
		g.JSON(http.StatusBadRequest, nil)
	}

	resp, err := s.Rep.Reversegeogrpc(Object.Latitude, Object.Longitude)
	if err != nil {
		log.Fatal(err)
	}

	type Result struct {
		Locality string `json:"locality"`
		Place    string `json:"place"`
		District string `json:"district"`
		State    string `json:"state"`
		Country  string `json:"country"`
	}

	var result Result
	result.Country = resp.Result.Country
	result.State = resp.Result.State
	result.District = resp.Result.Dsitrict
	result.Place = resp.Result.Place
	result.Locality = resp.Result.Locality

	g.JSON(http.StatusOK, &result)
}

///////////////////////////////////////////////////////////

func (s *methods) SuggestApiHandler(g *gin.Context) {
	var object parms
	err := g.BindJSON(&object)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	}

	fmt.Println(object)

	resp, err := s.Rep.Suggestapigrpc(object.Place)
	if err != nil {
		log.Fatal(err)
	}

	result := make([]struct {
		Location string `json:"location"`
		Details  string `json:"details"`
		Address  string `json:"address"`
		Actionid string `json:"actionid"`
	}, len(resp.Suggestion))

	for i := 0; i < len(resp.Suggestion); i++ {
		result[i].Location = resp.Suggestion[i].Location
		result[i].Details = resp.Suggestion[i].Details
		result[i].Address = resp.Suggestion[i].Address
		result[i].Actionid = resp.Suggestion[i].ActionID
	}

	g.JSON(http.StatusOK, &result)
}

///////////////////////////////////////////////////////////

type parm struct {
	Id string `json:"id"`
}

func (s *methods) RetrieveApiHandler(g *gin.Context) {
	var object parm
	err := g.BindJSON(&object)
	if err != nil {
		g.JSON(http.StatusBadRequest, err)
	}

	fmt.Println(object)

	resp, err := s.Rep.Retieveapigrpc(object.Id)
	if err != nil {
		log.Fatal(err)
	}

	type Result struct {
		Location string `json:"location"`
		Details  string `json:"details"`
		Address  string `json:"address"`
	}

	var result Result

	result.Location = resp.Location
	result.Details = resp.Details
	result.Address = resp.Address

	g.JSON(http.StatusOK, &result)
}
