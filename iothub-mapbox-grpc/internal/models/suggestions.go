package models

type SuggestResponse struct {
	Suggestions []struct {
		FeatureName     string   `json:"feature_name"`
		MatchingName    string   `json:"matching_name"`
		HighlightedName string   `json:"highlighted_name"`
		Description     string   `json:"description"`
		ResultType      []string `json:"result_type"`
		Language        string   `json:"language"`
		Action          struct {
			Endpoint string `json:"endpoint"`
			Method   string `json:"method"`
			Body     struct {
				ID string `json:"id"`
			} `json:"body"`
			MultiRetrievable bool `json:"multi_retrievable"`
		} `json:"action"`
		Coordinates string   `json:"coordinates"`
		Maki        string   `json:"maki"`
		Category    []string `json:"category"`
		InternalID  string   `json:"internal_id"`
		ExternalIds struct {
			MbxPoi      string `json:"mbx_poi"`
			Foursquare  string `json:"foursquare"`
			Federated   string `json:"federated"`
			Tripadvisor string `json:"tripadvisor"`
		} `json:"external_ids"`
		MapboxID string `json:"mapbox_id"`
		Context  []struct {
			Layer          string `json:"layer"`
			LocalizedLayer string `json:"localized_layer"`
			Name           string `json:"name"`
		} `json:"context"`
		Metadata struct {
			Iso31661 string `json:"iso_3166_1"`
			Iso31662 string `json:"iso_3166_2"`
		} `json:"metadata"`
	} `json:"suggestions"`
	Attribution  string `json:"attribution"`
	Version      string `json:"version"`
	ResponseUUID string `json:"response_uuid"`
}
