package DefectDojoApi

type TestResult struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Test `json:"results"`
}
