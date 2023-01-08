package DefectDojoApi

type ProductResult struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Product `json:"results"`
}
