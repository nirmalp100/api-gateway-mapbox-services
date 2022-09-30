package models

type ReverseGeoAddress struct {
	Type     string    `json:"type"`
	Query    []float64 `json:"query"`
	Features []struct {
		ID         string   `json:"id"`
		Type       string   `json:"type"`
		PlaceType  []string `json:"place_type"`
		Relevance  int      `json:"relevance"`
		Properties struct {
			Foursquare string `json:"foursquare"`
			Landmark   bool   `json:"landmark"`
			Category   string `json:"category"`
		} `json:"properties"`
		Text      string    `json:"text"`
		PlaceName string    `json:"place_name"`
		Center    []float64 `json:"center"`
		Geometry  struct {
			Coordinates []float64 `json:"coordinates"`
			Type        string    `json:"type"`
		} `json:"geometry"`
		Context []struct {
			ID        string `json:"id"`
			Text      string `json:"text"`
			Wikidata  string `json:"wikidata,omitempty"`
			ShortCode string `json:"short_code,omitempty"`
		} `json:"context"`
	} `json:"features"`
	Attribution string `json:"attribution"`
}
