package DefectDojoApi

const (
	YYYYMMDD = "2006-01-02"
)

const (
	SeverityCritical               = "Critical"
	SeverityHigh                   = "High"
	SeverityMedium                 = "Medium"
	SeverityLog                    = "Low"
	SeverityInformational          = "Informational"
	NumericalSeverityCritical      = "S0"
	NumericalSeverityHigh          = "S1"
	NumericalSeverityMedium        = "S2"
	NumericalSeverityLow           = "S3"
	NumericalSeverityInformational = "S4"
)

type EngagementResult struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []Engagement `json:"results"`
}
