package DefectDojoApi

type FindingResults struct {
	Count    int       `json:"count"`
	Next     *string   `json:"next"`
	Previous *string   `json:"previous"`
	Results  []Finding `json:"results"`
	Prefetch struct {
	} `json:"prefetch"`
}
