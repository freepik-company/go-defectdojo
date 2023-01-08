package DefectDojoApi

type EngagementResult struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []Engagement `json:"results"`
}
