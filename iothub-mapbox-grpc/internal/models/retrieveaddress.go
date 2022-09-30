package models

type RetrieveResp struct {
	Type     string `json:"type"`
	Features []struct {
		Type     string `json:"type"`
		Geometry struct {
			Coordinates []float64 `json:"coordinates"`
			Type        string    `json:"type"`
		} `json:"geometry"`
		Properties struct {
			FeatureName     string   `json:"feature_name"`
			MatchingName    string   `json:"matching_name"`
			HighlightedName string   `json:"highlighted_name"`
			Description     string   `json:"description"`
			PlaceName       string   `json:"place_name"`
			ID              string   `json:"id"`
			PlaceType       []string `json:"place_type"`
			Context         []struct {
				Layer          string `json:"layer"`
				LocalizedLayer string `json:"localized_layer"`
				Name           string `json:"name"`
			} `json:"context"`
			Maki        string   `json:"maki"`
			PoiCategory []string `json:"poi_category"`
			InternalID  string   `json:"internal_id"`
			ExternalIds struct {
				Tripadvisor string `json:"tripadvisor"`
				Foursquare  string `json:"foursquare"`
				Federated   string `json:"federated"`
				MbxPoi      string `json:"mbx_poi"`
			} `json:"external_ids"`
			MapboxID string `json:"mapbox_id"`
			Metadata struct {
				ReviewCount         int     `json:"review_count"`
				Phone               string  `json:"phone"`
				Website             string  `json:"website"`
				AverageRating       float64 `json:"average_rating"`
				DetailedDescription string  `json:"detailed_description"`
				PrimaryPhoto        []struct {
					Width  int    `json:"width"`
					Height int    `json:"height"`
					URL    string `json:"url"`
				} `json:"primary_photo"`
				OpenHours struct {
					OpenType    string `json:"open_type"`
					OpenPeriods []struct {
						Open struct {
							Day  int `json:"day"`
							Time struct {
								Hour   int `json:"hour"`
								Minute int `json:"minute"`
							} `json:"time"`
						} `json:"open"`
						Close struct {
							Day  int `json:"day"`
							Time struct {
								Hour   int `json:"hour"`
								Minute int `json:"minute"`
							} `json:"time"`
						} `json:"close"`
					} `json:"open_periods"`
				} `json:"open_hours"`
				Iso31661 string `json:"iso_3166_1"`
				Iso31662 string `json:"iso_3166_2"`
			} `json:"metadata"`
		} `json:"properties"`
	} `json:"features"`
	Attribution  string `json:"attribution"`
	Version      string `json:"version"`
	ResponseUUID string `json:"response_uuid"`
}
