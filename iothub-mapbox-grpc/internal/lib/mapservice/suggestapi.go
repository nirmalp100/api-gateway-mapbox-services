package mapservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"m/internal/models"
	"m/pb"
	"strings"
)

func (s MapService) MapboxSuggestions(ctx context.Context, Req *pb.SuggestApiReq) (*pb.SuggestApiResp, error) {

	content, err := ioutil.ReadFile("./internal/models/sampleresponse/suggestions.json")

	if err != nil {
		fmt.Println(err)
	}

	var resp models.SuggestResponse

	json.Unmarshal(content, &resp)

	var data []string

	for i := 0; i < len(resp.Suggestions); i++ {

		data = append(data, resp.Suggestions[i].MatchingName+"$"+resp.Suggestions[i].Description+"$"+resp.Suggestions[i].Coordinates+"$"+resp.Suggestions[i].Action.Body.ID)

	}

	var res pb.SuggestApiResp

	for i := 0; i < len(data); i++ {
		d := strings.Split(data[i], "$")

		res.Suggestion = append(res.Suggestion, &pb.Suggestions{Details: d[0], Address: d[1], Location: d[2], ActionID: d[3]})
	}

	return &res, nil
}
