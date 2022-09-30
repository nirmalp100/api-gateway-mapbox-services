package models

type NearbySearchModel struct {
	Type     string   `json:"type"`
	Query    []string `json:"query"`
	Features []struct {
		ID         string   `json:"id"`
		Type       string   `json:"type"`
		PlaceType  []string `json:"place_type"`
		Relevance  float64  `json:"relevance"`
		Properties struct {
			Wikidata string `json:"wikidata"`
		} `json:"properties"`
		TextEnUS          string    `json:"text_en-US"`
		LanguageEnUS      string    `json:"language_en-US,omitempty"`
		PlaceNameEnUS     string    `json:"place_name_en-US"`
		Text              string    `json:"text"`
		Language          string    `json:"language,omitempty"`
		PlaceName         string    `json:"place_name"`
		MatchingPlaceName string    `json:"matching_place_name,omitempty"`
		Bbox              []float64 `json:"bbox,omitempty"`
		Center            []float64 `json:"center"`
		Geometry          struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Context []struct {
			ID           string `json:"id"`
			Wikidata     string `json:"wikidata"`
			TextEnUS     string `json:"text_en-US"`
			LanguageEnUS string `json:"language_en-US"`
			Text         string `json:"text"`
			Language     string `json:"language"`
			ShortCode    string `json:"short_code,omitempty"`
		} `json:"context"`
	} `json:"features"`
	Attribution string `json:"attribution"`
}
